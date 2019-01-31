package firebase

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func FetchToken(r *http.Request) (*auth.Token, error) {
	credentials := os.Getenv("FIREBASE_AUTH_CREDENTIALS")
	if credentials == "" {
		return &auth.Token{}, fmt.Errorf("%v", "not found credentials.json")
	}

	app, err := firebase.NewApp(r.Context(),
		nil, option.WithCredentialsFile(credentials))
	if err != nil {
		return &auth.Token{}, err
	}

	client, err := app.Auth(r.Context())
	if err != nil {
		return &auth.Token{}, err
	}

	jwt := strings.Replace(
		r.Header.Get("Authorization"), "Bearer ", "", 1)

	token, err := client.VerifyIDToken(r.Context(), jwt)
	return token, err
}
