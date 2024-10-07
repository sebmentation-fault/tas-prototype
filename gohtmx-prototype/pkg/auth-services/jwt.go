package authservices

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/db"
)

var JWTSecret = []byte("A super-secret secret")

// generate a signed token from a JWT claim
// errors if can't
func NewSignedJWTTokenWithClaims(account *db.Account) (string, error) {
	claim := jwt.MapClaims{
		"id":   account.ID,                            // the user's id
		"type": account.Type,                          // stores user's type
		"exp":  time.Now().Add(time.Hour * 72).Unix(), // expires in 72 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(JWTSecret)

	return signedToken, err
}

// validate a JWT token
func ValidateJWTToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return the secret key for validation
		return JWTSecret, nil
	})
}

// attempts to get an account object from the token
//
// if unsuccessful, returns an error
func GetAccountIdAndTypeFromToken(t *jwt.Token) (int64, int, error) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return -1, -1, errors.New("could not produce claim")
	}

	// Extract and convert the ID
	idFloat, ok := claims["id"].(float64)
	if !ok {
		return -1, -1, errors.New("invalid ID type in claims")
	}
	id := int64(idFloat)

	// Extract and convert the user type
	typeFloat, ok := claims["type"].(float64)
	if !ok {
		return -1, -1, errors.New("invalid user type in claims")
	}
	userType := int(typeFloat)

	return id, userType, nil
}
