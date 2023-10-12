package passwordgenerator

import (
	"math/rand"
	"time"
)

type PasswordGenerator struct {
}

func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{}
}

func (p *PasswordGenerator) GenerateSafePassword(size int, withCapitalizedChar, withNumbers, withSpecialChar bool) string {
	// Set default values if not provided
	if size == 0 {
		size = 8
	}
	if !withCapitalizedChar && !withNumbers && !withSpecialChar {
		withCapitalizedChar = false
		withNumbers = false
		withSpecialChar = false
	}

	// Define character sets based on the requirements
	charSet := "abcdefghijklmnopqrstuvwxyz"
	if withCapitalizedChar {
		charSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if withNumbers {
		charSet += "0123456789"
	}
	if withSpecialChar {
		charSet += "!@#$%^&*()-_"
	}

	// Generate the password
	password := make([]byte, size)
	randSrc := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSrc)
	for i := 0; i < size; i++ {
		password[i] = charSet[r.Intn(len(charSet))]
	}

	return string(password)
}
