package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/estudeIA-golang/internal/controllers/api"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.ListarUsuarios,
		RequerAutenticacao: false,
		TiposPermitidos:    []string{"A", "U"},
	},
	{
		URI:                "/usuarios",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.CadastroUsuario,
		RequerAutenticacao: false,
		TiposPermitidos:    []string{"A", "U"},
	},
	{
		URI:                "/usuarios",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
		TiposPermitidos:    []string{"A", "U"},
	},
	{
		URI:                "/usuarios",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
		TiposPermitidos:    []string{"A", "U"},
	},
}
