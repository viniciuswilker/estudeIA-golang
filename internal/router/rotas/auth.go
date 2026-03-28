package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/estudeIA-golang/internal/controllers/api"
)

var rotasAuth = []Rota{
	{
		URI:                "/auth/login",
		Metodo:             []string{http.MethodPost},
		Funcao:             controllers.LoginAPI,
		RequerAutenticacao: false,
	},
	{
		URI:                "/auth/logout",
		Metodo:             []string{http.MethodPost},
		Funcao:             func(w http.ResponseWriter, r *http.Request) {},
		RequerAutenticacao: true,
	},
}
