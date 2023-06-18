package gorm

import (
	"context"
	"fmt"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
)

// Update holds all the database record update methods
type Update interface {
	InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) error
	UpdateUser(ctx context.Context, user *User, updateData map[string]interface{}) error

	UpdateProduct(ctx context.Context, product *Product, updateData map[string]interface{}) error
}

// InvalidatePIN invalidates a pin that is linked to the user profile when a new one is created
func (db *PGInstance) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) error {
	err := db.DB.WithContext(ctx).Model(&UserPIN{}).Where(&UserPIN{UserID: userID, Active: true, Flavour: flavour}).Select("active").Updates(UserPIN{Active: false}).Error
	if err != nil {
		return fmt.Errorf("an error occurred while invalidating the pin: %v", err)
	}

	return nil
}

// UpdateUser updates a user record
func (db *PGInstance) UpdateUser(ctx context.Context, user *User, updateData map[string]interface{}) error {
	err := db.DB.WithContext(ctx).Model(&user).Updates(updateData).Error
	if err != nil {
		return fmt.Errorf("an error occurred while updating the user: %v", err)
	}

	return nil
}

// UpdateProduct updates product details
func (db *PGInstance) UpdateProduct(ctx context.Context, product *Product, updateData map[string]interface{}) error {
	return nil
}
