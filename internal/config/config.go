package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	Porta       = 0
	StringBanco = ""
	SecretKey   []byte
)
var Store = sessions.NewCookieStore([]byte("sua-chave-secreta-muito-segura"))

func CarregarConfigs() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 8000
	}

	StringBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
