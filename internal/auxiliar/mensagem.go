package auxiliar

import (
	"net/http"

	"github.com/viniciuswilker/estudeIA-golang/internal/config"
)

func MensagemFlash(w http.ResponseWriter, r *http.Request, nome, mensagem, tipo string, url string) {
	session, _ := config.Store.Get(r, nome)
	session.AddFlash(mensagem, tipo)
	session.Save(r, w)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
