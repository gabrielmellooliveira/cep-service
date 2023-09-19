package services

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ViaCepService struct {
	Url string
}

func NewViaCepService(url string) *ViaCepService {
	return &ViaCepService{
		Url: url,
	}
}

type ViaCepResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func (s *ViaCepService) GetCepInfo(cep string) (*ViaCepResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", s.Url+cep+"/json", nil)
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

	var data ViaCepResponse
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Printf("Erro ao fazer a conversão do JSON para a cotação: %v", err)
	}

	return &data, nil
}
