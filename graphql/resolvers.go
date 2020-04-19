package graphql

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/dgrijalva/jwt-go"
	graphqlgo "github.com/graphql-go/graphql"
	"github.com/jaebradley/savr/database"
)

func getCurrentUser(params graphqlgo.ResolveParams) (interface{}, error) {
	userContext := params.Context.Value("user")
	fmt.Println("user context is", userContext)
	if userContext == nil {
		return database.User{}, errors.New("Unable to identify current user")
	}
	claims := userContext.(*jwt.Token).Claims.(jwt.MapClaims)
	emailAddress, err := mail.ParseAddress(claims["sub"].(string))

	if err != nil {
		return database.User{}, errors.New("Unable to parse email address for current user")
	}

	user := database.GetUserByEmailAddress(emailAddress)
	return user, nil
}
