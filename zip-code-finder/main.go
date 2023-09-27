package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep        string `json:"cep"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func main() {
	zipCode := os.Args[1]

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var address Address
	err = json.Unmarshal(res, &address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	addressData := []string{}
	addressData = append(addressData, fmt.Sprintf("CEP: %v\n", address.Cep))
	addressData = append(addressData, fmt.Sprintf("Cidade/Estado: %v/%v\n", address.Localidade, address.Uf))
	addressData = append(addressData, fmt.Sprintf("Bairro: %v\n", address.Bairro))
	addressData = append(addressData, fmt.Sprintf("Logradouro: %v", address.Logradouro))

	file, err := os.Create("zip-code-finder/endereco.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, v := range addressData {
		fmt.Print(v)
		file.WriteString(v)
	}
}
