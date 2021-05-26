package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		//Cadastro de usuários
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutentitacao: false,
	},
	{
		//Busca todos os usuários
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutentitacao: true,
	},
	{
		//Buscar apenas um usuário por id
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutentitacao: true,
	},
	{
		//Atualiza um usuário pelo ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutentitacao: true,
	},
	{
		//Deleta usuário pelo ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutentitacao: true,
	},
}
