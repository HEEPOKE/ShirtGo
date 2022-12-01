package amin

import (
	"fmt"
	"shirt/Server/database"
	"shirt/Server/routes"

	"github.com/jho/godotenv"
)

func main() {
	err := godotenvLoad(".env")
	if err != nil {
		mt.Println("Error loading .env file")
	}

	database.Connect()
	routes.Router()

}
