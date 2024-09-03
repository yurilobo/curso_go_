package main

import "fmt"

func main() {
	var minhaVar interface{} = "Yuri Anderson"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)

	fmt.Printf(" o valor de res é %v o valor de ok é %v\n", res, ok)
}
