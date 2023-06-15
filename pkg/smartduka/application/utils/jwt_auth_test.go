package utils_test

import (
	"context"
	"testing"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
)

func TestGenerateJWTToken(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: generate JWT token",
			args: args{
				userID: "123",
			},
			wantErr: false,
		},
		{
			name:    "Sad case: unable to generate JWT token",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := utils.GenerateJWTToken(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestValidateJWTToken(t *testing.T) {
	token, err := utils.GenerateJWTToken("123")
	if err != nil {
		t.Error("Unable to generate JWT token")
	}

	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: validate JWT token",
			args: args{
				tokenString: token.Token,
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to validate JWT token",
			args: args{
				tokenString: "123",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := utils.ValidateJWTToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetLoggedInUser(t *testing.T) {
	token, err := utils.GenerateJWTToken("123")
	if err != nil {
		t.Error("Unable to generate JWT token")
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: get logged in user",
			args: args{
				ctx: context.WithValue(context.Background(), common.AuthTokenContextKey, token.Token),
			},
			wantErr: false,
		},
		{
			name: "Sad case: get logged in user",
			args: args{
				ctx: context.WithValue(context.Background(), common.AuthTokenContextKey, "UID"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := utils.GetLoggedInUser(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoggedInUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
