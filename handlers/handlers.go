package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/go-url-shortner/helpers"
)

type UrlShortener struct {
	Urls map[string]string
}

// initialize above struct with an in-memory map
func NewUrlShortner() *UrlShortener {
	return &UrlShortener{
		Urls: make(map[string]string),
	}
}

// payload for shortening URLs.
type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

func (urlShortner *UrlShortener) ShortenUrl(c *fiber.Ctx) error {
	var req ShortenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// empty url validation
	if req.LongURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL parameter is missing",
		})
	}
	// Format validation of URL
	if !helpers.ValidateURL(req.LongURL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL format",
		})
	}

	shortKey := helpers.GenerateShortKey()
	urlShortner.Urls[shortKey] = req.LongURL

	//generated short key
	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"short_url": shortURL,
	})
}

func (urlShortner *UrlShortener) RedirectUser(c *fiber.Ctx) error {
	//take shortCode from request parameter
	shortCode := c.Params("shortcode")
	if shortCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Shortened Key is missing",
		})
	}

	//extract original url from request shortCode
	originalUrl, exists := urlShortner.Urls[shortCode]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Shortened Key not found",
		})
	}
	//redirect users to original url
	return c.Redirect(originalUrl, fiber.StatusMovedPermanently)
}
