package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fir1/rest-api/http/rest/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading .env File")
	}

	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	userHandler := &handlers.UserHandler{DB: db}

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("SERVER_PORT")
	fmt.Printf("Starting Server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
