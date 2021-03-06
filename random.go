package utils

import "crypto/rand"

const (
	AlphabetDigits       = "0123456789"
	AlphabetLettersUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphabetLettersLower = "abcdefghijklmnopqrstuvwxyz"
)

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

//Generate generates tan digits string (10 digits length)
func GenerateRandomString(length int, alphabet string) (string, error) {
	bytes, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = alphabet[b%byte(len(alphabet))]
	}
	return string(bytes), nil
}
