package main

import "net/http"

func main() {
	mux := http.NewServeMux() //responsavel por atachar as urls/hendlws
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8080", mux) //se usar o nil em vez de mux vai ter o statudo defou e só um servidor
	mux2 := http.NewServeMux()        ///é possível criar mais de um servidor
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Segundo servido. este em uma porta difetente"))
	})
	http.ListenAndServe(":8081", mux2)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
