package db

import (
	"context"
	"fmt"

	"github.com/oryx-systems/smartduka/pkg/smartduka/domain"
	"github.com/oryx-systems/smartduka/pkg/smartduka/infrastructure/datastore/db/gorm"
)

// RegisterUser registers a new user in the database
func (d *DbServiceImpl) RegisterUser(ctx context.Context, user *domain.User, contact *domain.Contact) (*domain.User, error) {
	usr := &gorm.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Active:      user.Active,
		UserName:    user.UserName,
		UserType:    user.UserType,
		UserContact: gorm.Contact{},
		DeviceToken: user.DeviceToken,
		Email:       user.Email,
	}

	contactData := &gorm.Contact{
		Active:       contact.Active,
		ContactType:  contact.ContactType,
		ContactValue: contact.ContactValue,
		Flavour:      contact.Flavour,
		UserID:       usr.ID,
	}

	response, err := d.create.RegisterUser(ctx, usr, contactData)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        *response.ID,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Active:    response.Active,
		UserName:  response.UserName,
		UserType:  response.UserType,
		UserContact: domain.Contact{
			ID:           response.UserContact.ID,
			Active:       response.UserContact.Active,
			ContactType:  response.UserContact.ContactType,
			ContactValue: response.UserContact.ContactValue,
			Flavour:      response.UserContact.Flavour,
			UserID:       *response.UserContact.UserID,
		},
		DeviceToken: response.DeviceToken,
		Email:       response.Email,
	}, nil
}

// SaveOTP saves an OTP in the database
func (d *DbServiceImpl) SaveOTP(ctx context.Context, otp *domain.OTP) (*domain.OTP, error) {
	otpData := &gorm.OTP{
		IsValid:     otp.IsValid,
		ValidUntil:  otp.ValidUntil,
		PhoneNumber: otp.PhoneNumber,
		OTP:         otp.OTP,
		Flavour:     otp.Flavour,
		Medium:      otp.Medium,
		UserID:      otp.UserID,
	}

	result, err := d.create.SaveOTP(ctx, otpData)
	if err != nil {
		return nil, err
	}

	return &domain.OTP{
		ID:          result.ID,
		IsValid:     result.IsValid,
		ValidUntil:  result.ValidUntil,
		PhoneNumber: result.PhoneNumber,
		OTP:         result.OTP,
		Flavour:     result.Flavour,
		Medium:      result.Medium,
		UserID:      result.UserID,
	}, nil
}

// SavePIN saves a PIN in the database
func (d *DbServiceImpl) SavePIN(ctx context.Context, pinInput *domain.UserPIN) (*domain.UserPIN, error) {
	pinObj := &gorm.UserPIN{
		UserID:    pinInput.UserID,
		HashedPIN: pinInput.HashedPIN,
		ValidFrom: pinInput.ValidFrom,
		ValidTo:   pinInput.ValidTo,
		Active:    pinInput.Active,
		Flavour:   pinInput.Flavour,
		Salt:      pinInput.Salt,
	}

	result, err := d.create.SavePIN(ctx, pinObj)
	if err != nil {
		return nil, fmt.Errorf("failed to save user pin: %v", err)
	}

	return &domain.UserPIN{
		ID:        result.ID,
		Active:    result.Active,
		Flavour:   result.Flavour,
		ValidFrom: result.ValidFrom,
		ValidTo:   result.ValidTo,
		HashedPIN: result.HashedPIN,
		Salt:      result.Salt,
		UserID:    result.UserID,
	}, nil
}

// Adds a product into the database
func (d *DbServiceImpl) AddProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	productObj := &gorm.Product{
		Active:       product.Active,
		Name:         product.Name,
		Category:     product.Category,
		Quantity:     product.Quantity,
		Unit:         product.Unit,
		Price:        product.Price,
		Description:  product.Description,
		Manufacturer: product.Manufacturer,
		InStock:      product.InStock,
	}

	result, err := d.create.AddProduct(ctx, productObj)
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		ID:           result.ID,
		Active:       result.Active,
		Name:         result.Name,
		Category:     result.Category,
		Quantity:     result.Quantity,
		Unit:         result.Unit,
		Price:        result.Price,
		Description:  result.Description,
		Manufacturer: result.Manufacturer,
		InStock:      result.InStock,
	}, nil
}

// AddSaleRecord adds sale record in the database
func (d *DbServiceImpl) AddSaleRecord(ctx context.Context, sale *domain.Sale) (*domain.Sale, error) {
	saleObj := &gorm.Sale{
		ProductID: sale.ProductID,
		Quantity:  sale.Quantity,
		Unit:      sale.Unit,
		Price:     sale.Price,
		SoldBy:    sale.SoldBy,
	}

	result, err := d.create.AddSaleRecord(ctx, saleObj)
	if err != nil {
		return nil, err
	}

	return &domain.Sale{
		ID:        result.ID,
		ProductID: result.ProductID,
		Quantity:  result.Quantity,
		Unit:      result.Unit,
		Price:     result.Price,
		SoldBy:    result.SoldBy,
	}, nil
}
