package cep

import (
	"encoding/json"
	"net/http"
)

func RequestCepInfo(cep string) Cep {
	api := "https://viacep.com.br/ws/" + cep + "/json"

	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var parsedCep Cep

	err = json.NewDecoder(res.Body).Decode(&parsedCep)

	if err != nil {
		panic(err)
	}

	return parsedCep
}
