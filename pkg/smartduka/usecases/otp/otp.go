package otp

import (
	"context"
	"fmt"
	"time"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/helpers"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/extension"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore"
	"github.com/sirupsen/logrus"
)

const (
	appName = "Smartduka"
)

// UseCasesOTP contain all the method required for OTP delivery
type UseCasesOTP interface {
	GenerateAndSendOTP(ctx context.Context, phoneNumber string, flavour enums.Flavour) (string, error)
}

// UseCasesOTPImpl represents the user otp usecase implementation
type UseCasesOTPImpl struct {
	Create datastore.Create
	Query  datastore.Query
	Ext    extension.Extension
}

// NewUseCaseOTP initializes the new otp implementation
func NewUseCaseOTP(
	create datastore.Create,
	query datastore.Query,
) UseCasesOTP {
	ext := extension.NewExtension()
	return &UseCasesOTPImpl{
		Create: create,
		Query:  query,
		Ext:    ext,
	}
}

// GenerateAndSendOTP generates and sends an OTP to the user
func (o *UseCasesOTPImpl) GenerateAndSendOTP(ctx context.Context, phoneNumber string, flavour enums.Flavour) (string, error) {
	validatePhoneNumber, err := helpers.NormalizeMSISDN(phoneNumber)
	if err != nil {
		return "", err
	}

	userProfile, err := o.Query.GetUserProfileByPhoneNumber(ctx, *validatePhoneNumber, flavour)
	if err != nil {
		return "", err
	}

	if !flavour.IsValid() {
		return "", fmt.Errorf("invalid flavour")
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		return "", fmt.Errorf("failed to generate an OTP")
	}

	var message string
	switch flavour {
	case enums.FlavourConsumer:
		message = fmt.Sprintf("Your %v verification code is %s", appName, otp)
	case enums.FlavourPro:
		message = fmt.Sprintf("Your %v verification code is %s", appName, otp)
	}

	// TODO: 1. Implement send sms logic here
	logrus.Print("OTP MESSAGE: ", message)

	otpData := &domain.OTP{
		IsValid:     true,
		ValidUntil:  time.Now().Add(time.Minute * 5),
		PhoneNumber: *validatePhoneNumber,
		OTP:         otp,
		Flavour:     flavour,
		Medium:      "SMS",
		UserID:      userProfile.ID,
	}

	// Save the OTP to the database
	_, err = o.Create.SaveOTP(ctx, otpData)
	if err != nil {
		return "", err
	}

	return otp, nil
}
