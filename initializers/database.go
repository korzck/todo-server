package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB() {
	db_url := os.Getenv("DB_URL")
	fmt.Println(db_url)
	var err error
	DB, err = sql.Open("postgres", db_url)
	if err != nil {
		fmt.Println(err.Error())
		// println("error from db")
	}
	// defer DB.Close()
}
