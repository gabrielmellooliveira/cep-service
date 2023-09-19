package handlers

import (
	"cep-service/src/internal/services"
	"time"
)

type CepHandler struct {
	ApiCepService services.ApiCepService
	ViaCepService services.ViaCepService
}

func NewCepHandler(apiCepService services.ApiCepService, viaCepService services.ViaCepService) *CepHandler {
	return &CepHandler{
		ApiCepService: apiCepService,
		ViaCepService: viaCepService,
	}
}

type CepResponse struct {
	Cep      string `json:"cep"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
	Provider string `json:"provider"`
}

func (h *CepHandler) GetCepInfo(cep string) *CepResponse {
	channelApiCepUrl := make(chan CepResponse)
	channelViaCepUrl := make(chan CepResponse)

	// ApiCepService
	go func() {
		response, err := h.ApiCepService.GetCepInfo(cep)
		if err != nil {
			panic(err)
		}

		message := CepResponse{
			response.Code,
			response.State,
			response.City,
			response.District,
			response.Address,
			"Api Cep",
		}

		channelApiCepUrl <- message
	}()

	// ViaCepService
	go func() {
		response, err := h.ViaCepService.GetCepInfo(cep)
		if err != nil {
			panic(err)
		}

		message := CepResponse{
			response.Cep,
			response.Uf,
			response.Localidade,
			response.Bairro,
			response.Logradouro,
			"Via Cep",
		}

		channelApiCepUrl <- message
	}()

	select {
	case cepResponse := <-channelApiCepUrl:
		return &cepResponse

	case cepResponse := <-channelViaCepUrl:
		return &cepResponse

	case <-time.After(time.Second * 1):
		panic("Timeout")
		return nil
	}
}
