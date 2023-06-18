package domain

import (
	"time"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
)

type OTP struct {
	ID          string        `json:"id"`
	IsValid     bool          `json:"is_valid"`
	ValidUntil  time.Time     `json:"valid_until"`
	PhoneNumber string        `json:"phone_number"`
	OTP         string        `json:"otp"`
	Flavour     enums.Flavour `json:"flavour"`
	Medium      string        `json:"medium"`
	UserID      string        `json:"user_id"`
}
