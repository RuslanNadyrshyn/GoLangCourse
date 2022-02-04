package Connection

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	envName := "cfg/database.env"
	err := godotenv.Load(envName)
	if err != nil {
		log.Println("Error loading .env file")
	}
	dataSource := os.Getenv("DB_USER") + ":" + os.Getenv("PASS") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB")

	db, err := sql.Open(
		"mysql",
		dataSource,
	)
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected.")

	return db, err
}
