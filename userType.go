package ahsay

import "encoding/xml"

// UserType is used to keep track of the type of user
type UserType int

// Paid is the UserType for users who have paid
// Trial is the UserType for users who are still on trial
const (
	Paid UserType = iota + 1
	Trial
)

func (t UserType) String() string {
	switch t {
	case Paid:
		return "Paid"
	case Trial:
		return "Trial"
	default:
		return "User type not set"
	}
}

// UnmarshalXMLAttr ensures UserType implements the xml.UnmarshalerAttr interface
func (t *UserType) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "PAID" {
		*t = Paid
	} else if attr.Value == "TRIAL" {
		*t = Trial
	}

	return nil
}
