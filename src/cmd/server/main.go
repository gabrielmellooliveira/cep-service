package main

import (
	"cep-service/src/internal/infra/webserver/handlers"
	"cep-service/src/internal/services"
	"fmt"
	"os"
)

func main() {
	apiCepService := services.NewApiCepService("https://cdn.apicep.com/file/apicep/")
	viaCepService := services.NewViaCepService("https://viacep.com.br/ws/")

	cepHandler := handlers.NewCepHandler(*apiCepService, *viaCepService)

	var cep string

	fmt.Print("Digite seu CEP: ")
	fmt.Fscan(os.Stdin, &cep)

	cepResponse := cepHandler.GetCepInfo(cep)

	fmt.Println()
	fmt.Printf("Cep: %s \n", cepResponse.Cep)
	fmt.Printf("Address: %s \n", cepResponse.Address)
	fmt.Printf("District: %s \n", cepResponse.District)
	fmt.Printf("City: %s \n", cepResponse.City)
	fmt.Printf("State: %s \n", cepResponse.State)
	fmt.Printf("Provider: %s \n", cepResponse.Provider)

}
