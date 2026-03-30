package router

import (
	"github.com/gorilla/mux"
	"github.com/viniciuswilker/estudeIA-golang/internal/middlewares"
	"github.com/viniciuswilker/estudeIA-golang/internal/router/rotas"
)

func CarregarRotas() *mux.Router {
	r := mux.NewRouter()

	r = rotas.Configurar(r)

	r.Use(middlewares.Logger)

	return r
}
