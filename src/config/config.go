package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	AppPort           = ""
	DbSringConnection = ""
)

func LoadEnvs() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var (
		appPort    = os.Getenv("APP_PORT")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbName     = os.Getenv("DB_NAME")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
	)

	if dbPort == "" {
		dbPort = "3306"
	}

	AppPort = appPort
	DbSringConnection = makeStringConnection(dbHost, dbPort, dbName, dbUser, dbPassword)
}

func makeStringConnection(dbHost string, dbPort string, dbName string, dbUser string, dbPassword string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
}
