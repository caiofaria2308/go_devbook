package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

//Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CriandoEm time.Time `json:"criandoEm,omitempty"`
}

//UsuarioToken representa um token gerado no login
type UsuarioToken struct {
	Token string `json:"token,omitempty"`
}

//Preparar vai chamar os metodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("'nome' é um campo obrigatório")
	} else if usuario.Nick == "" {
		return errors.New("'nick' é um campo obrigatório")
	} else if usuario.Email == "" {
		return errors.New("'email' é um campo obrigatório")
	} else if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("'email' inválido")
	} else if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("'senha' é um campo obrigatório")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)

	}
	return nil
}
