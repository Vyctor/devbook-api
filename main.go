package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvs()
	fmt.Printf("Rodando API na port %s\n", config.AppPort)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.AppPort), r))
}
