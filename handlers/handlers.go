package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/whalelogic/evbagel/templates"
)

func render(c *fiber.Ctx, component templ.Component) error {
	return adaptor.HTTPHandler(templ.Handler(component))(c)
}

func HomeHandler(c *fiber.Ctx) error {
	return render(c, templates.HomePage())
}

func AboutHandler(c *fiber.Ctx) error {
	return render(c, templates.AboutPage())
}

func ContactHandler(c *fiber.Ctx) error {
	return render(c, templates.ContactPage())
}