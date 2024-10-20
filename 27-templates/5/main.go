package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		} //para ter varias paginas
		t := template.Must(template.New("content.html").ParseFiles(templates...)) //trago a funcao variatica
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 20},
			{"Python", 10},
		})

		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)

}
