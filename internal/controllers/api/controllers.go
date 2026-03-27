package controllers

import "net/http"

func CadastroUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CRIANDO USUARIO"))

}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("LISTANDO USUARIOS"))

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeletarUsuario "))

}


func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizar Usuario"))

}
