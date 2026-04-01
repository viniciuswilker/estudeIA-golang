package rotas

import (
	"net/http"

	controllers "github.com/viniciuswilker/estudeIA-golang/internal/controllers/web"
)

var rotasFazendas = []Rota{
	{
		URI:                "/fazendas",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             controllers.CadastroFazendaWeb,
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},
}
