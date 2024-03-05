package password

import "golang.org/x/crypto/bcrypt"

// Generate - Generate bcrypt hash password
func Generate(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// Verify - Verify bcrypt hash password
func Verify(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
