package domain

import (
	"time"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
)

type OTP struct {
	ID          string        `gorm:"column:id"`
	IsValid     bool          `gorm:"column:is_valid"`
	ValidUntil  time.Time     `gorm:"column:valid_until"`
	PhoneNumber string        `gorm:"column:phone_number"`
	OTP         string        `gorm:"column:otp"`
	Flavour     enums.Flavour `gorm:"column:flavour"`
	Medium      string        `gorm:"column:medium"`
	UserID      string        `gorm:"column:user_id"`
}
