package security

import "golang.org/x/crypto/bcrypt"

func Hash(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	return string(bytes), err
}

func ValidateHash(hashed, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(text))
	return err == nil
}
