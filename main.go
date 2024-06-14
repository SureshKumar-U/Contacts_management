package main

import (
	"log"
	"net/http"

	"github.com/SureshKumar-U/Contacts_management/Routes"
	"github.com/SureshKumar-U/Contacts_management/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")

	}

	router := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := config.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	Routes.ContactRoutes(router)
	Routes.UserRoutes(router)
	log.Fatal(server.ListenAndServe())

}
