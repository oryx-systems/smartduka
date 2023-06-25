package enums

import (
	"fmt"
	"io"
	"strconv"
)

// Category determines the 'category' of a product
type Category string

const (
	//CategoryCereals represents cereals product type
	CategoryCereals Category = "CEREALS"

	// CategoryMedicine represents the medicine product category type
	CategoryMedicine Category = "MEDICINE"

	// CategoryFoodStuff represents the food stuff category type
	CategoryFoodStuff Category = "FOOD_STUFF"
)

// IsValid returns true if a Category type is valid
func (u Category) IsValid() bool {
	switch u {
	case CategoryCereals, CategoryMedicine, CategoryFoodStuff:
		return true
	}
	return false
}

func (u Category) String() string {
	return string(u)
}

// UnmarshalGQL converts the supplied value to a category type.
func (u *Category) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*u = Category(str)
	if !u.IsValid() {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}

// MarshalGQL writes the Category type to the supplied
func (u Category) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(u.String()))
}
