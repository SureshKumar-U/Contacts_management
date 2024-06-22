package main

import (
	"log"

	"github.com/SureshKumar-U/Contacts_management/config"

	"github.com/joho/godotenv"
)

func main() {
	app := config.NewAppConig()

	if err := godotenv.Load(); err != nil { // load environment variables
		log.Fatal("Error loading .env file")

	}

	if err := config.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}
	if err := app.Run(); err != nil {
		app.ErrLog.Printf(err.Error())
	}

}
