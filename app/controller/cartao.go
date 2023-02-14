package controller

import (
	"b2/model"
	"b2/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "b2/docs"

	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger"
)

// getCartoes godoc
// @Summary Busca todos os cartões
// @Description Busca todos os cartões
// @Tags cartao
// @Produce  json
// @Success 200
// @Router /api/cartoes [get]
func getCartoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := model.GetCartoes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(posts)
	}
}

// getCartao godoc
// @Summary Busca um cartão em específico
// @Description Busca um cartão em específico
// @Tags cartao
// @Param id path int true "Número do Cartão"
// @Produce  json
// @Success 200
// @Router /api/cartoes/{id} [get]
func getCartao(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	cartao_numero, erro := strconv.ParseInt(params["id"], 10, 64)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número do cartao é inválido, forneça somente números")
		return
	}

	if !util.ValidadorLuhn(cartao_numero, false, false, false) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número do cartao é inválido")
		return
	}

	post, erro := model.GetCartao(cartao_numero)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	if post.Mes_ano_venc != "" {
		json.NewEncoder(w).Encode(post)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Cartão não encontrado")
}

// createCartao godoc
// @Summary Cria um cartão
// @Description Cria um cartão
// @Tags cartao
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/cartoes [post]
func createCartao(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var cartao model.Cartao
	erro := decoder.Decode(&cartao)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	//verifica se todos os campos foram preenchidos
	//assumindo que limite restante seja igual limite na crição, ativo = true e bloqueado = false
	if cartao.Agencia_numero == 0 || cartao.Cartao_numero == 0 || cartao.Conta_numero == 0 || cartao.Mes_ano_venc == "" || cartao.Cvc == 0 || cartao.Limite == 0 {
		json.NewEncoder(w).Encode("Todos os campos precisam ser preenchidos")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//verifica se o cartão é valido usando Luhn
	if !util.ValidadorLuhn(cartao.Cartao_numero, false, false, false) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número do cartao é inválido")
		return
	}

	_, erro = model.CreateCartao(cartao)

	if erro != nil {

		if strings.Contains(erro.Error(), "duplicate key") {
			json.NewEncoder(w).Encode("Esse cartão já existe nessa agência")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if strings.Contains(erro.Error(), "violates foreign key") {
			json.NewEncoder(w).Encode("Essa conta ou agência não existem")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cartao)
	}
}

// updateCartao godoc
// @Summary Atualiza dados de um cartão
// @Description Atualiza dados de um cartão
// @Tags cartao
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/cartoes [put]
func updateCartao(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var cartao model.Cartao
	erro := decoder.Decode(&cartao)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	//verifica se todos os campos foram preenchidos
	//assumindo que limite restante seja igual limite na crição, ativo = true e bloqueado = false
	if cartao.Agencia_numero == 0 || cartao.Cartao_numero == 0 || cartao.Conta_numero == 0 || cartao.Mes_ano_venc == "" || cartao.Cvc == 0 || cartao.Limite == 0 {
		json.NewEncoder(w).Encode("Todos os campos precisam ser preenchidos")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//verifica se o cartão é valido usando Luhn
	if !util.ValidadorLuhn(cartao.Cartao_numero, false, false, false) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número do cartao é inválido")
		return
	}

	result, erro := model.UpdateCartao(cartao)

	if result == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Cartão não encontrado")
		return
	}

	if erro != nil {

		if strings.Contains(erro.Error(), "violates foreign key") {
			json.NewEncoder(w).Encode("Essa conta ou agência não existem")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cartao)
	}
}

// deleteCartao godoc
// @Summary Deleta um cartão
// @Description Deleta um cartão
// @Tags cartao
// @Param id path int true "Número do Cartão"
// @Produce  json
// @Success 200
// @Router /api/cartoes/{id} [delete]
func deleteCartao(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	cartao_numero, erro := strconv.ParseInt(params["id"], 10, 64)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número do cartao é inválido, forneça somente números")
		return
	}

	if !util.ValidadorLuhn(cartao_numero, false, false, false) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número do cartao é inválido")
		return
	}

	result, erro := model.DeleteCartao(cartao_numero)

	if result == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Cartão não encontrado")
		return
	}

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

// updateSaldo godoc
// @Summary Deleta um cartão
// @Description Deleta um cartão
// @Tags cartao
// @Param cartao_numero path int true "Número do Cartão"
// @Param limite_disponivel path number true "Limite Disponível"
// @Produce  json
// @Success 200
// @Router /api/cartoes/saldo [put]
func updateSaldo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var cartao model.Cartao
	erro := decoder.Decode(&cartao)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	cartao_old, erro := model.GetCartao(cartao.Cartao_numero)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	if cartao_old.Mes_ano_venc == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Cartão não encontrado")
		return
	}

	if cartao.Limite_disponivel > cartao_old.Limite {
		s := fmt.Sprintf("%0.2f", cartao_old.Limite)

		json.NewEncoder(w).Encode("O novo limite disponível não pode ser maior do que: R$ " + s)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	_, erro = model.UpdateSaldo(cartao.Cartao_numero, cartao.Limite_disponivel)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cartao.Limite_disponivel)
	}
}

// getSaldo godoc
// @Summary Deleta uam conta
// @Description Deleta uam conta
// @Tags cartao
// @Param id path int true "Número do Cartão"
// @Produce  json
// @Success 200
// @Router /api/cartoes/saldo/{id} [get]
func getSaldo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	cartao_numero, erro := strconv.ParseInt(params["id"], 10, 64)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número do cartao é inválido, forneça somente números")
		return
	}

	saldo, erro := model.GetSaldo(cartao_numero)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	if saldo != 0 {
		json.NewEncoder(w).Encode(saldo)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Cartão não encontrado")
}
