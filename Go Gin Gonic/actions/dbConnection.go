package actions

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DbConecction() (db *sql.DB) {

	user := getEnv("USER")
	pwd := getEnv("PWD")
	port := getEnv("PORT")
	schema := getEnv("SCHEMA")

	path := (user + ":" + pwd + "@tcp(" + port + ")" + "/" + schema + "")

	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func getEnv(key string) string {

	err1 := godotenv.Load("config.yml")
	if err1 != nil {
		log.Fatal(err1.Error())
	}
	return os.Getenv(key)
}
