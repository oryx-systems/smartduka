package rest

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
)

// authCheckFn is a function type for authorization and authentication checks
// there can be several e.g an authentication check runs first then an authorization
// check runs next if the authentication passes etc
type authCheckFn = func(
	c *gin.Context,
) (bool, map[string]string, string)

// AuthMiddleware is a gin middleware that checks if the request is authorized
// and authenticated
func AuthMiddleware() gin.HandlerFunc {
	// multiple checks will be run in sequence (order matters)
	// the first check to succeed will call `c.Next()` and `return`
	// this means that more permissive checks (e.g exceptions) should come first
	checkFuncs := []authCheckFn{HasValidFirebaseBearerToken}

	return func(c *gin.Context) {
		errs := []map[string]string{}
		for _, checkFn := range checkFuncs {
			// run the check function
			isAuthorized, errMap, token := checkFn(c)

			// if the check function returned true, call `c.Next()` and `return`
			if isAuthorized {
				ctx := context.WithValue(c.Request.Context(), common.AuthTokenContextKey, token)
				c.Request = c.Request.WithContext(ctx)

				c.Next()
				return
			}

			errs = append(errs, errMap)
		}

		// if all the checks failed, return a 401
		c.JSON(401, gin.H{"errors": errs})
		c.Abort()
	}
}

// HasValidFirebaseBearerToken returns true with no errors if the request has a valid bearer token in the authorization header.
// Otherwise, it returns false and the error in a map with the key "error"
func HasValidFirebaseBearerToken(c *gin.Context) (bool, map[string]string, string) {
	bearerToken, err := ExtractBearerToken(c)
	if err != nil {
		// this error here will only be returned to the user if all the verification functions in the chain fail
		return false, utils.ErrorMap(err), ""
	}

	validatedToken, err := utils.ValidateJWTToken(bearerToken)
	if err != nil {
		return false, utils.ErrorMap(err), ""
	}

	return true, nil, validatedToken.Token
}

// ExtractBearerToken gets a bearer token from an Authorization header.
//
// This is expected to contain a Firebase idToken prefixed with "Bearer "
func ExtractBearerToken(r *gin.Context) (string, error) {
	return ExtractToken(r, "Authorization", "Bearer")
}

// ExtractToken extracts a token with the specified prefix from the specified header
func ExtractToken(c *gin.Context, header string, prefix string) (string, error) {
	if c.Request == nil {
		return "", fmt.Errorf("request is nil")
	}

	if c.Request.Header == nil {
		return "", fmt.Errorf("no headers, can't extract bearer token")
	}

	authHeader := c.Request.Header.Get(header)
	if authHeader == "" {
		return "", fmt.Errorf("expected an `%s` request header", header)
	}

	if !strings.HasPrefix(authHeader, prefix) {
		return "", fmt.Errorf("the `Authorization` header contents should start with `Bearer`")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	return token, nil
}
