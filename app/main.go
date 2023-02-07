package main

import (
	"b2/controller"
	"b2/model"
	"fmt"
)

// /*
// * usar sync.Mutex para fazer o travamento para edição
//  */
// type TipoConta string

// const (
// 	F TipoConta = "Pessoa Física"
// 	J TipoConta = "Pessoa Jurídica"
// )

// type Conta struct {
// 	Conta_numero   int       `json:"conta_numero"`
// 	Agencia_numero int       `json:"agencia_numero"`
// 	Titular        string    `json:"titular"`
// 	Tipo           TipoConta `json:"tipo"`          //F = pessoa fisica / J = pessoa juridica
// 	Identificador  string    `json:"identificador"` //CPF ou CNPJ a depender do tipo da conta
// 	Ativa          bool      `json:"ativa"`
// }

// type Cartao struct {
// 	Cartao_numero     int     `json:"cartao_numero"` //TODO validar mastercard
// 	Cvc               int     `json:"cvc"`
// 	Conta_numero      int     `json:"conta_numero"`
// 	Agencia_numero    int     `json:"agencia_numero"` //F = pessoa fisica / J = pessoa juridica
// 	Limite            float32 `json:"limite"`         //CPF ou CNPJ a depender do tipo da conta
// 	Limite_disponivel float32 `json:"limite_disponivel"`
// 	Ativo             bool    `json:"ativo"`
// 	Bloqueado         bool    `json:"bloqueado"`
// }

// Valor de inicialização
// var contas []Conta
// var cartoes []Cartao

// // Recupera todas as contas
// func getContas(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(contas)
// }

// // Recupera Uma conta em específico
// func getConta(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	conta_param, err := strconv.Atoi(params["conta"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode("O número da conta é inválido, forneça somente números")
// 		return
// 	}

// 	agencia_param, err := strconv.Atoi(params["agencia"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode("O número da agência é inválido, forneça somente números")
// 		return
// 	}

// 	for _, item := range contas {
// 		if item.Conta_numero == conta_param && item.Agencia_numero == agencia_param {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}

// 	w.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(w).Encode("Conta não encontrada")

// }

// // Cria uma conta
// func createConta(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var conta Conta
// 	json.NewDecoder(r.Body).Decode(&conta)

// 	//verifica se todos os campos foram preenchidos
// 	if conta.Conta_numero == 0 || conta.Agencia_numero == 0 || conta.Titular == "" || conta.Tipo == "" || conta.Identificador == "" {
// 		json.NewEncoder(w).Encode("Todos os campos precisam ser preenchidos")
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	//verifica se forneceu nome e sobrenome
// 	nome_sobrenome := strings.Split(conta.Titular, " ")

