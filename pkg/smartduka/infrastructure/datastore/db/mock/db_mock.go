package mock

import (
	"context"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
)

// DataStoreMock is a mock implementation of the datastore interface
type DataStoreMock struct {
	MockRegisterUserFn                func(ctx context.Context, user *domain.User, contact *domain.Contact) error
	MockSaveOTPFn                     func(ctx context.Context, otp *domain.OTP) error
	MockSavePINFn                     func(ctx context.Context, pinInput *domain.UserPIN) (bool, error)
	MockGetUserProfileByUserIDFn      func(ctx context.Context, userID string) (*domain.User, error)
	MockGetUserProfileByPhoneNumberFn func(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error)
	MockGetUserPINByUserIDFn          func(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error)
	MockInvalidatePINFn               func(ctx context.Context, userID string, flavour enums.Flavour) (bool, error)
	MockSearchUserFn                  func(ctx context.Context, searchTerm string) ([]*domain.User, error)
	MockUpdateUserFn                  func(ctx context.Context, user *domain.User, updateData map[string]interface{}) (bool, error)
}

// NewDataStoreMock returns a new instance of the mock datastore
func NewDataStoreMock() *DataStoreMock {
	user := &domain.User{
		ID:        uuid.New().String(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Active:    true,
		UserName:  gofakeit.Username(),
		UserType:  "ADMIN",
		UserContact: domain.Contact{
			ID: uuid.New().String(),
		},
		DeviceToken: uuid.New().String(),
	}

	return &DataStoreMock{
		MockRegisterUserFn: func(ctx context.Context, user *domain.User, contact *domain.Contact) error {
			return nil
		},
		MockSaveOTPFn: func(ctx context.Context, otp *domain.OTP) error {
			return nil
		},
		MockSavePINFn: func(ctx context.Context, pinInput *domain.UserPIN) (bool, error) {
			return true, nil
		},
		MockGetUserProfileByUserIDFn: func(ctx context.Context, userID string) (*domain.User, error) {
			return nil, nil
		},
		MockGetUserProfileByPhoneNumberFn: func(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error) {
			return nil, nil
		},
		MockGetUserPINByUserIDFn: func(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error) {
			return nil, nil
		},
		MockInvalidatePINFn: func(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
			return true, nil
		},
		MockSearchUserFn: func(ctx context.Context, searchTerm string) ([]*domain.User, error) {
			return []*domain.User{
				user,
			}, nil
		},
		MockUpdateUserFn: func(ctx context.Context, user *domain.User, updateData map[string]interface{}) (bool, error) {
			return true, nil
		},
	}
}

// RegisterUser mocks the RegisterUser method
func (m *DataStoreMock) RegisterUser(ctx context.Context, user *domain.User, contact *domain.Contact) error {
	return m.MockRegisterUserFn(ctx, user, contact)
}

// SaveOTP mocks the SaveOTP method
func (m *DataStoreMock) SaveOTP(ctx context.Context, otp *domain.OTP) error {
	return m.MockSaveOTPFn(ctx, otp)
}

// SavePIN mocks the SavePIN method
func (m *DataStoreMock) SavePIN(ctx context.Context, pinInput *domain.UserPIN) (bool, error) {
	return m.MockSavePINFn(ctx, pinInput)
}

// GetUserProfileByUserID mocks the GetUserProfileByUserID method
func (m *DataStoreMock) GetUserProfileByUserID(ctx context.Context, userID string) (*domain.User, error) {
	return m.MockGetUserProfileByUserIDFn(ctx, userID)
}

// GetUserProfileByPhoneNumber mocks the GetUserProfileByPhoneNumber method
func (m *DataStoreMock) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error) {
	return m.MockGetUserProfileByPhoneNumberFn(ctx, phoneNumber, flavour)
}

// GetUserPINByUserID mocks the GetUserPINByUserID method
func (m *DataStoreMock) GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error) {
	return m.MockGetUserPINByUserIDFn(ctx, userID, flavour)
}

// InvalidatePIN mocks the InvalidatePIN method
func (m *DataStoreMock) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
	return m.MockInvalidatePINFn(ctx, userID, flavour)
}

// SearchUser mocks the SearchUser method
func (m *DataStoreMock) SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error) {
	return m.MockSearchUserFn(ctx, searchTerm)
}

// UpdateUser mocks the UpdateUser method
func (m *DataStoreMock) UpdateUser(ctx context.Context, user *domain.User, updateData map[string]interface{}) (bool, error) {
	return m.MockUpdateUserFn(ctx, user, updateData)
}
