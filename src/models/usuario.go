package models

import (
	"errors"
	"strings"
	"time"
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

//Preparar vai chamar os metodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("'nome' é um campo obrigatório")
	} else if usuario.Nick == "" {
		return errors.New("'nick' é um campo obrigatório")
	} else if usuario.Email == "" {
		return errors.New("'email' é um campo obrigatório")
	} else if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("'senha' é um campo obrigatório")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
