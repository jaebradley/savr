package authentication

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/jaebradley/savr/database"
)

// CreateToken creates a signed JWT for a given user
func CreateToken(user *database.User) (signedToken string, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&jwt.StandardClaims{
			Subject:   user.EmailAddress,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(0, 0, 14).Unix(),
		},
	)
	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
}

// ParseToken parses a signed token into a parsed token
func ParseToken(signedToken string) (parsedToken *jwt.Token, err error) {
	token, err := jwt.Parse(
		signedToken,
		func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("JWT_SIGNING_KEY"), nil
		})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return token, nil
	}

	return nil, nil
}
