package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID int) (string, error) {
	var SECRET_KEY = []byte(GetAsString("STAGE", "kuncirahasia"))

	//set payload
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	//set algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	var SECRET_KEY = []byte(GetAsString("STAGE", "kuncirahasia"))

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil

	})

	if err != nil {
		return token, err
	}

	return token, nil
}
