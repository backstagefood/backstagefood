package main

import (
	"log"
	"os"

	db "github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/backstagefood/backstagefood/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/backstagefood/backstagefood/docs"
)

// @title Backstage Food API
// @version 1.0
// @description API for managing products and orders for Backstage Food.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
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
