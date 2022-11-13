package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mfsiat/go_api/database"
)

func main() {
	database.ConnectDb()
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	setupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, Siat 👋!")
}