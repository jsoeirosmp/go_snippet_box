package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	// Extraindo o valor da string da url e tentando converter em int usando o Atoi(), se nao conseguir converter
	// ou o valor for menor do que 1, retorna um 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific Snippet with ID %d", id)
}

// Adicionando func com handler pra criar uma anotacao
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Usando r.method pra checar se a request esta usando o metodo que eu quero
	if r.Method != http.MethodPost {
		// Informando ao user o tipo de metodo permitido, primeiro param é o nome do header e segundo o método
		w.Header().Set("Allow", http.MethodPost)
		// Se nao estiver, retorno um 405 usando http.StatusMethodNotAllowed
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create new Snippet"))
}