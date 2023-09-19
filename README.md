# CEP Service (Serviço de CEP)

Esse repositório faz parte de um desafio do curso [Go Expert](https://goexpert.fullcycle.com.br/curso).

## Descrição do desafio

Você precisará criar uma api que busque o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:
- `https://cdn.apicep.com/file/apicep/" + cep + ".json`
- `http://viacep.com.br/ws/" + cep + "/json/`
 
Os requisitos para cumprir este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Como rodar o projeto

- Entrar na pasta `/src`
- Rodar o comando `go run main.go`
