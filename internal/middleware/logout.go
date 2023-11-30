package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func Logout(w http.ResponseWriter, r *http.Request) {

	claims := MyCustomClaims{
		"",
		jwt.StandardClaims{
			ExpiresAt: -15000,
			Issuer:    "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString([]byte("keymaker"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating JWT token: " + err.Error()))
	}
	w.Header().Set("Authorization", "Bearer "+tokenstring)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token: " + tokenstring))
}
