package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
)

// ConnectDatabase ...
func ConnectDatabase() *gorm.DB {
	dbURI := os.Getenv("DATABASE_URL")
	if dbURI == "" {
		username := ViperEnvVariable("DB_USER")
		password := ViperEnvVariable("DB_PASS")
		dbName := ViperEnvVariable("DB_NAME")
		dbHost := ViperEnvVariable("DB_HOST")
		dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	}
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	return db
}