package ahsay

import (
	"encoding/xml"
	"errors"
)

// Boolean is used as a faux class to unmarshal boolean values
type Boolean uint8

// BooleanTrue represents a boolean value of true
// BooleanFalse represents a boolean value of false
const (
	BooleanTrue Boolean = iota + 1
	BooleanFalse
)

func (b Boolean) String() string {
	switch b {
	case BooleanTrue:
		return "True"
	case BooleanFalse:
		return "False"
	default:
		return "Not set"
	}
}

// ToBool is able to transform the Boolean to a bool
func (b Boolean) ToBool() (bool, error) {
	switch b {
	case BooleanTrue:
		return true, nil
	case BooleanFalse:
		return false, nil
	default:
		return false, errors.New("Boolean not set")
	}
}

// UnmarshalXMLAttr ensures Boolean implements the xml.UnmarshalerAttr interface
func (b *Boolean) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "Y" {
		*b = BooleanTrue
	} else if attr.Value == "N" {
		*b = BooleanFalse

	}
	return nil
}
