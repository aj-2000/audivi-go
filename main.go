package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aj-2000/audivi-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", routes.Hw)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "0.0.0.0:6969"
	} else {
		port = "0.0.0.0:" + port
	}

	return port
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()

	//app.Use(csrf.New())
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(getPort()))
}