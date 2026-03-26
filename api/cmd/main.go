package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/viniciuswilker/estudeIA-golang/internal/config"
	"github.com/viniciuswilker/estudeIA-golang/internal/router"
)

func main() {
	config.CarregarConfigs()

	log.Printf("Rodando na porta: %d", config.Porta)

	r := router.CarregarRotas()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
