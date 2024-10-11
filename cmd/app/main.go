package main

import (
	"log"
	"os"

	db "github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/backstagefood/backstagefood/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	databaseRepository := db.New()

	echoEngine := echo.New()
	echoEngine.Use(middleware.Logger())
	echoEngine.Use(middleware.Recover())

	routes := routes.New(echoEngine)
	routes.Start(os.Getenv("SERVER_PORT"), databaseRepository)
}
