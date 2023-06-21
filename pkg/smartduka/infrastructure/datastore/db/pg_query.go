package db

import (
	"context"
	"fmt"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/enums"
	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
)

// GetUserProfileByUserID fetches and returns a userprofile using their user ID
func (d *DbServiceImpl) GetUserProfileByUserID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := d.query.GetUserProfileByUserID(ctx, &userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile by user ID: %v", err)
	}

	// contact := &domain.Contact{
	// 	ID:           user.UserContact.ID,
	// 	Active:       user.UserContact.Active,
	// 	ContactType:  user.UserContact.ContactType,
	// 	ContactValue: user.UserContact.ContactValue,
	// 	Flavour:      user.UserContact.Flavour,
	// 	UserID:       *user.ID,
	// }

	return &domain.User{
		ID:          *user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Active:      user.Active,
		UserName:    user.UserName,
		UserType:    user.UserType,
		DeviceToken: user.PushToken,
	}, nil
}

// GetUserProfileByPhoneNumber fetches and returns a userprofile using their phone number
func (d *DbServiceImpl) GetUserProfileByPhoneNumber(ctx context.Context, phoneNumber string, flavour enums.Flavour) (*domain.User, error) {
	user, err := d.query.GetUserProfileByPhoneNumber(ctx, phoneNumber, flavour)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile by phonenumber: %v", err)
	}

	// contact := &domain.Contact{
	// 	ID:           user.UserContact.ID,
	// 	Active:       user.UserContact.Active,
	// 	ContactType:  user.UserContact.ContactType,
	// 	ContactValue: user.UserContact.ContactValue,
	// 	Flavour:      user.UserContact.Flavour,
	// 	UserID:       *user.ID,
	// }

	return &domain.User{
		ID:          *user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Active:      user.Active,
		UserName:    user.UserName,
		UserType:    user.UserType,
		DeviceToken: user.PushToken,
	}, nil
}

// GetUserPINByUserID fetches and returns a user PIN using their user ID
func (d *DbServiceImpl) GetUserPINByUserID(ctx context.Context, userID string, flavour enums.Flavour) (*domain.UserPIN, error) {
	if userID == "" {
		return nil, fmt.Errorf("user id cannot be empty")
	}
	pinData, err := d.query.GetUserPINByUserID(ctx, userID, flavour)
	if err != nil {
		return nil, fmt.Errorf("failed query and retrieve user PIN data: %s", err)
	}

	return &domain.UserPIN{
		UserID:    pinData.UserID,
		HashedPIN: pinData.HashedPIN,
		ValidFrom: pinData.ValidFrom,
		ValidTo:   pinData.ValidTo,
		Active:    pinData.Active,
		Salt:      pinData.Salt,
	}, nil
}

// SearchUser searches for users in the system using a search term
func (d *DbServiceImpl) SearchUser(ctx context.Context, searchTerm string) ([]*domain.User, error) {
	var users []*domain.User

	records, err := d.query.SearchUser(ctx, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search user: %v", err)
	}

	for _, record := range records {
		// contact := &domain.Contact{
		// 	ID:           record.UserContact.ID,
		// 	Active:       record.UserContact.Active,
		// 	ContactType:  record.UserContact.ContactType,
		// 	ContactValue: record.UserContact.ContactValue,
		// 	Flavour:      record.UserContact.Flavour,
		// 	UserID:       *record.ID,
		// }

		users = append(users, &domain.User{
			ID:          *record.ID,
			FirstName:   record.FirstName,
			LastName:    record.LastName,
			Active:      record.Active,
			UserName:    record.UserName,
			UserType:    record.UserType,
			DeviceToken: record.PushToken,
		})
	}

	return users, nil
}

// GetProductByID retrieves a product using its ID
func (d *DbServiceImpl) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	return nil, nil
}

// GetDailySale retrieves daily sales
func (d *DbServiceImpl) GetDailySale(ctx context.Context) ([]*domain.Sale, error) {
	return nil, nil
}

// SearchProduct searches a product using the term provided by the user
func (d *DbServiceImpl) SearchProduct(ctx context.Context, searchTerm string) (*domain.Product, error) {
	return nil, nil
}
