package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(50, 20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

//	func sum(a int, b int) int {
//		return a + b
//	}

//	func sum(a, b int) (int, bool) {
//		if a+b >= 50 {
//			return a + b, true
//		}
//		return a + b, false
//	}
func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
