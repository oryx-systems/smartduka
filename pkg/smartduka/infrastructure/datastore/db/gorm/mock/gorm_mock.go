package mock

// GormDatastoreMock is a mocks the database layer
type GormDatastoreMock struct {
}

// NewGormDatastoreMock initializes a new GormDatastoreMock
func NewGormDatastoreMock() *GormDatastoreMock {
	// UUID := uuid.New().String()

	// residence := &gorm.Residence{
	// 	ID:                 UUID,
	// 	Active:             true,
	// 	Name:               gofakeit.Name(),
	// 	RegistrationNumber: gofakeit.Name(),
	// 	Location:           gofakeit.Name(),
	// 	LivingRoomsCount:   10,
	// 	Owner:              gofakeit.Name(),
	// }

	// contact := &gorm.Contact{
	// 	ID:           UUID,
	// 	Active:       true,
	// 	ContactType:  "PHONE",
	// 	ContactValue: gofakeit.Phone(),
	// 	Flavour:      enums.FlavourPro,
	// 	UserID:       &UUID,
	// }

	// user := &gorm.User{
	// 	ID:             &UUID,
	// 	FirstName:      gofakeit.BeerAlcohol(),
	// 	MiddleName:     gofakeit.BeerAlcohol(),
	// 	LastName:       gofakeit.BeerAlcohol(),
	// 	Active:         true,
	// 	Flavour:        enums.FlavourPro,
	// 	UserName:       gofakeit.BeerAlcohol(),
	// 	UserType:       "STAFF",
	// 	DeviceToken:    gofakeit.BeerAlcohol(),
	// 	Residence:      uuid.New().String(),
	// 	UserContact:    *contact,
	// }

	return &GormDatastoreMock{}
}
