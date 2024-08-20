//funcao variática

package main

import "fmt"

func main() {
	fmt.Println(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 5435)
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
