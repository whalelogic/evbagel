package handlers

import (
	"bytes"
	"context"
	"html/template"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/whalelogic/evbagel/components"
)

// renderTemplComponent renders a templ component to a string
func renderTemplComponent(component templ.Component) (string, error) {
	buf := new(bytes.Buffer)
	err := component.Render(context.Background(), buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// HomeHandler handles the home page
func HomeHandler(c *fiber.Ctx) error {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading template")
	}

	// Render templ components
	header, _ := renderTemplComponent(components.Header("EvBagel"))
	footer, _ := renderTemplComponent(components.Footer())
	buttonAbout, _ := renderTemplComponent(components.Button("Learn More", "/about", "primary"))
	buttonContact, _ := renderTemplComponent(components.Button("Contact Us", "/contact", "secondary"))

	data := map[string]interface{}{
		"Title":         "Home - EvBagel",
		"Heading":       "Welcome to EvBagel",
		"Header":        template.HTML(header),
		"Footer":        template.HTML(footer),
		"ButtonAbout":   template.HTML(buttonAbout),
		"ButtonContact": template.HTML(buttonContact),
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error rendering template")
	}

	c.Set("Content-Type", "text/html")
	return c.SendString(buf.String())
}

// AboutHandler handles the about page
func AboutHandler(c *fiber.Ctx) error {
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading template")
	}

	// Render templ components
	header, _ := renderTemplComponent(components.Header("EvBagel"))
	footer, _ := renderTemplComponent(components.Footer())
	buttonHome, _ := renderTemplComponent(components.Button("Back Home", "/", "primary"))
	buttonContact, _ := renderTemplComponent(components.Button("Contact Us", "/contact", "success"))

	data := map[string]interface{}{
		"Title":         "About - EvBagel",
		"Heading":       "About Us",
		"Header":        template.HTML(header),
		"Footer":        template.HTML(footer),
		"ButtonHome":    template.HTML(buttonHome),
		"ButtonContact": template.HTML(buttonContact),
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error rendering template")
	}

	c.Set("Content-Type", "text/html")
	return c.SendString(buf.String())
}

// ContactHandler handles the contact page
func ContactHandler(c *fiber.Ctx) error {
	tmpl, err := template.ParseFiles("templates/contact.html")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading template")
	}

	// Render templ components
	header, _ := renderTemplComponent(components.Header("EvBagel"))
	footer, _ := renderTemplComponent(components.Footer())
	buttonHome, _ := renderTemplComponent(components.Button("Back Home", "/", "primary"))
	buttonAbout, _ := renderTemplComponent(components.Button("About Us", "/about", "success"))

	data := map[string]interface{}{
		"Title":       "Contact - EvBagel",
		"Heading":     "Contact Us",
		"Header":      template.HTML(header),
		"Footer":      template.HTML(footer),
		"ButtonHome":  template.HTML(buttonHome),
		"ButtonAbout": template.HTML(buttonAbout),
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error rendering template")
	}

	c.Set("Content-Type", "text/html")
	return c.SendString(buf.String())
}
