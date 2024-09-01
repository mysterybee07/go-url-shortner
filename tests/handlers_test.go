package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/go-url-shortner/handlers"
	"github.com/stretchr/testify/assert"
)

func TestUrlShortener_ShortenUrl(t *testing.T) {
	app := fiber.New()
	urlShortener := handlers.NewUrlShortner()
	app.Post("/shorten", urlShortener.ShortenUrl)

	tests := []struct {
		name         string
		body         string
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			name:         "Valid URL",
			body:         `{"long_url": "http://example.com"}`,
			expectedCode: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				// Check if the short URL starts with the base URL and is not empty
				"short_url": "http://localhost:8080/",
			},
		},
		{
			name:         "Invalid URL format",
			body:         `{"long_url": "htp://invalid-url"}`,
			expectedCode: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid URL format",
			},
		},
		{
			name:         "Missing URL parameter",
			body:         `{"long_url": ""}`,
			expectedCode: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "URL parameter is missing",
			},
		},
		{
			name:         "Invalid JSON format",
			body:         `{"long_url": "http://example.com"`,
			expectedCode: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid request payload",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := bytes.NewReader([]byte(tt.body))
			httpReq, err := http.NewRequest("POST", "/shorten", req)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}
			httpReq.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(httpReq)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			var response map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			if tt.name == "Valid URL" {
				shortURL, ok := response["short_url"].(string)
				if !ok {
					t.Fatalf("Expected short_url to be a string")
				}
				assert.True(t, len(shortURL) > len("http://localhost:8080/"), "Short URL should have a short key")
				assert.True(t, shortURL[:len("http://localhost:8080/")] == "http://localhost:8080/", "Short URL should start with base URL")
			} else {
				for key, value := range tt.expectedBody {
					assert.Equal(t, value, response[key])
				}
			}
		})
	}
}

func TestUrlShortener_RedirectUser(t *testing.T) {
	app := fiber.New()
	urlShortener := handlers.NewUrlShortner()
	// Initialize with some test data
	urlShortener.Urls = map[string]string{
		"abc123": "http://example.com",
		"xyz789": "http://example.org",
	}

	app.Get("/:shortcode", urlShortener.RedirectUser)

	tests := []struct {
		name         string
		shortcode    string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid Shortcode",
			shortcode:    "abc123",
			expectedCode: fiber.StatusMovedPermanently,
			expectedBody: "http://example.com",
		},
		{
			name:         "Invalid Shortcode",
			shortcode:    "invalid",
			expectedCode: fiber.StatusNotFound,
			expectedBody: `{"error": "Shortened Key not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := bytes.NewReader([]byte{})
			httpReq, err := http.NewRequest("GET", "/"+tt.shortcode, req)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			resp, err := app.Test(httpReq)
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}

			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if tt.expectedCode == fiber.StatusMovedPermanently {
				// Check the Location header
				assert.Equal(t, tt.expectedBody, resp.Header.Get("Location"))
			} else {
				// Read response body
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("Failed to read response body: %v", err)
				}
				defer resp.Body.Close()

				// Compare response body
				assert.JSONEq(t, tt.expectedBody, string(body))
			}
		})
	}
}
