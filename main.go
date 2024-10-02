package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}

	// get server port from .env file
	serverPort := os.Getenv("SERVER_PORT")
	if "" == serverPort {
		serverPort = "8080"
	}
	log.Println("startins server on port ", serverPort)

	// define routes
	http.HandleFunc("/", greet)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}
