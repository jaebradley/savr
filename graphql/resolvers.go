package graphql

import (
	"errors"
	"fmt"
	"strconv"

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
	userID, err := strconv.ParseUint(claims["sub"].(string), 10, 64)
	if err != nil {
		return database.User{}, errors.New("Unable to parse user id from claims subject")
	}
	user := database.GetUserByID(userID)
	return user, nil
}
