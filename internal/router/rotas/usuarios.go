package rotas

import "net/http"

var rotasUsuarios = []Rota{
	{
		URI:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("LISTANDO USUARIOS"))
		},
	},
	{
		URI:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("CRIANDO USUARIO"))
		},
	},
}
