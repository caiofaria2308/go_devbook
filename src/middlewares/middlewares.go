package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(rw, r)
	}
}

//Autenticar  verifica se o usuário fazendo requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando token ...")
		next(w, r)
	}
}