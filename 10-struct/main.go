package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// nas interfaces de go so e possivel setar metodos
type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativo", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}
func main() {
	yuri := Cliente{
		Nome:  "Yuri",
		Idade: 29,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
	yuri.Endereco.Cidade = "Fortaleza"
	yuri.Ativo = false
	yuri.Desativar()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", yuri.Nome, yuri.Idade, yuri.Ativo)
}
