package model

type Conta struct {
	Conta_numero   int       `json:"conta_numero"`
	Agencia_numero int       `json:"agencia_numero"`
	Titular        string    `json:"titular"`
	Tipo           TipoConta `json:"tipo"`          //F = pessoa fisica / J = pessoa juridica
	Identificador  string    `json:"identificador"` //CPF ou CNPJ a depender do tipo da conta
	Ativa          bool      `json:"ativa"`
}

//TODO retornar os sqlresults e conferir a execução das operações

// Recupera todas as contas
func GetContas() ([]Conta, error) {
	var contas []Conta

	query := `select conta_numero, agencia_numero, titular, tipo, identificador, ativa from contas;`

	rows, erro := db.Query(query)
	if erro != nil {
		return contas, erro
	}

	defer rows.Close()

	for rows.Next() {
		var conta_numero, agencia_numero int
		var titular, identificador string
		var ativa bool
		var tipo TipoConta

		err := rows.Scan(&conta_numero, &agencia_numero, &titular, &tipo, &identificador, &ativa)
		if err != nil {
			return contas, err
		}

		conta := Conta{
			Conta_numero:   conta_numero,
			Agencia_numero: agencia_numero,
			Titular:        titular,
			Tipo:           tipo,
			Identificador:  identificador,
			Ativa:          ativa,
		}

		contas = append(contas, conta)

	}
	return contas, nil
}

// Recupera Uma conta em específico
func GetConta(conta_numero int, agencia_numero int) (Conta, error) {
	var conta Conta

	query := `select conta_numero, agencia_numero, titular, tipo, identificador, ativa from contas where conta_numero=$1 and agencia_numero=$2;`
	row, erro := db.Query(query, conta_numero, agencia_numero)

	if erro != nil {
		return conta, erro
	}

	defer row.Close()

	if row.Next() {
		var conta_numero, agencia_numero int
		var titular, identificador string
		var ativa bool
		var tipo TipoConta

		erro := row.Scan(&conta_numero, &agencia_numero, &titular, &tipo, &identificador, &ativa)
		if erro != nil {
			return conta, erro
		}

		conta = Conta{
			Conta_numero:   conta_numero,
			Agencia_numero: agencia_numero,
			Titular:        titular,
			Tipo:           tipo,
			Identificador:  identificador,
			Ativa:          ativa,
		}
	}

	return conta, nil

}

// Cria uma conta
func CreateConta(conta Conta) (int64, error) {
	query := `insert into contas(conta_numero, agencia_numero, titular, tipo, identificador, ativa) values($1, $2, $3, $4, $5, $6);`

	result, erro := db.Exec(query, conta.Conta_numero, conta.Agencia_numero, conta.Titular, conta.Tipo, conta.Identificador, conta.Ativa)

	if erro != nil {
		return -1, erro
	}

	return result.RowsAffected()

}

// atualiza uma conta
func UpdateConta(conta Conta) (int64, error) {
	query := `update contas set conta_numero=$1, agencia_numero=$2, titular=$3, tipo=$4, identificador=$5, ativa=$6 where conta_numero=$1 and agencia_numero=$2;`

	result, erro := db.Exec(query, conta.Conta_numero, conta.Agencia_numero, conta.Titular, conta.Tipo, conta.Identificador, conta.Ativa)
	if erro != nil {
		return -1, erro
	}
	return result.RowsAffected()
}

// deletea uma conta
func DeleteConta(conta_numero int, agencia_numero int) (int64, error) {
	query := `delete from contas where conta_numero=$1 and agencia_numero=$2;`
	result, erro := db.Exec(query, conta_numero, agencia_numero)
	if erro != nil {
		return -1, erro
	}
	return result.RowsAffected()

}
