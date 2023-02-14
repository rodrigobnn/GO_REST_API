package model

type Cartao struct {
	Cartao_numero     int64   `json:"cartao_numero"`
	Mes_ano_venc      string  `json:"mes_ano_venc"` //mm/aa
	Cvc               int     `json:"cvc"`
	Conta_numero      int     `json:"conta_numero"`
	Agencia_numero    int     `json:"agencia_numero"` //F = pessoa fisica / J = pessoa juridica
	Limite            float32 `json:"limite"`         //CPF ou CNPJ a depender do tipo da conta
	Limite_disponivel float32 `json:"limite_disponivel"`
	Ativo             bool    `json:"ativo"`
	Bloqueado         bool    `json:"bloqueado"`
}

// Recupera todos cartões
func GetCartoes() ([]Cartao, error) {
	var cartoes []Cartao

	query := `select cartao_numero, mes_ano_venc, cvc, conta_numero, agencia_numero, limite, limite_disponivel, ativo, bloqueado from "cartoes";`

	rows, erro := db.Query(query)
	if erro != nil {
		return cartoes, erro
	}

	for rows.Next() {
		var cvc, conta_numero, agencia_numero int
		var mes_ano_venc string
		var ativo, bloqueado bool
		var limite, limite_disponivel float32
		var cartao_numero int64

		erro := rows.Scan(&cartao_numero, &mes_ano_venc, &cvc, &conta_numero, &agencia_numero, &limite, &limite_disponivel, &ativo, &bloqueado)

		if erro != nil {
			return cartoes, erro
		}

		cartao := Cartao{
			Cartao_numero:     cartao_numero,
			Mes_ano_venc:      mes_ano_venc,
			Cvc:               cvc,
			Conta_numero:      conta_numero,
			Agencia_numero:    agencia_numero,
			Limite:            limite,
			Limite_disponivel: limite_disponivel,
			Ativo:             ativo,
			Bloqueado:         bloqueado,
		}

		cartoes = append(cartoes, cartao)

	}
	return cartoes, nil
}

// Recupera um ou mais cartões
func GetCartao(cartao_numero int64) (Cartao, error) {
	var cartao Cartao

	query := `select cartao_numero, mes_ano_venc, cvc, conta_numero, agencia_numero, limite, limite_disponivel, ativo, bloqueado from cartoes where cartao_numero=$1;`
	row, erro := db.Query(query, cartao_numero)

	if erro != nil {
		return cartao, erro
	}

	defer row.Close()

	if row.Next() {
		var cvc, conta_numero, agencia_numero int
		var mes_ano_venc string
		var ativo, bloqueado bool
		var limite, limite_disponivel float32
		var cartao_numero int64

		erro := row.Scan(&cartao_numero, &mes_ano_venc, &cvc, &conta_numero, &agencia_numero, &limite, &limite_disponivel, &ativo, &bloqueado)

		if erro != nil {
			return cartao, erro
		}

		cartao = Cartao{
			Cartao_numero:     cartao_numero,
			Mes_ano_venc:      mes_ano_venc,
			Cvc:               cvc,
			Conta_numero:      conta_numero,
			Agencia_numero:    agencia_numero,
			Limite:            limite,
			Limite_disponivel: limite_disponivel,
			Ativo:             ativo,
			Bloqueado:         bloqueado,
		}
	}

	return cartao, nil
}

// Cria um cartão
func CreateCartao(cartao Cartao) (int64, error) {
	query := `insert into cartoes(cartao_numero, mes_ano_venc, cvc, conta_numero, agencia_numero, limite, limite_disponivel, ativo, bloqueado) values($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	result, erro := db.Exec(query, cartao.Cartao_numero, cartao.Mes_ano_venc, cartao.Cvc, cartao.Conta_numero, cartao.Agencia_numero, cartao.Limite, cartao.Limite_disponivel, cartao.Ativo, cartao.Bloqueado)

	if erro != nil {
		return -1, erro
	}

	return result.RowsAffected()
}

// Atualiza um cartão
func UpdateCartao(cartao Cartao) (int64, error) {
	query := `update cartoes set cartao_numero=$1, mes_ano_venc=$2, cvc=$3, conta_numero=$4, agencia_numero=$5, limite=$6, limite_disponivel=$7, ativo=$8, bloqueado=$9 where cartao_numero=$1;`

	result, erro := db.Exec(query, cartao.Cartao_numero, cartao.Mes_ano_venc, cartao.Cvc, cartao.Conta_numero, cartao.Agencia_numero, cartao.Limite, cartao.Limite_disponivel, cartao.Ativo, cartao.Bloqueado)

	if erro != nil {
		return -1, erro
	}
	return result.RowsAffected()
}

// Deleta um cartão
func DeleteCartao(cartao_numero int64) (int64, error) {
	query := `delete from cartoes where cartao_numero=$1;`

	result, erro := db.Exec(query, cartao_numero)

	if erro != nil {
		return -1, erro
	}

	return result.RowsAffected()
}

// Atualiza o saldo diposnível
func UpdateSaldo(cartao_numero int64, new_value float32) (int64, error) {
	query := `update cartoes set limite_disponivel=$1 where cartao_numero=$2;`

	result, erro := db.Exec(query, new_value, cartao_numero)

	if erro != nil {
		return -1, erro
	}

	return result.RowsAffected()
}

// Atualiza o saldo diposnível
func GetSaldo(cartao_numero int64) (float32, error) {
	query := `select limite_disponivel from cartoes where cartao_numero=$1;`

	row, erro := db.Query(query, cartao_numero)

	if erro != nil {
		return -1, erro
	}

	var saldo_diponivel float32
	if row.Next() {
		row.Scan(&saldo_diponivel)
	}

	return saldo_diponivel, nil
}
