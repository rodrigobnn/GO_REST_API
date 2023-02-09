package model

type Cartao struct {
	Cartao_numero     int64   `json:"cartao_numero"`  //TODO validar mastercard
	Mes_vencimento    int     `json:"mes_vencimento"` //mm/
	Ano_vencimento    int     `json:"ano_vencimento"` //aa
	Cvc               int     `json:"cvc"`
	Conta_numero      int     `json:"conta_numero"`
	Agencia_numero    int     `json:"agencia_numero"` //F = pessoa fisica / J = pessoa juridica
	Limite            float32 `json:"limite"`         //CPF ou CNPJ a depender do tipo da conta
	Limite_disponivel float32 `json:"limite_disponivel"`
	Ativo             bool    `json:"ativo"`
	Bloqueado         bool    `json:"bloqueado"`
}

//TODO IMPLEMETAR
