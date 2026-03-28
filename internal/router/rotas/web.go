package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/estudeIA-golang/internal/controllers/web"
)

var rotasWeb = []Rota{
	{
		URI:                "/login",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.LoginWeb,
		RequerAutenticacao: false,
	},

	{
		URI:                "/home",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.HomeWeb,
		RequerAutenticacao: false,
	},
	{
		URI:                "/cadastro",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao: controllers.CadastroWeb,
		RequerAutenticacao: false,
	},

}
