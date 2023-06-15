package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/imroc/req"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/testutils"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
	"github.com/oryx-systems/smartduka/pkg/smartduka/presentation"
)

const (
	testHTTPClientTimeout = 180
)

var (
	srv       *http.Server
	baseURL   string
	serverErr error
)

func mapToJSONReader(m map[string]interface{}) (io.Reader, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal map to JSON: %w", err)
	}

	buf := bytes.NewBuffer(bs)
	return buf, nil
}

func TestMain(m *testing.M) {
	log.Printf("Setting tests up ...")

	initialEnv := os.Getenv("ENVIRONMENT")
	os.Setenv("ENVIRONMENT", "staging")

	setupFixtures()

	ctx := context.Background()

	srv, baseURL, serverErr = testutils.StartTestServer(
		ctx,
		presentation.PrepareServer,
		presentation.SmartdukaServiceAllowedOrigins,
	)
	if serverErr != nil {
		log.Printf("unable to start test server: %s", serverErr)
	}

	// run tests
	log.Printf("Running tests ...")
	code := m.Run()

	// restore envs
	os.Setenv("ENVIRONMENT", initialEnv)

	log.Printf("finished running tests")

	// cleanup here
	defer func() {
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("test server shutdown error: %s", err)
		}
	}()

	os.Exit(code)
}

// GetGraphQLHeaders gets relevant GraphQLHeaders
func GetGraphQLHeaders(ctx context.Context) (map[string]string, error) {
	authorization, err := GetBearerTokenHeader(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't Generate Bearer Token: %s", err)
	}
	return req.Header{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": authorization,
	}, nil
}

// GetBearerTokenHeader gets bearer Token Header
func GetBearerTokenHeader(ctx context.Context) (string, error) {
	customToken, err := utils.GenerateJWTToken(userID)
	if err != nil {
		return "", fmt.Errorf("can't create custom token: %s", err)
	}

	if customToken.Token == "" {
		return "", fmt.Errorf("blank custom token: %s", err)
	}

	idTokens, err := utils.ValidateJWTToken(customToken.Token)
	if err != nil {
		return "", fmt.Errorf("can't validate custom token: %s", err)
	}
	if idTokens == nil {
		return "", fmt.Errorf("nil idTokens")
	}

	return fmt.Sprintf("Bearer %s", idTokens.Token), nil
}
