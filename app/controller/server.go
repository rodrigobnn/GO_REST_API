package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	//Rotas / Endpoints
	router.HandleFunc("/api/contas", GetContas).Methods("GET")
	router.HandleFunc("/api/contas/{conta}/{agencia}", GetConta).Methods("GET")
	router.HandleFunc("/api/contas", CreateConta).Methods("POST")
	router.HandleFunc("/api/contas", UpdateConta).Methods("PUT")
	router.HandleFunc("/api/contas/{conta}/{agencia}", DeleteConta).Methods("DELETE")

	router.HandleFunc("/api/cartoes", getCartoes).Methods("GET")
	router.HandleFunc("/api/cartoes/{id}", getCartao).Methods("GET")
	router.HandleFunc("/api/cartoes", createCartao).Methods("POST")
	router.HandleFunc("/api/cartoes", updateCartao).Methods("PUT")
	router.HandleFunc("/api/cartoes/{id}", deleteCartao).Methods("DELETE")
	router.HandleFunc("/api/cartoes/saldo", updateSaldo).Methods("PUT")
	router.HandleFunc("/api/cartoes/saldo/{id}", getSaldo).Methods("GET")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
