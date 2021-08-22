package security

import "golang.org/x/crypto/bcrypt"

func Hash(text string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	return string(bytes)
}

func ValidateHash(hashed, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(text))
	return err == nil
}
