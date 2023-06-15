package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
)

func TestSendOTP(t *testing.T) {
	ctx := context.Background()
	graphQLURL := fmt.Sprintf("%s/%s", baseURL, "v1/auth/graphql")

	headers, err := GetGraphQLHeaders(ctx)
	if err != nil {
		t.Errorf("failed to get GraphQL headers: %v", err)
		return
	}

	graphqlMutation := `
	mutation sendOTP($phoneNumber: String!, $flavour: Flavour!) {
		sendOTP(phoneNumber: $phoneNumber, flavour: $flavour)
	  }
	`

	type args struct {
		query map[string]interface{}
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "success: send OTP",
			args: args{
				query: map[string]interface{}{
					"query": graphqlMutation,
					"variables": map[string]interface{}{
						"phoneNumber": testPhone,
						"flavour":     enums.FlavourPro,
					},
				},
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "success: unable to send OTP",
			args: args{
				query: map[string]interface{}{
					"query": graphqlMutation,
					"variables": map[string]interface{}{
						"phoneNumber": "testPhone",
						"flavour":     enums.FlavourPro,
					},
				},
			},
			wantStatus: http.StatusOK,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := mapToJSONReader(tt.args.query)
			if err != nil {
				t.Errorf("unable to get GQL JSON io Reader: %s", err)
				return
			}

			r, err := http.NewRequest(
				http.MethodPost,
				graphQLURL,
				body,
			)
			if err != nil {
				t.Errorf("unable to compose request: %s", err)
				return
			}

			if r == nil {
				t.Errorf("nil request")
				return
			}

			for k, v := range headers {
				r.Header.Add(k, v)
			}
			client := http.Client{
				Timeout: time.Second * testHTTPClientTimeout,
			}
			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("request error: %s", err)
				return
			}

			dataResponse, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("can't read request body: %s", err)
				return
			}
			if dataResponse == nil {
				t.Errorf("nil response data")
				return
			}

			data := map[string]interface{}{}
			err = json.Unmarshal(dataResponse, &data)
			if err != nil {
				t.Errorf("bad data returned")
				return
			}

			if tt.wantErr {
				errMsg, ok := data["errors"]
				if !ok {
					t.Errorf("GraphQL error: %s", errMsg)
					return
				}
			}

			if !tt.wantErr {
				_, ok := data["errors"]
				if ok {
					t.Errorf("error not expected, got %v", data["errors"])
					return
				}
			}
			if tt.wantStatus != resp.StatusCode {
				t.Errorf("Bad status response returned, expected %v, got %v", tt.wantStatus, resp.StatusCode)
				return
			}
		})
	}
}