// 	if len(nome_sobrenome) < 2 {
// 		json.NewEncoder(w).Encode("Favor fornecer Nome e Sobrenome")
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	//vefica se cpf está de acordo com a conta e está correto
// 	if conta.Tipo == F && len(conta.Identificador) != 14 {
// 		json.NewEncoder(w).Encode("Favor fornecer CPF com os separadores")
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	//vefica se cnpj está de acordo com a conta e está correto
// 	if conta.Tipo == J && len(conta.Identificador) != 18 {
// 		json.NewEncoder(w).Encode("Favor fornecer CPNJ com os separadores")
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	idx := slices.IndexFunc(contas, func(c Conta) bool {
// 		return c.Conta_numero == conta.Conta_numero && c.Agencia_numero == conta.Agencia_numero
// 	})

// 	if idx != -1 {
// 		json.NewEncoder(w).Encode("Essa conta já existe nessa agência")
// 		w.WriteHeader(http.StatusForbidden)
// 		return
// 	}

// 	contas = append(contas, conta)
// 	json.NewEncoder(w).Encode(conta)

// }

// // atualiza uma conta
// func updateConta(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	conta_param, err := strconv.Atoi(params["conta"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode("O número da conta é inválido, forneça somente números")
// 		return
// 	}

// 	agencia_param, err := strconv.Atoi(params["agencia"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode("O número da agência é inválido, forneça somente números")
// 		return
// 	}

// 	idx := slices.IndexFunc(contas, func(c Conta) bool {
// 		return c.Conta_numero == conta_param && c.Agencia_numero == agencia_param
// 	})

// 	if idx != -1 {
// 		var conta Conta
// 		json.NewDecoder(r.Body).Decode(&conta)

// 		contas[idx] = conta
// 		json.NewEncoder(w).Encode(contas[idx])
// 		return
// 	} else {
// 		json.NewEncoder(w).Encode("Conta não encontrad")
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// }

// // deletea uma conta
// func deleteConta(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	for i, item := range contas {
// 		if strconv.Itoa(item.Conta_numero) == params["conta"] && strconv.Itoa(item.Agencia_numero) == params["agencia"] {
// 			contas = append(contas[:i], contas[i+1:]...)
// 			json.NewEncoder(w).Encode(contas)
// 			return
// 		}
// 	}

// 	json.NewEncoder(w).Encode("Conta não encontrad")
// 	w.WriteHeader(http.StatusNotFound)
// 	return

// }

// // Recupera todos cartões
// func getCartoes(w http.ResponseWriter, r *http.Request) {

// }

// // Recupera um ou mais cartões
// func getCartao(w http.ResponseWriter, r *http.Request) {

// }

// // Cria um cartão
// func createCartao(w http.ResponseWriter, r *http.Request) {

// }

// // Atualiza um cartão
// func updateCartao(w http.ResponseWriter, r *http.Request) {

// }

// // Deleta um cartão
// func deleteCartao(w http.ResponseWriter, r *http.Request) {

// }

// // Atualiza o saldo diposnível
// func updateSaldo(w http.ResponseWriter, r *http.Request) {

// }

// type Conta struct {
// 	Conta_numero   int       `json:"conta_numero"`
// 	Agencia_numero int       `json:"agencia_numero"`
// 	Titular        string    `json:"titular"`
// 	Tipo           TipoConta `json:"tipo"`         //F = pessoa fisica / J = pessoa juridica
// 	Identificador  string    `json:"idetificador"` //CPF ou CNPJ a depender do tipo da conta
// 	Ativa          bool      `json:"ativa"`
// }

// type Cartao struct {
// 	Cartao_numero     int     `json:"cartao_numero"` //TODO validar mastercard
// 	Cvc               int     `json:"cvc"`
// 	Conta_numero      int     `json:"conta_numero"`
// 	Agencia_numero    int     `json:"agencia_numero"` //F = pessoa fisica / J = pessoa juridica
// 	Limite            float32 `json:"limite"`         //CPF ou CNPJ a depender do tipo da conta
// 	Limite_disponivel float32 `json:"limite_disponivel"`
// 	Ativo             bool    `json:"ativo"`
// 	Bloqueado         bool    `json:"bloqueado"`

func main() {
	// //Inicia Router
	// router := mux.NewRouter()

	model.Init()
	controller.Start()

	fmt.Println("Iniciou")

	// //Mock Data
	// contas = append(contas, Conta{Conta_numero: 12345, Agencia_numero: 1234, Titular: "Rodrigo Barbosa", Tipo: F, Identificador: "067.757.446-09", Ativa: true})
	// contas = append(contas, Conta{Conta_numero: 23456, Agencia_numero: 2345, Titular: "Thais Helena", Tipo: F, Identificador: "084.155.596-94", Ativa: true})
	// contas = append(contas, Conta{Conta_numero: 34567, Agencia_numero: 3456, Titular: "BnnCode", Tipo: J, Identificador: "60.813.719/0001-73", Ativa: false})

	// cartoes = append(cartoes, Cartao{Cartao_numero: 1234567898765432, Cvc: 597, Conta_numero: 12345, Agencia_numero: 1234, Limite: 1000.00, Limite_disponivel: 500.00, Ativo: true, Bloqueado: false})
	// cartoes = append(cartoes, Cartao{Cartao_numero: 1111222233334444, Cvc: 011, Conta_numero: 12345, Agencia_numero: 1234, Limite: 2000.00, Limite_disponivel: 1500.00, Ativo: true, Bloqueado: false})
	// cartoes = append(cartoes, Cartao{Cartao_numero: 9999888877776666, Cvc: 666, Conta_numero: 23456, Agencia_numero: 2345, Limite: 500.00, Limite_disponivel: 10.00, Ativo: false, Bloqueado: false})
	// cartoes = append(cartoes, Cartao{Cartao_numero: 0011002200330044, Cvc: 001, Conta_numero: 34567, Agencia_numero: 3456, Limite: 12000.00, Limite_disponivel: 2500.00, Ativo: true, Bloqueado: false})

	// //Rotas / Endpoints
	// router.HandleFunc("/api/contas", getContas).Methods("GET")
	// router.HandleFunc("/api/contas/{conta}/{agencia}", getConta).Methods("GET")
	// router.HandleFunc("/api/contas", createConta).Methods("POST")
	// router.HandleFunc("/api/contas/{conta}/{agencia}", updateConta).Methods("PUT")
	// router.HandleFunc("/api/contas/{conta}/{agencia}", deleteConta).Methods("DELETE")

	// router.HandleFunc("/api/cartao", getCartoes).Methods("GET")
	// router.HandleFunc("/api/cartao/{id}", getCartao).Methods("GET")
	// router.HandleFunc("/api/cartao", createCartao).Methods("POST")
	// router.HandleFunc("/api/cartao/{id}", updateCartao).Methods("PUT")
	// router.HandleFunc("/api/cartao/{id}", deleteCartao).Methods("DELETE")
	// router.HandleFunc("/api/cartao/{id}", updateSaldo).Methods("PUT")

	//inicia servidor
	// log.Fatal(http.ListenAndServe(":8080", router))

	// conta := Conta{1, 2, "Rodrigo Barbosa", TipoConta(F), "06775744609", true}
	// fmt.Println(conta)

}
