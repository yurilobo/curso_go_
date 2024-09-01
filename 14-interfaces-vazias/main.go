package main

import "fmt"

// interface que nao tem metodo é usada para todo mundo
type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Helllo World!"
	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
