package gorm

import (
	"context"
)

// Create holds all the database record creation methods
type Create interface {
	RegisterUser(ctx context.Context, user *User, contact *Contact) (*User, error)
	SaveOTP(ctx context.Context, otp *OTP) (*OTP, error)
	SavePIN(ctx context.Context, pinData *UserPIN) (*UserPIN, error)

	AddProduct(ctx context.Context, product *Product) (*Product, error)
	AddSaleRecord(ctx context.Context, sale *Sale) (*Sale, error)
}

// RegisterUser creates a new user record.
// The user can be a resident or a staff member
func (db *PGInstance) RegisterUser(ctx context.Context, user *User, contact *Contact) (*User, error) {
	tx := db.DB.WithContext(ctx).Begin()

	// create user
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// create contact
	contact.UserID = user.ID
	if err := tx.Create(&contact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return user, nil
}

// SaveOTP saves an OTP in the database
func (db *PGInstance) SaveOTP(ctx context.Context, otp *OTP) (*OTP, error) {
	if err := db.DB.WithContext(ctx).Create(&otp).Error; err != nil {
		return nil, err
	}

	return otp, nil
}

// SavePIN saves a pin in the database
func (db *PGInstance) SavePIN(ctx context.Context, pinData *UserPIN) (*UserPIN, error) {
	if err := db.DB.WithContext(ctx).Create(&pinData).Error; err != nil {
		return nil, err
	}

	return pinData, nil
}

// Adds a product into the database
func (db *PGInstance) AddProduct(ctx context.Context, product *Product) (*Product, error) {
	if err := db.DB.WithContext(ctx).Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// AddSaleRecord adds sale record in the database
func (db *PGInstance) AddSaleRecord(ctx context.Context, sale *Sale) (*Sale, error) {
	if err := db.DB.WithContext(ctx).Create(&sale).Error; err != nil {
		return nil, err
	}

	return sale, nil
}
