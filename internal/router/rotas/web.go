package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/estudeIA-golang/internal/controllers/web"
)

var rotasWeb = []Rota{
	{
		URI:                "/login",
		Metodo:             http.MethodGet,
		Funcao:             controllers.LoginWeb,
		RequerAutenticacao: false,
	},

	{
		URI:                "/home",
		Metodo:             http.MethodGet,
		Funcao:             controllers.HomeWeb,
		RequerAutenticacao: false,
	},
}
