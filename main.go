package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/whalelogic/evbagel/handlers"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "EvBagel",
	})

	app.Use(logger.New())
	app.Use(recover.New())
	
	app.Static("/static", "./static")


	app.Get("/", handlers.HomeHandler)
	app.Get("/about", handlers.AboutHandler)
	app.Get("/contact", handlers.ContactHandler)

	log.Println("Starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
