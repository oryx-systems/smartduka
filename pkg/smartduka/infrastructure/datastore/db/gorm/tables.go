package gorm

import (
	"time"

	"github.com/google/uuid"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"gorm.io/gorm"
)

// Base is the base table for all tables
type Base struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy *string   `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy *string   `gorm:"column:updated_by"`
}

// User models the system user
type User struct {
	Base

	ID        *string `gorm:"column:id"`
	FirstName string  `gorm:"column:first_name"`
	LastName  string  `gorm:"column:last_name"`
	Active    bool    `gorm:"column:active"`
	UserName  string  `gorm:"column:username"`
	UserType  string  `gorm:"column:user_type"`
	PushToken string  `gorm:"column:push_token"`
	Email     string  `gorm:"column:email"`
}

// BeforeCreate is a hook run before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	u.Base.CreatedAt = time.Now()
	u.Base.UpdatedAt = time.Now()
	u.ID = &id

	return
}

// TableName customizes how the table name is generated
func (User) TableName() string {
	return "user"
}

// Contact is a contact model for a user
type Contact struct {
	Base

	ID           string        `gorm:"column:id"`
	Active       bool          `gorm:"column:active"`
	ContactType  string        `gorm:"column:contact_type"`
	ContactValue string        `gorm:"column:contact_value"`
	Flavour      enums.Flavour `gorm:"column:flavour"`
	UserID       *string       `gorm:"column:user_id"`
}

// BeforeCreate is a hook run before creating user contact
func (c *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	c.CreatedAt = time.Now()
	c.ID = id

	return
}

// TableName customizes how the table name is generated
func (Contact) TableName() string {
	return "contact"
}

// UserPIN models the user's PIN table
type UserPIN struct {
	Base

	ID        string    `gorm:"column:id"`
	Active    bool      `gorm:"column:active"`
	ValidFrom time.Time `gorm:"column:valid_from"`
	ValidTo   time.Time `gorm:"column:valid_to"`
	HashedPIN string    `gorm:"column:hashed_pin"`
	Salt      string    `gorm:"column:salt"`
	UserID    string    `gorm:"column:user_id"`
}

// BeforeCreate is a hook run before creating user PIN
func (u *UserPIN) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New().String()
	u.CreatedAt = time.Now()
	u.ID = id

	return
}

// TableName customizes how the table name is generated
func (UserPIN) TableName() string {
	return "user_pin"
}

// OTP is model for one time password
type OTP struct {
	Base

	ID          string        `gorm:"column:id"`
	IsValid     bool          `gorm:"column:is_valid"`
	ValidUntil  time.Time     `gorm:"column:valid_until"`
	PhoneNumber string        `gorm:"column:phone_number"`
	OTP         string        `gorm:"column:otp"`
	Flavour     enums.Flavour `gorm:"column:flavour"`
	UserID      string        `gorm:"column:user_id"`
}

// BeforeCreate is a hook run before creating an OTP
func (o *OTP) BeforeCreate(tx *gorm.DB) (err error) {
	o.CreatedAt = time.Now()
	o.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (OTP) TableName() string {
	return "user_otp"
}

// Sale is used to show sales data
type Sale struct {
	Base

	ID        string  `gorm:"column:id"`
	Active    bool    `gorm:"column:active"`
	ProductID string  `gorm:"column:product_id"`
	Quantity  float64 `gorm:"column:quantity"`
	Unit      string  `gorm:"column:unit"`
	Price     float64 `gorm:"column:price"`
}

// BeforeCreate is a hook run before creating an OTP
func (s *Sale) BeforeCreate(tx *gorm.DB) (err error) {
	s.Base.CreatedAt = time.Now()
	s.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (Sale) TableName() string {
	return "sale"
}

// Product is used to display product info
type Product struct {
	Base

	ID           string  `gorm:"column:id"`
	Active       bool    `gorm:"column:active"`
	Name         string  `gorm:"column:name"`
	Category     string  `gorm:"column:category"`
	Quantity     float64 `gorm:"column:quantity"`
	Unit         string  `gorm:"column:unit"`
	Price        float64 `gorm:"column:price"`
	VAT          float64 `gorm:"column:vat"`
	Description  string  `gorm:"column:description"`
	Manufacturer string  `gorm:"column:manufacturer"`
	InStock      bool    `gorm:"column:in_stock"`
}

// BeforeCreate is a hook run before creating an OTP
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Base.CreatedAt = time.Now()
	p.ID = uuid.New().String()
	return
}

// TableName customizes how the table name is generated
func (Product) TableName() string {
	return "product"
}
