package gorm

import (
	"context"
	"fmt"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"gorm.io/gorm/clause"
)

// Query holds all the database record query methods
type Query interface {
	GetUserProfileByUserID(ctx context.Context, userID *string) (*User, error)
	GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*User, error)
	GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*UserPIN, error)
	SearchUser(ctx context.Context, searchTerm string) ([]*User, error)

	GetProductByID(ctx context.Context, id string) (*Product, error)
	GetDailySale(ctx context.Context) ([]*Sale, error)
	SearchProduct(ctx context.Context, searchTerm string) (*Product, error)
}

// GetUserProfileByUserID fetches a user profile using the user ID
func (db *PGInstance) GetUserProfileByUserID(ctx context.Context, userID *string) (*User, error) {
	var user User
	if err := db.DB.Where(&User{ID: userID, Active: true}).Preload(clause.Associations).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by user ID %v: %v", userID, err)
	}
	return &user, nil
}

// GetUserProfileByPhoneNumber fetches a user profile using the phone number
func (db *PGInstance) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*User, error) {
	var user User

	if err := db.DB.Joins("JOIN smartduka_contact on smartduka_user.id = smartduka_contact.user_id").Where("smartduka_contact.contact_value = ? AND smartduka_contact.flavour = ?", phoneNumber, flavour).Preload(clause.Associations).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by phonenumber %v: %v", phoneNumber, err)
	}

	return &user, nil
}

// GetUserPINByUserID fetches a user's pin using the user ID and Flavour
func (db *PGInstance) GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*UserPIN, error) {
	if !flavour.IsValid() {
		return nil, fmt.Errorf("flavour is not valid")
	}
	var pin UserPIN
	if err := db.DB.Where(&UserPIN{UserID: userID, Active: true, Flavour: flavour}).First(&pin).Error; err != nil {
		return nil, fmt.Errorf("failed to get pin: %v", err)
	}

	return &pin, nil
}

// SearchUser searches for a user using the search term
func (db *PGInstance) SearchUser(ctx context.Context, searchTerm string) ([]*User, error) {
	var users []*User
	if err := db.DB.Joins("JOIN smartduka_contact on smartduka_user.id = smartduka_contact.user_id").
		Joins("JOIN smartduka_identifier on smartduka_user.id = smartduka_identifier.user_id").
		Where("smartduka_contact.contact_value ILIKE ? OR smartduka_user.first_name ILIKE ? "+
			"OR smartduka_user.last_name ILIKE ? OR smartduka_user.username ILIKE ? OR smartduka_identifier.identifier_value ILIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%").
		Where("smartduka_user.active = ?", true).
		Preload(clause.Associations).Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to search user: %v", err)
	}

	return users, nil
}

// GetProductByID retrieves a product using its ID
func (db *PGInstance) GetProductByID(ctx context.Context, id string) (*Product, error) {
	return nil, nil
}

// GetDailySale retrieves daily sales
func (db *PGInstance) GetDailySale(ctx context.Context) ([]*Sale, error) {
	return nil, nil
}

// SearchProduct searches a product using the term provided by the user
func (db *PGInstance) SearchProduct(ctx context.Context, searchTerm string) (*Product, error) {
	return nil, nil
}
