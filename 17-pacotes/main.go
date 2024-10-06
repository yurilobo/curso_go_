package main

import (
	"curso/17-pacotes/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado : ", s)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro)
	fmt.Println(matematica.A)
	fmt.Println(carro.Andar())
}
