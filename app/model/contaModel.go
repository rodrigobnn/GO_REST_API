package model

import "time"

// TableName overrides the table name used by User to `profiles`
func (Conta) TableName() string {
	return "contas"
}

type Conta struct {
	Conta_numero   int       `gorm:"primaryKey;autoIncrement:false" json:"conta_numero"`
	Agencia_numero int       `gorm:"primaryKey;autoIncrement:false" json:"agencia_numero"`
	Titular        string    `json:"titular"`
	Tipo           TipoConta `json:"tipo"`          //F = pessoa fisica / J = pessoa juridica
	Identificador  string    `json:"identificador"` //CPF ou CNPJ a depender do tipo da conta
	Ativa          bool      `json:"ativa"`
	CreatedAt      time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime:false"`
}

// Recupera todas as contas
func GetContas() ([]Conta, error) {
	var contas []Conta

	result := db.Find(&contas)

	if result.Error != nil {
		return contas, result.Error
	}

	return contas, nil
}

// Recupera Uma conta em específico
func GetConta(conta_numero int, agencia_numero int) (Conta, error) {
	var conta Conta

	result := db.Where("conta_numero = ? AND agencia_numero = ?", conta_numero, agencia_numero).First(&conta)

	if result.Error != nil {
		return conta, result.Error
	}

	return conta, nil

}

// Cria uma conta
func CreateConta(conta Conta) (int64, error) {

	result := db.Create(&conta)

	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil

}

// atualiza uma conta
func UpdateConta(conta Conta) (int64, error) {

	result := db.Model(Conta{}).Where("conta_numero = ? AND agencia_numero = ?", conta.Conta_numero, conta.Agencia_numero).Updates(conta)

	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil

}

// deletea uma conta
func DeleteConta(conta_numero int, agencia_numero int) (int64, error) {
	result := db.Where("conta_numero = ? AND agencia_numero = ?", conta_numero, agencia_numero).Delete(&Conta{})

	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil

}
