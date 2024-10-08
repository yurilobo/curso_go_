package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	//tamanho, err := f.WriteString("Hello, world!!") se soouber que todo que vai ser retornado é uma string
	if err != nil {
		panic(err)
	}
	fmt.Printf("Aquivo criado com sucesso!!, Tamanho %d bytes", tamanho)

	f.Close()
	///para leitura de arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(arquivo))

	//leitura pouco a pouco
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[n]))
	}
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
