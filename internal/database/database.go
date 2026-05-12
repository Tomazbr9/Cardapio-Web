package database

import (
	"cw/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado. Usando variáveis de ambiente do sistema.")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port, sslmode)

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Erro fatal ao conectar ao banco de dados:", err)
	}

	log.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")

	log.Println("Iniciando AutoMigrate dos modelos GORM")

	err = DB.AutoMigrate(
		&models.Tenant{},
		&models.Categories{},
		&models.Products{},
		&models.PizzaSizes{},
		&models.PizzaFlavors{},
		&models.PizzaFlavorPrices{},
		&models.Order{},
		&models.OrderItem{},
		&models.PizzaFlavorPrices{},
	)

	if err != nil {
		log.Fatal("Erro fatal durante AutoMigrate", err)
	}

	log.Println("AutoMigrate concluído com sucesso! Tabelas criadas/atualizadas.")
}
