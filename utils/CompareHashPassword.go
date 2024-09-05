package utils

import "golang.org/x/crypto/bcrypt"

func ComapareHasPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
