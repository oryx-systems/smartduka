package gorm_test

import (
	"context"
	"testing"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore/db/gorm"
)

func TestPGInstance_GetUserProfileByUserID(t *testing.T) {
	invalidUserID := "invalid"
	type args struct {
		ctx    context.Context
		userID *string
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.User
		wantErr bool
	}{
		{
			name: "Happy case: get user profile by user id",
			args: args{
				ctx:    context.Background(),
				userID: &userID,
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get user profile by user id",
			args: args{
				ctx:    context.Background(),
				userID: &invalidUserID,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.GetUserProfileByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.GetUserProfileByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_GetUserProfileByPhoneNumber(t *testing.T) {
	type args struct {
		ctx         context.Context
		phoneNumber string
		flavour     enums.Flavour
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.User
		wantErr bool
	}{
		{
			name: "Happy case: get user profile by phone",
			args: args{
				ctx:         context.Background(),
				phoneNumber: testPhone,
				flavour:     "PRO",
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get user profile by phone",
			args: args{
				ctx:     context.Background(),
				flavour: "PRO",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.GetUserProfileByPhoneNumber(tt.args.ctx, tt.args.phoneNumber, tt.args.flavour)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.GetUserProfileByPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_GetUserPINByUserID(t *testing.T) {
	invalidUserID := "invalid"
	type args struct {
		ctx     context.Context
		userID  string
		flavour enums.Flavour
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.UserPIN
		wantErr bool
	}{
		{
			name: "Happy case: get user PIN by user id",
			args: args{
				ctx:     context.Background(),
				userID:  userID,
				flavour: "PRO",
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get user PIN by user id",
			args: args{
				ctx:     context.Background(),
				userID:  invalidUserID,
				flavour: "PRO",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.GetUserPINByUserID(tt.args.ctx, tt.args.userID, tt.args.flavour)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.GetUserPINByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_SearchUser(t *testing.T) {
	type args struct {
		ctx        context.Context
		searchTerm string
	}
	tests := []struct {
		name    string
		args    args
		want    []*gorm.User
		wantErr bool
	}{
		{
			name: "Happy case: search user",
			args: args{
				ctx:        context.Background(),
				searchTerm: "test",
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to search user",
			args: args{
				ctx:        context.Background(),
				searchTerm: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.SearchUser(tt.args.ctx, tt.args.searchTerm)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_GetProductByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.Product
		wantErr bool
	}{
		{
			name: "Happy case: get product by id",
			args: args{
				ctx: context.Background(),
				id:  productID,
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to get product by id",
			args: args{
				ctx: context.Background(),
				id:  "productID",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.GetProductByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_GetDailySale(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*gorm.Sale
		wantErr bool
	}{
		{
			name: "Happy case: Get daily sales",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.GetDailySale(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.GetDailySale() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_SearchProduct(t *testing.T) {
	type args struct {
		ctx        context.Context
		searchTerm string
	}
	tests := []struct {
		name    string
		args    args
		want    []*gorm.Product
		wantErr bool
	}{
		{
			name: "Happy case: search product",
			args: args{
				ctx:        context.Background(),
				searchTerm: "Panadol",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.SearchProduct(tt.args.ctx, tt.args.searchTerm)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.SearchProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
