package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aaron-g-sanchez/PROTOTYPE/PROJECT-ATHENA-PROTO/backend/database/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

/* USER table
- ID
- Name
- User_id
- created_at
- Updated_at
*/

func loadEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return os.Getenv("DB_CONNECTION_STRING")
}

func main() {
	connStr := loadEnv()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	data := []models.User{}

	rows, err := db.Query(`SELECT name, user_id FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string
	var userId string

	for rows.Next() {
		err := rows.Scan(&name, &userId)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, models.User{name, userId})
	}

	fmt.Println(data)
}
