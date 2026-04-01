package models

import "time"

type Fazenda struct {
	ID            uint64    `json:"id,omitempty"`
	Nome          string    `json:"nome,omitempty"`
	Endereco      string    `json:"endereco,omitempty"`
	CodigoFazenda string    `json:"codigo_fazenda,omitempty"`
	DataCriacao   time.Time `json:"data_criacao,omitempty"`
}
