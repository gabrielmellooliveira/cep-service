package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type ApiCepService struct {
	Url string
}

func NewApiCepService(url string) *ApiCepService {
	return &ApiCepService{
		Url: url,
	}
}

type ApiCepResponse struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func (s *ApiCepService) GetCepInfo(cep string) (*ApiCepResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	cepFirstPart := strings.TrimSpace(cep[0:5])
	cepSecondPart := strings.TrimSpace(cep[5:])

	cepFormatted := cepFirstPart + "-" + cepSecondPart

	req, err := http.NewRequestWithContext(ctx, "GET", s.Url+cepFormatted+".json", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erro ao buscar a cotação: %v", err)
	}

	var data ApiCepResponse
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Printf("Erro ao fazer a conversão do JSON para a cotação: %v", err)
	}

	return &data, nil
}
