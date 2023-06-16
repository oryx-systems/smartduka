package domain

import (
	"time"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
)

type User struct {
	ID              string          `json:"id"`
	FirstName       string          `json:"firstName"`
	LastName        string          `json:"lastName"`
	Active          bool            `json:"active"`
	UserName        string          `json:"username"`
	UserType        string          `json:"userType"`
	UserContact     Contact         `json:"userContact"`
	DeviceToken     string          `json:"deviceToken"`
	Email           string          `json:"email"`
	AuthCredentials AuthCredentials `json:"authCredentials"`
}

type AuthCredentials struct {
	RefreshToken string    `json:"refreshToken"`
	IDToken      string    `json:"idToken"`
	ExpiresIn    time.Time `json:"expiresIn"`
}

type Contact struct {
	ID           string        `json:"id"`
	Active       bool          `json:"active"`
	ContactType  string        `json:"contact_type"`
	ContactValue string        `json:"contact_value"`
	Flavour      enums.Flavour `json:"flavour"`
	UserID       string        `json:"user_id"`
}

type UserPIN struct {
	ID        string        `json:"id"`
	Active    bool          `json:"active"`
	Flavour   enums.Flavour `json:"flavour"`
	ValidFrom time.Time     `json:"valid_from"`
	ValidTo   time.Time     `json:"valid_to"`
	HashedPIN string        `json:"hashed_pin"`
	Salt      string        `json:"salt"`
	UserID    string        `json:"user_id"`
}

type LoginResponse struct {
	ID           string  `json:"id"`
	Username     string  `json:"username"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	UserContact  Contact `json:"contact"`
	AuthToken    string  `json:"auth_token"`
	RefreshToken string  `json:"refresh_token"`
}
