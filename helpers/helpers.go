package helpers

import (
	"sync"
)

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	counter int64
	mu      sync.Mutex
)

// Base62Encode converts a number to a Base62 encoded string.
func Base62Encode(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	encoded := ""
	for num > 0 {
		remainder := num % 62
		encoded = string(base62Chars[remainder]) + encoded
		num = num / 62
	}
	return encoded
}

// GenerateShortKey generates a unique short key using Base62 encoding.
func GenerateShortKey() string {
	mu.Lock()
	defer mu.Unlock()

	// Increment the counter to get a new unique identifier.
	counter++
	return Base62Encode(counter)
}
