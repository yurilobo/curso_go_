package main

import (
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}
type Cursos []Curso

// para transformar as funcoes em up
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		} //para ter varias paginas
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper}) ///acrescenta um par chave valor
		t = template.Must(t.ParseFiles(templates...))
		//t := template.Must(template.New("content.html").ParseFiles(templates...)) //trago a funcao variatica
		err := t.Execute(os.Stdout, Cursos{
			{"Golang", 40},
			{"Java", 20},
			{"Python", 10},
		})

		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)

}
