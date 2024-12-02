package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() string {
	godotenv.Load()

	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")

	fmt.Println(db_name, db_user, db_host, db_pass)

	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require", db_name, db_user, db_pass, db_host)

	return connStr
}
