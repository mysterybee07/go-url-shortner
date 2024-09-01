package helpers

import (
	"crypto/rand"
	"math/big"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeToBase62(data []byte) string {
	var result string
	num := new(big.Int).SetBytes(data)
	base := big.NewInt(62)
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		result = string(base62Chars[mod.Int64()]) + result
	}

	return result
}

// generates a random short key for URLs
func GenerateShortKey() string {

	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return EncodeToBase62(b)
}
