package main

import (
	"log"
	"net/http"
)

// Definindo handler pra home que escreve um slice de byte contendo um texto como response body (preciso converter
// o byte dps?)
func home(w http.ResponseWriter, r *http.Request) {
	// Checando se a URL é valida, caso nao, retorna um 404 com http.NotFound
	// Importante dar o return pra ele parar de executar
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from J`s Snippetbox!"))
}

// Adicionando func com handler pra visualizar apenas uma anotacao
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific Snippet"))
}

// Adicionando func com handler pra criar uma anotacao
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Usando r.method pra checar se a request esta usando o metodo que eu quero
	if r.Method != "POST" {
		// Informando ao user o tipo de metodo permitido, primeiro param é o nome do header e segundo o método
		w.Header().Set("Allow", "POST")
		// Se nao estiver, retorno um 405 com a mensagem abaixo. Retorna a func pro codigo a seguir nao ser executado
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed."))
		return
	}
	w.Write([]byte("Create new Snippet"))
}

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
