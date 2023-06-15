package extension

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/helpers"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
)

// Extension holds the methods that are used by the extension
type Extension interface {
	MakeRequest(ctx context.Context, method string, path string, body interface{}) (*http.Response, error)
	GetLoggedInUserUID(ctx context.Context) (string, error)
}

// ExtImpl implements the Extension interface
type ExtImpl struct {
}

// NewExtension initializes the extension methods
func NewExtension() Extension {
	return &ExtImpl{}
}

// MakeRequest performs a http request and returns a response
func (e ExtImpl) MakeRequest(ctx context.Context, method string, path string, body interface{}) (*http.Response, error) {
	apiKey := helpers.MustGetEnvVar("AIT_API_KEY")

	client := &http.Client{}

	// A GET request should not send data when doing a request. We should use query parameters
	// instead of having a request body. In some cases where a GET request has an empty body {},
	// it might result in status code 400 with the error:
	//  `Your client has issued a malformed or illegal request. That’s all we know.`
	if method == http.MethodGet {
		req, reqErr := http.NewRequestWithContext(ctx, method, path, nil)
		if reqErr != nil {
			return nil, reqErr
		}

		req.Header.Set("apiKey", apiKey)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		return client.Do(req)
	}

	encoded, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(encoded)
	req, reqErr := http.NewRequestWithContext(ctx, method, path, payload)
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Set("apiKey", apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

// GetLoggedInUserUID returns the UID of the logged in user
func (e ExtImpl) GetLoggedInUserUID(ctx context.Context) (string, error) {
	uid, err := utils.GetLoggedInUser(ctx)
	if err != nil {
		return "", err
	}

	return uid, nil
}
