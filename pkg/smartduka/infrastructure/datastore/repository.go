package datastore

import (
	"context"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
)

// Create is a collection of methods to carry out create operations on the database
type Create interface {
	RegisterUser(ctx context.Context, user *domain.User, contact *domain.Contact) (*domain.User, error)
	SaveOTP(ctx context.Context, otp *domain.OTP) (*domain.OTP, error)
	SavePIN(ctx context.Context, pinInput *domain.UserPIN) (*domain.UserPIN, error)

	AddProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	AddSaleRecord(ctx context.Context, sale *domain.Sale) (*domain.Sale, error)
}

// Query hold a collection of methods to interact with the querying of any data
type Query interface {
	GetUserProfileByUserID(ctx context.Context, userID string) (*domain.User, error)
	GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error)
	GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error)
	SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error)

	GetProductByID(ctx context.Context, id string) (*domain.Product, error)
	GetDailySale(ctx context.Context) ([]*domain.Sale, error)
	SearchProduct(ctx context.Context, searchTerm string) (*domain.Product, error)
}

// Update is a collection of methods with the ability to update any data
type Update interface {
	InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) error
	UpdateUser(ctx context.Context, user *domain.User, updateData map[string]interface{}) error

	UpdateProduct(ctx context.Context, product *domain.Product, updateData map[string]interface{}) error
}
