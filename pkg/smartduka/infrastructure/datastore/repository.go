package datastore

import (
	"context"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
)

// Create is a collection of methods to carry out create operations on the database
type Create interface {
	RegisterUser(ctx context.Context, user *domain.User, contact *domain.Contact) error
	SaveOTP(ctx context.Context, otp *domain.OTP) error
	SavePIN(ctx context.Context, pinInput *domain.UserPIN) (bool, error)
}

// Query hold a collection of methods to interact with the querying of any data
type Query interface {
	GetUserProfileByUserID(ctx context.Context, userID string) (*domain.User, error)
	GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error)
	GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error)
	SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error)
}

// Update is a collection of methods with the ability to update any data
type Update interface {
	InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error)
	UpdateUser(ctx context.Context, user *domain.User, updateData map[string]interface{}) (bool, error)
}
