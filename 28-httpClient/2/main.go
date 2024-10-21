package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	//c := http.Client{Timeout: time.Second} ///time.Microsecond
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "yuri"}`)) //recebe um slice de bates e devolve um buffer-espaço com informacao
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil) //pega os dados do stdout
}
