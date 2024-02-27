package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Registrando as rotas e linkando elas com as minhas funcs (controllers) acima
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Starting server on :4000")

	// Usando http.ListenAndServe() pra iniciar um server novo. Passo dois parametros, o TCP network adress
	// (:4000 nese caso) e o servemux que criei agora. Se essa func retornar um erro, uso log.Fatal pra logar o erro
	// e sair do app. Todo erro retornado por essa func é non-nil
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
