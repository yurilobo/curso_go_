package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
}

func main() {
	yuri := Cliente{
		Nome:  "Yuri",
		Idade: 29,
		Ativo: true,
	}

	yuri.Endereco.Cidade = "Fortaleza"
	yuri.Ativo = false
	yuri.Desativar()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", yuri.Nome, yuri.Idade, yuri.Ativo)
}
