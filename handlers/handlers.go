package handlers

import (
	"fmt"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/go-url-shortner/helpers"
)

type UrlShortener struct {
	Urls map[string]string
}

func NewUrlShortner() *UrlShortener {
	return &UrlShortener{
		Urls: make(map[string]string),
	}
}

func (urlShortner *UrlShortener) Index(c *fiber.Ctx) error {
	// Ensure the request method is GET
	if c.Method() != fiber.MethodGet {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"error": "Invalid request method",
		})
	}

	// Parse the templates
	tmpl, err := template.ParseFiles("frontend/index.html", "frontend/form.html")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to parse or load template",
		})
	}

	// Set the content type as HTML
	c.Type("html")

	// Render the templates
	if err := tmpl.Execute(c.Response().BodyWriter(), nil); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to render template",
		})
	}

	return nil
}

func (urlShortner *UrlShortener) ShortenUrl(c *fiber.Ctx) error {
	// Ensure the request method is POST
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"error": "Invalid request method",
		})
	}

	// Get the original URL from the form data
	originalUrl := c.FormValue("url")
	if originalUrl == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL parameter is missing",
		})
	}

	// Generate a short key and store the original URL
	shortKey := helpers.GenerateShortKey()
	urlShortner.Urls[shortKey] = originalUrl

	// Create the shortened URL
	shortenedUrl := fmt.Sprintf("http://localhost:8080/%s", shortKey)

	// Parse the HTML template files
	tmpl, err := template.ParseFiles("frontend/shorturl.html", "frontend/form.html")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to parse or load template",
		})
	}

	data := map[string]string{
		"OriginalUrl":  originalUrl,
		"ShortenedUrl": shortenedUrl,
	}

	c.Type("html")

	// dynamic data to the template
	if err := tmpl.Execute(c.Response().BodyWriter(), data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to render template",
		})
	}

	return nil
}

func (urlShortner *UrlShortener) RedirectUser(c *fiber.Ctx) error {
	//shortcode from url path
	shortCode := c.Params("shortcode")
	if shortCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Shortened Key is missing",
		})
	}

	// original url of the shortcode retrieval
	originalUrl, exists := urlShortner.Urls[shortCode]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shortened Key not found",
		})
	}
	return c.Redirect(originalUrl, fiber.StatusMovedPermanently)
}
