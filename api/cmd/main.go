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
	r := router.CarregarRotas()

	log.Printf("Rodando na porta: %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
