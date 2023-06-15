package dto

import (
	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
)

// LoginResponse represents the login response
type LoginResponse struct {
	UserProfile *domain.User `json:"user_profile"`
}
