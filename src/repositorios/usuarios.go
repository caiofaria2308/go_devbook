package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

//Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioUsuarios cria um repositório de usuários
func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar insere um usuário no banco de dados e retorna o ID
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}
	ultimoID, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoID), nil

}

//Buscar trará todos os usuários de acordo com o filtro
func (respositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%
	linhas, erro := respositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick,
		nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriandoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
