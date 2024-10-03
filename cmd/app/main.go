package main

import (
	"github.com/backstagefood/backstagefood/internal/app/adapter/driven/postgresql"
	"github.com/backstagefood/backstagefood/internal/app/adapter/driver/http"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Info("error loading .env file")
		//log.Fatal("error loading .env file")
	}

	postgresql.New()
	server := http.New()

	server.Start(os.Getenv("SERVER_PORT"))
}
