package main

type Conta struct {
	saldo int
}

// ao usar o * alteramos o endeço se não usamos entao ultilizamos apenas uma copia
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
	// c.nome = "Yuri Anderson"
	// fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	// yuri := Client{
	// 	nome: "Yuri",
	// }
	// yuri.andou()
	// fmt.Printf("O valor da Struct com o nome %v\n", yuri.nome)

	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
