package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// The secret key should be securely stored, ideally using a service like EnvKey, which provides
// a secure environment for managing and accessing sensitive information.
var jwtKey = []byte("my_secret_key_123")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// EnforceAuthentication enforces:
//
//	session token is passed in as the Authorization header
//	session token passed in is valid, and not expired
func EnforceAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "authorization token missing", http.StatusUnauthorized)
			return
		}

		if strings.Contains(tokenString, "Bearer") {
			sessionToken := strings.SplitAfter(tokenString, "Bearer ")

			if len(sessionToken) <= 1 {
				http.Error(w, "invalid token format", http.StatusUnauthorized)
				return
			}

			tokenString = sessionToken[1]
		}

		// You can add more validation logic here based on the claims if needed
		_, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
