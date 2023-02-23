package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type TipoConta string

const (
	F TipoConta = "Pessoa Física"
	J TipoConta = "Pessoa Jurídica"
)

type Tabler interface {
	TableName() string
}

var db *gorm.DB

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

	// // try to open our postgresql connection with our connection info
	// db, err = sql.Open("postgres", connToString(connInfo))

	// dsn := "host=" + os.Getenv("POSTGRES_URL") +
	// 	" user=" + os.Getenv("POSTGRES_USER") +
	// 	" password=" + os.Getenv("POSTGRES_PASSWORD") +
	// 	" dbname=" + os.Getenv("POSTGRES_DB") +
	// 	" port=" + os.Getenv("POSTGRES_PORT") +
	// 	" sslmode=disable"
	db, err = gorm.Open(postgres.Open(connToString(connInfo)), &gorm.Config{})

	if err != nil {
		fmt.Printf("Erro ao conectar no banco de dados: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB aberto\n")
	}

	// check if we can ping our DB
	sqlDB, erro := db.DB()

	if erro != nil {
		log.Fatal(erro)
	}

	err = sqlDB.Ping()
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
