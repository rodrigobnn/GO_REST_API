package model

// TableName overrides the table name used by User to `profiles`
func (Cartao) TableName() string {
	return "cartoes"
}

type Cartao struct {
	Cartao_numero     int64   `gorm:"primaryKey;autoIncrement:false" json:"cartao_numero"`
	Mes_ano_venc      string  `json:"mes_ano_venc"` //mm/aa
	Cvc               int     `json:"cvc"`
	Conta_numero      int     `gorm:"primaryKey;autoIncrement:false" json:"conta_numero"`
	Agencia_numero    int     `gorm:"primaryKey;autoIncrement:false" json:"agencia_numero"` //F = pessoa fisica / J = pessoa juridica
	Limite            float32 `json:"limite"`                                               //CPF ou CNPJ a depender do tipo da conta
	Limite_disponivel float32 `json:"limite_disponivel"`
	Ativo             bool    `json:"ativo"`
	Bloqueado         bool    `json:"bloqueado"`
}

// Recupera todos cartões
func GetCartoes() ([]Cartao, error) {
	var cartoes []Cartao

	result := db.Find(&cartoes)

	if result.Error != nil {
		return cartoes, result.Error
	}

	return cartoes, nil
}

// Recupera um ou mais cartões
func GetCartao(cartao_numero int64) (Cartao, error) {
	var cartao Cartao

	result := db.Where("cartao_numero = ?", cartao_numero).First(&cartao)

	if result.Error != nil {
		return cartao, result.Error
	}

	return cartao, nil
}

// Cria um cartão
func CreateCartao(cartao Cartao) (int64, error) {

	result := db.Create(&cartao)

	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil
}

// Atualiza um cartão
func UpdateCartao(cartao Cartao) (int64, error) {

	result := db.Model(Cartao{}).Where("cartao_numero = ?", cartao.Cartao_numero).Updates(cartao)

	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil
}

// Deleta um cartão
func DeleteCartao(cartao_numero int64) (int64, error) {

	result := db.Where("cartao_numero = ?", cartao_numero).Delete(&Cartao{})

	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil
}

// Atualiza o saldo diposnível
func UpdateSaldo(cartao_numero int64, new_value float32) (int64, error) {

	result := db.Model(&Cartao{}).Where("cartao_numero = ?", cartao_numero).Update("limite_disponivel", new_value)

	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil

}

// Atualiza o saldo diposnível
func GetSaldo(cartao_numero int64) (float32, error) {

	var saldo_diponivel float32

	result := db.Model(&Cartao{}).Select("limite_disponivel").Where("cartao_numero = ?", cartao_numero).Scan(&saldo_diponivel)

	if result.Error != nil {
		return -1, result.Error
	}

	return saldo_diponivel, nil
}
