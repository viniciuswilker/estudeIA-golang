package rotas

import "net/http"

var rotasAdmin = []Rota{
	{
		URI:                "/admin/painel",
		Metodo:             []string{http.MethodGet, http.MethodPost},
		Funcao:             func(w http.ResponseWriter, r *http.Request) {},
		RequerAutenticacao: false,
		TiposPermitidos:    []string{"A"},
	},
}
