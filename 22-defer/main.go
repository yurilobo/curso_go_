package main

func main() {
	println("Primeira linha")
	defer println("Segunda linha")
	println("Terceira linha")
	println("Quarta linha")
	println("Quinta linha")
	//defer req.Body.Close()///é usado para ser o ultimo a ser chamado

}
