package rotas

import (
	"api/src/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutentitacao bool
}

//Configurar irá configurar todas rotas no router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {
		fmt.Printf("Criando rota %s\n", rota.URI)
		if rota.RequerAutentitacao {
			r.HandleFunc(rota.URI,
				 middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
				 ).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, 
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}
	}
	return r
}
