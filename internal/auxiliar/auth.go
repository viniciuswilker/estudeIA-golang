package auxiliar

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/viniciuswilker/estudeIA-golang/internal/database"
	"github.com/viniciuswilker/estudeIA-golang/internal/models"
	"github.com/viniciuswilker/estudeIA-golang/internal/repositorios"
)

func ValidaSessao(r *http.Request) (models.Usuario, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return models.Usuario{}, err
	}

	token, err := jwt.Parse(cookie.Value, retornarChaveDeVerificacao)
	if err != nil || !token.Valid {
		return models.Usuario{}, errors.New("token inválido")
	}

	permissoes, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return models.Usuario{}, errors.New("erro ao ler claims")
	}
	usuarioID, _ := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)

	db, err := database.ConectaBanco()
	if err != nil {
		return models.Usuario{}, err
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	return repositorio.BuscarPorID(usuarioID)
}
