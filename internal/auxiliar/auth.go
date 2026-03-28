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
		fmt.Println("ERRO SESSÃO: Cookie não encontrado")
		return models.Usuario{}, err
	}

	token, err := jwt.Parse(cookie.Value, retornarChaveDeVerificacao)
	if err != nil {
		fmt.Printf("ERRO SESSÃO: Falha no Parse do JWT: %v\n", err)
		return models.Usuario{}, err
	}

	if !token.Valid {
		fmt.Println("ERRO SESSÃO: Token existe mas o JWT diz que é inválido")
		return models.Usuario{}, errors.New("token inválido")
	}

	permissoes, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("ERRO SESSÃO: Erro ao converter claims")
		return models.Usuario{}, errors.New("erro ao ler claims")
	}

	usuarioID, _ := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
	fmt.Printf("DEBUG SESSÃO: Token OK! Buscando usuário ID %d no banco...\n", usuarioID)

	db, err := database.ConectaBanco()
	if err != nil {
		fmt.Printf("ERRO SESSÃO: Falha na conexão com banco: %v\n", err)
		return models.Usuario{}, err
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		fmt.Printf("ERRO SESSÃO: Usuário %d não encontrado no banco: %v\n", usuarioID, err)
		return models.Usuario{}, err
	}

	fmt.Printf("DEBUG SESSÃO: Sucesso! Usuário %s encontrado.\n", usuario.Nome)
	return usuario, nil
}
