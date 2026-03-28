package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/estudeIA-golang/internal/controllers/web"
)

var rotasWeb = []Rota{
	{
		URI:                "/login",
		Metodo:             []string{http.MethodGet},
		Funcao:             controllers.LoginWeb,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},

	{
		URI:                "/home",
		Metodo:             []string{http.MethodGet},
		Funcao:             controllers.HomeWeb,
		RequerAutenticacao: true,
		TiposPermitidos:    []string{"A", "U"},
	},
	{
		URI:                "/cadastro",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.CadastroWeb,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},
}
