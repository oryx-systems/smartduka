package db

import (
	"context"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore/db/gorm"
)

// InvalidatePIN invalidates a pin that is linked to the user profile.
// This is done by toggling the IsValid field to false
func (d *DbServiceImpl) InvalidatePIN(ctx context.Context, userID string, flavour enums.Flavour) error {
	return d.update.InvalidatePIN(ctx, userID, flavour)
}

// UpdateUser updates a user record
func (d *DbServiceImpl) UpdateUser(ctx context.Context, user *domain.User, updateData map[string]interface{}) error {
	data := &gorm.User{
		ID: &user.ID,
	}

	return d.update.UpdateUser(ctx, data, updateData)
}

// UpdateProduct updates product details in the database
func (d *DbServiceImpl) UpdateProduct(ctx context.Context, product *domain.Product, updateData map[string]interface{}) error {
	data := &gorm.Product{
		ID: product.ID,
	}

	return d.update.UpdateProduct(ctx, data, updateData)
}
