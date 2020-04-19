package authentication

import (
	"os"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

// VerifyToken checks if a Google token is valid
func VerifyToken(token string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")

	v := googleAuthIDTokenVerifier.Verifier{}
	err := v.VerifyIDToken(token, []string{clientID})
	if err != nil {
		return nil, err
	}

	claimSet, err := googleAuthIDTokenVerifier.Decode(token)
	if err != nil {
		return nil, err
	}

	return claimSet, nil
}
