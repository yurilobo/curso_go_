package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	yuri := Cliente{
		Nome:  "Yuri",
		Idade: 29,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", yuri.Nome, yuri.Idade, yuri.Ativo)
}
