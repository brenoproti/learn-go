package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Address struct {
	Cep        string `json:"cep"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func main() {
	http.HandleFunc("/", GetAddressHandler)
	http.ListenAndServe(":8080", nil)
}

func GetAddressHandler(w http.ResponseWriter, r *http.Request) {
	zipCode := r.URL.Query().Get("cep")
	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP is required"))
		return
	}

	address, err := GetAddresss(zipCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// res, err := json.Marshal(address)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(res)

	json.NewEncoder(w).Encode(address)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func GetAddresss(zipCode string) (*Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var address Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
