package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/helpers"
)

// Create the JWT key used to create the signature
var (
	jwtKey = []byte(helpers.MustGetEnvVar("JWT_SECRET"))
)

// TokenResponse represents the response from the token endpoint
type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresIn time.Time `json:"expiresIn"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWTToken generates a JWT token
func GenerateJWTToken(userID string) (*TokenResponse, error) {
	if userID == "" {
		return nil, fmt.Errorf("user id is required")
	}
	// Create the Claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3)),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		Token:     tokenString,
		ExpiresIn: claims.ExpiresAt.Time,
	}, nil
}

// ValidateJWTToken validates a JWT token
func ValidateJWTToken(tokenString string) (*TokenResponse, error) {

	tkn, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		// Now, create a new token for the current use, with a renewed expiration time
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
		claims.IssuedAt = jwt.NewNumericDate(time.Now())

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Create the JWT string
		tokenString, err = token.SignedString(jwtKey)
		if err != nil {
			return nil, err
		}

		return &TokenResponse{
			Token:     tokenString,
			ExpiresIn: claims.ExpiresAt.Time,
		}, nil
	}

	return &TokenResponse{
		Token:     tokenString,
		ExpiresIn: claims.ExpiresAt.Time,
	}, nil

}

// GetLoggedInUser retrieves the logged in user from the context
func GetLoggedInUser(ctx context.Context) (string, error) {
	UID := ctx.Value(common.AuthTokenContextKey).(string)

	tkn, err := jwt.ParseWithClaims(UID, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := tkn.Claims.(*Claims)
	if !ok {
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}
