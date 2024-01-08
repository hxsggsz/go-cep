package main

import (
	"cep/api/cep"
	"fmt"
)

func main() {
	rawCep := ""

	fmt.Println("Insira seu cep: ")
	fmt.Scanln(&rawCep)

	cepReq := cep.RequestCepInfo(rawCep)

	fmt.Printf("Logradouro: %v, Complemento: %v, Bairro: %v e Localidade: %v", cepReq.Logradouro, cepReq.Complemento, cepReq.Bairro, cepReq.Localidade)
}
