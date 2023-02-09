package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"b2/model"
	tipos "b2/model"
)

// var contas []tipos.Conta

// Recupera todas as contas
func GetContas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := model.GetContas()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(posts)
	}
}

// Recupera Uma conta em específico
func GetConta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	conta_param, erro := strconv.Atoi(params["conta"])

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número da conta é inválido, forneça somente números")
		return
	}

	agencia_param, err := strconv.Atoi(params["agencia"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número da agência é inválido, forneça somente números")
		return
	}

	post, erro := model.GetConta(conta_param, agencia_param)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	if post.Titular != "" {
		json.NewEncoder(w).Encode(post)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Conta não encontrada")

}

// Cria uma conta
func CreateConta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var conta model.Conta
	erro := decoder.Decode(&conta)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	//verifica se todos os campos foram preenchidos
	if conta.Conta_numero == 0 || conta.Agencia_numero == 0 || conta.Titular == "" || conta.Tipo == "" || conta.Identificador == "" {
		json.NewEncoder(w).Encode("Todos os campos precisam ser preenchidos")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//verifica se forneceu nome e sobrenome
	nome_sobrenome := strings.Split(conta.Titular, " ")

	if len(nome_sobrenome) < 2 {
		json.NewEncoder(w).Encode("Favor fornecer Nome e Sobrenome")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//vefica se cpf está de acordo com a conta e está correto
	if conta.Tipo == tipos.F && len(conta.Identificador) != 14 {
		json.NewEncoder(w).Encode("Favor fornecer CPF com os separadores")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	//vefica se cnpj está de acordo com a conta e está correto
	if conta.Tipo == tipos.J && len(conta.Identificador) != 18 {
		json.NewEncoder(w).Encode("Favor fornecer CPNJ com os separadores")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	_, erro = model.CreateConta(conta)

	if erro != nil {

		if strings.Contains(erro.Error(), "duplicate key") {
			json.NewEncoder(w).Encode("Essa conta já existe nessa agência")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(conta)
	}

}

// atualiza uma conta
func UpdateConta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var conta model.Conta
	erro := decoder.Decode(&conta)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(erro.Error()))
		return
	}

	//verifica se todos os campos foram preenchidos
	if conta.Conta_numero == 0 || conta.Agencia_numero == 0 || conta.Titular == "" || conta.Tipo == "" || conta.Identificador == "" {
		json.NewEncoder(w).Encode("Todos os campos precisam ser preenchidos")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//verifica se forneceu nome e sobrenome
	nome_sobrenome := strings.Split(conta.Titular, " ")

	if len(nome_sobrenome) < 2 {
		json.NewEncoder(w).Encode("Favor fornecer Nome e Sobrenome")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//vefica se cpf está de acordo com a conta e está correto
	if conta.Tipo == tipos.F && len(conta.Identificador) != 14 {
		json.NewEncoder(w).Encode("Favor fornecer CPF com os separadores")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//vefica se cnpj está de acordo com a conta e está correto
	if conta.Tipo == tipos.J && len(conta.Identificador) != 18 {
		json.NewEncoder(w).Encode("Favor fornecer CPNJ com os separadores")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	result, erro := model.UpdateConta(conta)

	fmt.Printf("Results %d \n", result)

	if result == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Conta não encontrada")
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

// deletea uma conta
func DeleteConta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	conta_param, erro := strconv.Atoi(params["conta"])

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número da conta é inválido, forneça somente números")
		return
	}

	agencia_param, erro := strconv.Atoi(params["agencia"])

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("O número da agência é inválido, forneça somente números")
		return
	}

	result, erro := model.DeleteConta(conta_param, agencia_param)

	fmt.Printf("Results %d \n", result)

	if result == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Conta não encontrada")
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
