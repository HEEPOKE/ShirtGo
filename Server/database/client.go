package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	username = "root"
// 	password = "root"
// 	hostname = "localhost:3306"
// 	dbName   = "shirtgo"
// )

var (
	env      = make(map[string]string) // Simulated environoment
	password = env["DATABASE_PWD"]
)

func init() {
	// Simulate loading .env file (godotenv.Load())
	env["DATABASE_PWD"] = "foo"
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	fmt.Println("DB Connected!!")
	defer db.Close()
}
