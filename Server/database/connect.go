package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() {
	db, err := sql.Open("mysql")
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	fmt.Println("DB Connected!!")

}
