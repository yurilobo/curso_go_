package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 2000, "João": 3000, "Davi": 2100}
	// fmt.Println(salarios["Davi"])

	delete(salarios, "Wesley")
	salarios["Wes"] = 3500
	// fmt.Println(salarios["Wes"])

	// sal:=make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Wesley"]=1222

	//percorrendo salario
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s, é %d \n", nome, salario)

	}
	for _, salario := range salarios {
		fmt.Printf("O salario de %s, é %d \n", salario)

	}
}
