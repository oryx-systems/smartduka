package gorm_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore/db/gorm"
)

func TestPGInstance_RegisterUser(t *testing.T) {
	invalidID := "invalid"

	type args struct {
		ctx     context.Context
		user    *gorm.User
		contact *gorm.Contact
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case: Successfully register user",
			args: args{
				ctx: context.Background(),
				user: &gorm.User{
					ID:        &userID,
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
					Active:    true,
					UserName:  gofakeit.BeerHop(),
					UserType:  "ADMIN",
					PushToken: gofakeit.UUID(),
				},
				contact: &gorm.Contact{
					ID:           uuid.New().String(),
					Active:       true,
					ContactType:  "PHONE",
					ContactValue: testPhone,
					Flavour:      enums.FlavourConsumer,
					UserID:       &userID,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to register user",
			args: args{
				ctx: context.Background(),
				user: &gorm.User{
					ID:        &invalidID,
					FirstName: gofakeit.FirstName(),
					LastName:  gofakeit.LastName(),
					Active:    true,
					UserName:  gofakeit.BeerHop(),
					UserType:  "TENANT",
					PushToken: gofakeit.UUID(),
				},
				contact: &gorm.Contact{
					ID:           "uuid.New().String()",
					Active:       true,
					ContactType:  "PHONE",
					ContactValue: testPhone,
					Flavour:      enums.FlavourConsumer,
					UserID:       &invalidID,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.RegisterUser(tt.args.ctx, tt.args.user, tt.args.contact)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPGInstance_SaveOTP(t *testing.T) {
	type args struct {
		ctx context.Context
		otp *gorm.OTP
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.OTP
		wantErr bool
	}{
		{
			name: "Happy case: save otp",
			args: args{
				ctx: context.Background(),
				otp: &gorm.OTP{
					IsValid:     true,
					ValidUntil:  time.Now().Add(time.Duration(6)),
					PhoneNumber: gofakeit.PhoneFormatted(),
					OTP:         "1200",
					Flavour:     "PRO",
					UserID:      userID,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to save otp",
			args: args{
				ctx: context.Background(),
				otp: &gorm.OTP{
					IsValid:     true,
					ValidUntil:  time.Now().Add(time.Duration(6)),
					PhoneNumber: gofakeit.PhoneFormatted(),
					OTP:         "1200",
					Flavour:     "PRO",
					UserID:      "userID",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.SaveOTP(tt.args.ctx, tt.args.otp)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.SaveOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_SavePIN(t *testing.T) {
	type args struct {
		ctx     context.Context
		pinData *gorm.UserPIN
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.UserPIN
		wantErr bool
	}{
		{
			name: "Happy case: save otp",
			args: args{
				ctx: nil,
				pinData: &gorm.UserPIN{
					ID:        uuid.NewString(),
					Active:    true,
					ValidFrom: time.Now(),
					ValidTo:   time.Now().Add(time.Hour * 3),
					HashedPIN: "hashed",
					Salt:      "salt",
					UserID:    userID,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to save otp",
			args: args{
				ctx: nil,
				pinData: &gorm.UserPIN{
					ID:        uuid.NewString(),
					Active:    true,
					ValidFrom: time.Now(),
					ValidTo:   time.Now().Add(time.Hour * 3),
					HashedPIN: "hashed",
					Salt:      "salt",
					UserID:    "userID",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.SavePIN(tt.args.ctx, tt.args.pinData)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.SavePIN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_AddProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		product *gorm.Product
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.Product
		wantErr bool
	}{
		{
			name: "Happy case: add product",
			args: args{
				ctx: context.Background(),
				product: &gorm.Product{
					Base: gorm.Base{
						CreatedBy: &userID,
					},
					ID:           uuid.NewString(),
					Active:       true,
					Name:         gofakeit.BeerName(),
					Category:     "DETERGENT",
					Quantity:     12.00,
					Unit:         "DOZEN",
					Price:        900.89,
					VAT:          16.00,
					Description:  gofakeit.HipsterSentence(30),
					Manufacturer: gofakeit.BeerMalt(),
					InStock:      true,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to add product",
			args: args{
				ctx: context.Background(),
				product: &gorm.Product{
					ID:           "test",
					Active:       true,
					Name:         gofakeit.BeerName(),
					Category:     "DETERGENT",
					Quantity:     12,
					Unit:         "DOZEN",
					Price:        900.89,
					VAT:          16.00,
					Description:  gofakeit.HipsterSentence(30),
					Manufacturer: gofakeit.BeerMalt(),
					InStock:      true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.AddProduct(tt.args.ctx, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPGInstance_AddSaleRecord(t *testing.T) {
	type args struct {
		ctx  context.Context
		sale *gorm.Sale
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.Sale
		wantErr bool
	}{
		{
			name: "Happy case: record sale",
			args: args{
				ctx: context.Background(),
				sale: &gorm.Sale{
					Base: gorm.Base{
						CreatedBy: &userID,
					},
					ID:        uuid.NewString(),
					ProductID: productID,
					Quantity:  2.00,
					Unit:      "DOZEN",
					Price:     15.40,
				},
			},
			wantErr: false,
		},
		{
			name: "Sad case: unable to record sale",
			args: args{
				ctx: context.Background(),
				sale: &gorm.Sale{
					Base: gorm.Base{
						CreatedBy: &userID,
					},
					ID:        uuid.NewString(),
					ProductID: uuid.NewString(),
					Quantity:  97.0,
					Unit:      "DOZEN",
					Price:     15.40,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := testingDB.AddSaleRecord(tt.args.ctx, tt.args.sale)
			if (err != nil) != tt.wantErr {
				t.Errorf("PGInstance.AddSaleRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
