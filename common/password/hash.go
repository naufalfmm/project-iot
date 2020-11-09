package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	bytePass := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
