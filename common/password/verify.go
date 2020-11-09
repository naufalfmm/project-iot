package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Verify(hashedPassword string, plainPassword string) (bool, error) {
	byteHashedPass := []byte(hashedPassword)
	bytePlainPass := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(byteHashedPass, bytePlainPass)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
