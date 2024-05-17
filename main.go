package main

import (
	"github.com/gofiber/fiber/v2"
	"go-project-management/database"
	"go-project-management/routes"
	"log"
)

func main() {
	// Code here
	database.InitDatabase()

	defer database.CloseDatabase()

	app := fiber.New()

	routes.RouteInit(app)

	err := app.Listen("localhost:3000")
	if err != nil {
		log.Println("Failed to listen Go Fiber Server")
	}
}
