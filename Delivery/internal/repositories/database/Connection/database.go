package Connection

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		"nadyrshyn:683225@tcp(127.0.0.1:3306)/nadyrshyn_db",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected.")

	return db, err
}
