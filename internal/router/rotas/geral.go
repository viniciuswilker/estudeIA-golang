package rotas

import "net/http"

var rotasGerais = []Rota{
	{
		URI:    "/ping",
		Metodo: []string{http.MethodGet},
		Funcao: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		},
		RequerAutenticacao: false,
		TiposPermitidos:    nil,
	},
}
