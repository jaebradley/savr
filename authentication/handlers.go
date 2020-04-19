package authentication

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/jaebradley/savr/database"
)

type authenticationData struct {
	Token string `json:"token"`
}

// GoogleAuthenticationHandler handles Google Authentication Callback
func GoogleAuthenticationHandler(response http.ResponseWriter, request *http.Request) {
	var data authenticationData

	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&data)
	if err != nil {
		http.Error(response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	claimSet, err := VerifyToken(data.Token)

	if err != nil {
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	emailAddress, err := mail.ParseAddress(claimSet.Email)
	if err != nil {
		http.Error(response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var user database.User
	if !database.UserWithEmailAddressExists(emailAddress) {
		user = database.CreateUser(emailAddress)
	} else {
		user = database.GetUserByEmailAddress(emailAddress)
	}

	token, err := CreateToken(&user)

	if err != nil {
		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(
		response,
		&http.Cookie{
			Name:     "access_token",
			Value:    token,
			HttpOnly: true,
		},
	)

	response.WriteHeader(http.StatusNoContent)

	return
}
