package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

//TODO
/*
* usar sync.Mutex para fazer o travamento para edição
 */
type TipoConta string

const (
	F TipoConta = "Pessoa Física"
	J TipoConta = "Pessoa Jurídica"
)

// declare a db object, where we can use throughout the model package
// so in blog.go, we have access to this object
var db *sql.DB

// a struct to hold all the db connection information
type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("Erro ao carregar arquivo .env: %s\n", err.Error())
		return
	}

	connInfo := connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	// try to open our postgresql connection with our connection info
	db, err = sql.Open("postgres", connToString(connInfo))
	if err != nil {
		fmt.Printf("Erro ao conectar no banco de dados: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB aberto\n")
	}

	// check if we can ping our DB
	err = db.Ping()
	if err != nil {
		fmt.Printf("Erro, não foi possível dar ping no BD: %s\n", err.Error())
		return
	} else {
		fmt.Printf("BD pingado com sucesso\n")
	}
}

// Take our connection struct and convert to a string for our db connection info
func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)

}
