package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

//Logger escreve informações da requisição no terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(rw, r)
	}
}

//Autenticar  verifica se o usuário fazendo requisição está autenticado
func Autenticar(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunction(w, r)
	}
}
