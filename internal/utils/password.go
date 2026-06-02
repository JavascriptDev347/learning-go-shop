package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword using generate from password and return the hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword for compare password and hashPassword, return true if password is correct, otherwise return false
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
