package database

import (
	"assignment-02/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitializeDB() {
	godotenv.Load()

	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")

	var psqlInfo string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Sad! Error connecting to database:", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
	fmt.Println("Congrats! Success, connected to database")
}

func GetDB() *gorm.DB {
	return db
}
