package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/whalelogic/evbagel/handlers"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "EvBagel",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve static files
	app.Static("/static", "./static")

	// Routes
	app.Get("/", handlers.HomeHandler)
	app.Get("/about", handlers.AboutHandler)
	app.Get("/contact", handlers.ContactHandler)

	// Start server
	log.Println("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
