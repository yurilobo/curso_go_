package main

func main() {

	//Memoria aponta para o endereço que aponta para um valor

	//variavel aponta para um ponteiro que tem um indereço na memoria e tem um valor
	a := 10
	//println(&a)
	var ponteiro *int = &a
	//println(ponteiro)

	*ponteiro = 20
	b := &a
	println(&b) //lugar que esta apontando
	println(*b) //valor do que esta sendo apontado

	*b = 504
	println(&b) //lugar que esta apontando
	println(*b) //valor do que esta sendo apontado

}
