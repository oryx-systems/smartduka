package gorm

import (
	"context"
	"fmt"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
)

// Update holds all the database record update methods
type Update interface {
	InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error)
	UpdateUser(ctx context.Context, user *User, updateData map[string]interface{}) (bool, error)
}

// InvalidatePIN invalidates a pin that is linked to the user profile when a new one is created
func (db *PGInstance) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) (bool, error) {
	err := db.DB.WithContext(ctx).Model(&UserPIN{}).Where(&UserPIN{UserID: userID, Active: true, Flavour: flavour}).Select("active").Updates(UserPIN{Active: false}).Error
	if err != nil {
		return false, fmt.Errorf("an error occurred while invalidating the pin: %v", err)
	}

	return true, nil
}

// UpdateUser updates a user record
func (db *PGInstance) UpdateUser(ctx context.Context, user *User, updateData map[string]interface{}) (bool, error) {
	err := db.DB.WithContext(ctx).Model(&user).Updates(updateData).Error
	if err != nil {
		return false, fmt.Errorf("an error occurred while updating the user: %v", err)
	}

	return true, nil
}
