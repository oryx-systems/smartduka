package enums

import (
	"fmt"
	"io"
	"strconv"
)

// Flavour is a list of application types.
type Flavour string

const (
	//FlavourPro represents the admin app
	FlavourPro Flavour = "PRO"
	// FlavourConsumer represents the tenant app user
	FlavourConsumer Flavour = "CONSUMER"
)

// IsValid returns true if a flavour type is valid
func (f Flavour) IsValid() bool {
	switch f {
	case FlavourPro, FlavourConsumer:
		return true
	}
	return false
}

func (f Flavour) String() string {
	return string(f)
}

// UnmarshalGQL converts the supplied value to a flavour type.
func (f *Flavour) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*f = Flavour(str)
	if !f.IsValid() {
		return fmt.Errorf("%s is not a valid Flavour", str)
	}
	return nil
}

// MarshalGQL writes the flavour type to the supplied
func (f Flavour) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(f.String()))
}
