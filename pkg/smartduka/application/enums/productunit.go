package enums

import (
	"fmt"
	"io"
	"strconv"
)

// Unit determines the 'unit' of sale of a product
type Unit string

const (
	//UnitSingle represents the admin app
	UnitSingle Unit = "ONE"

	// UnitHalfDozen represents the half dozen unit of ocean
	UnitHalfDozen Unit = "HALF_DOZEN"

	// UnitDozen represents product dozen unit
	UnitDozen Unit = "DOZEN"

	// UnitOuter represents product outer unit
	UnitOuter Unit = "OUTER"

	// UnitCarton represents product carton unit
	UnitCarton Unit = "CARTON"

	// UnitBale represents product bale unit
	UnitBale Unit = "BALE"

	// UnitBag represents product bag unit
	UnitBag Unit = "BAG"

	// UnitPacket represents product packet unit
	UnitPacket Unit = "PACKET"
)

// IsValid returns true if a Unit type is valid
func (u Unit) IsValid() bool {
	switch u {
	case UnitSingle, UnitHalfDozen, UnitDozen, UnitOuter, UnitCarton, UnitBale, UnitBag, UnitPacket:
		return true
	}
	return false
}

func (u Unit) String() string {
	return string(u)
}

// UnmarshalGQL converts the supplied value to a Unit type.
func (u *Unit) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*u = Unit(str)
	if !u.IsValid() {
		return fmt.Errorf("%s is not a valid Unit", str)
	}
	return nil
}

// MarshalGQL writes the Unit type to the supplied
func (u Unit) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(u.String()))
}
