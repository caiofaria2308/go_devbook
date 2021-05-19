package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//iniciando router
	fmt.Println("Iniciando api...")
	//Carregando configuração
	config.Carregar()
	//gerando rotas
	r := router.Gerar()
	//disponibilizando api na porta configurada
	fmt.Printf("Rodando API na porta :%d \n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
