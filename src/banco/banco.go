package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Driver MYSQL

//Conectar no banco de dados
func Conectar() (*sql.DB, error) {
	//Carregando configuração
	config.Carregar()

	//Conectando...
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, erro
}
