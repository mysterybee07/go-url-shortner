package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/go-url-shortner/handlers"
)

func main() {
	app := fiber.New()

	shortner := handlers.NewUrlShortner()

	// routes setup
	app.Post("/shorten", shortner.ShortenUrl)
	app.Get("/:shortcode", shortner.RedirectUser)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}

}
