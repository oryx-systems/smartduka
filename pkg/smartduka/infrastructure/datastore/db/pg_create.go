package db

import (
	"context"
	"fmt"

	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore/db/gorm"
)

// RegisterUser registers a new user in the database
func (d *DbServiceImpl) RegisterUser(ctx context.Context, user *domain.User, contact *domain.Contact) error {
	usr := &gorm.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Active:      true,
		UserName:    user.UserName,
		UserType:    user.UserType,
		DeviceToken: user.DeviceToken,
	}

	ct := &gorm.Contact{
		Active:       true,
		ContactType:  contact.ContactType,
		ContactValue: contact.ContactValue,
		Flavour:      contact.Flavour,
		UserID:       usr.ID,
	}

	return d.create.RegisterUser(ctx, usr, ct)
}

// SaveOTP saves an OTP in the database
func (d *DbServiceImpl) SaveOTP(ctx context.Context, otp *domain.OTP) error {
	otpData := &gorm.OTP{
		IsValid:     otp.IsValid,
		ValidUntil:  otp.ValidUntil,
		PhoneNumber: otp.PhoneNumber,
		OTP:         otp.OTP,
		Flavour:     otp.Flavour,
		Medium:      otp.Medium,
		UserID:      otp.UserID,
	}

	return d.create.SaveOTP(ctx, otpData)
}

// SavePIN saves a PIN in the database
func (d *DbServiceImpl) SavePIN(ctx context.Context, pinInput *domain.UserPIN) (bool, error) {
	pinObj := &gorm.UserPIN{
		UserID:    pinInput.UserID,
		HashedPIN: pinInput.HashedPIN,
		ValidFrom: pinInput.ValidFrom,
		ValidTo:   pinInput.ValidTo,
		Active:    pinInput.Active,
		Flavour:   pinInput.Flavour,
		Salt:      pinInput.Salt,
	}

	_, err := d.create.SavePIN(ctx, pinObj)
	if err != nil {
		return false, fmt.Errorf("failed to save user pin: %v", err)
	}

	return true, nil
}
