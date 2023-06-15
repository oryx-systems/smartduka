package dto

import "github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"

// LoginInput represents the login input
type LoginInput struct {
	PhoneNumber string        `json:"phone_number"`
	PIN         string        `json:"pin"`
	Flavour     enums.Flavour `json:"flavour"`
}

// RegisterUserInput represents the register user input
type RegisterUserInput struct {
	IdentificationDocumentNumber string        `json:"identification_number"`
	FirstName                    string        `json:"first_name"`
	MiddleName                   string        `json:"middle_name"`
	LastName                     string        `json:"last_name"`
	PhoneNumber                  string        `json:"phone_number"`
	Flavour                      enums.Flavour `json:"flavour"`
	UserName                     string        `json:"username"`
	DeviceToken                  string        `json:"device_token"`
	Residence                    string        `json:"residence"`
	PIN                          string        `json:"pin"`
	ConfirmPIN                   string        `json:"confirm_pin"`
}

// UserPINInput represents the user pin input
type UserPINInput struct {
	UserID     string        `json:"user_id"`
	PIN        string        `json:"pin"`
	ConfirmPIN string        `json:"confirm_pin"`
	Flavour    enums.Flavour `json:"flavour"`
}
