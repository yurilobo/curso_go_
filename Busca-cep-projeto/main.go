package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}
func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	} //se o complemento da url nao vier com / trazemos uma resultado false
	cepParam := r.URL.Query().Get("cep") //peguei o paremetro do cep
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} //se o bad requesto for igual a em branco retornar a funcao
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json") //se nao der erro muda o cabeçcalho para json
	w.WriteHeader(http.StatusOK)                       //fala que o status
	json.NewEncoder(w).Encode(cep)                     //pega a requisicao e o resultado joga dentro de cep
}
func BuscaCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/") //faz uma busca
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body) //pega o valor que foi retornado
	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c) //pega todos os dados qe estao no endereco em json e transforma em struct
	if error != nil {
		return nil, error

	}
	return &c, nil
}
