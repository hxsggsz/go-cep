package cep

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
}

func (c *Cep) InitialCep() Cep {
	return Cep{
		Cep:         "Cep",
		Logradouro:  "Logradouro",
		Complemento: "Complemento",
		Bairro:      "Bairro",
		Localidade:  "Localidade",
	}
}
