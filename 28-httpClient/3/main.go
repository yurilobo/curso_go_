package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{} // client httop

	req, err := http.NewRequest("GET", "http://google.com", nil) //obsejto de requisição
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json") //cria um header
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
