package main

import "fmt"

const a = "hello, world!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 30
	meuArray[2] = 40

	fmt.Print(meuArray[0])
	fmt.Print(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("  O valor do indice %d e o valor é é %d\n", i, v)
	}
}
