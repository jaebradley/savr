package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	fmt.Printf("Claim set is %v", claimSet)

	return
}
