package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {

	api := r.PathPrefix("/api").Subrouter()
	web := r.PathPrefix("/").Subrouter()

	for _, rota := range rotasUsuarios {
		api.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	for _, rota := range rotasWeb {
		web.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
