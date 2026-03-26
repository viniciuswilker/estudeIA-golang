package router

import (
	"github.com/gorilla/mux"
	"github.com/viniciuswilker/estudeIA-golang/internal/router/rotas"
)

func CarregarRotas() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
