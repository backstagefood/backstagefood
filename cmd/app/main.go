package main

import (
	"log"
	"os"

	database "github.com/backstagefood/backstagefood/internal/app/adapter/driven/postgresql"
	server "github.com/backstagefood/backstagefood/internal/app/adapter/driver/http"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	db := database.New()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := server.New(e, db)
	server.Start(os.Getenv("SERVER_PORT"))
}
