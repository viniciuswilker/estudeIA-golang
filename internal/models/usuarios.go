package models

import "time"

type Usuario struct {
	ID          uint64    `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	Nome        string    `json:"nome,omitempty"`
	Sobrenome   string    `json:"sobrenome,omitempty"`
	Email       string    `json:"email,omitempty"`
	Senha       string    `json:"senha,omitempty"`
	TipoUsuario string    `json:"tipo_usuario,omitempty"`
	CriadoEm    time.Time `json:"criadoEm,omitempty"`
}
