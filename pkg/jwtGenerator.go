package pkg

import (
	"github.com/golang-jwt/jwt/v5"
)

// CreateUnsignedJwt creates an unsigned JWT with the provided claims.
func CreateUnsignedJwt(claims map[string]interface{}) (string, error) {
	jwtClaims := jwt.MapClaims{}
	for key, value := range claims {
		jwtClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodNone, jwtClaims)

	// Since we're using SigningMethodNone, we don't sign the token.
	jwtString, err := token.SigningString()
	if err != nil {
		return "", err
	}

	return jwtString + ".", nil
}
