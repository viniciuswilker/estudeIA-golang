package middlewares

import (
	"net/http"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auxiliar.ValidarToken(r); erro != nil {
			http.Error(w, "Não autorizado: "+erro.Error(), http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func VerificarPermissao(next http.HandlerFunc, tiposPermitidos []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tipoNoToken, err := auxiliar.ExtrairTipoUsuario(r)
		if err != nil {
			http.Error(w, "Erro ao validar permissão", http.StatusForbidden)
			return
		}

		permitido := false
		for _, t := range tiposPermitidos {
			if t == tipoNoToken {
				permitido = true
				break
			}
		}

		if !permitido {
			http.Error(w, "Acesso negado: nível insuficiente", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
