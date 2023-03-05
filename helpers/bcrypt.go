package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	hashResult, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	return string(hashResult)
}

func ComparePassword(inputPass, dbPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(inputPass), []byte(dbPass))
	return err == nil
}
