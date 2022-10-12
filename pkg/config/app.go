package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load("netlify.tolm")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_User")
	dbPass := os.Getenv("DB_Password")
	dbName := os.Getenv("DB_Name")
	dbHost := os.Getenv("DB_Host")
	dbPort := os.Getenv("DB_Port")

	var DSN = " host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort

	log.Println(DSN)

	d, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		log.Println("DB Connected")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
