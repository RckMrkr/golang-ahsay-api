package ahsay

import (
	"encoding/xml"
	"errors"
)

// Status is used to idenify the status of a user
type Status uint8

// StatusEnabled means a user is activated
// StatusSuspended means a user is deactivated
const (
	StatusEnabled Status = iota + 1
	StatusSuspended
)

func (status Status) String() string {
	switch status {
	case StatusEnabled:
		return "Enabled"
	case StatusSuspended:
		return "Suspended"
	default:
		return "Status not set"
	}
}

func (status Status) toBool() (bool, error) {
	switch status {
	case StatusEnabled:
		return true, nil
	case StatusSuspended:
		return false, nil
	default:
		return false, errors.New("Status not set")
	}
}

// UnmarshalXMLAttr ensures Status implements the xml.UnmarshalerAttr interface
func (status *Status) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "ENABLE" {
		*status = StatusEnabled
	} else if attr.Value == "SUSPENDED" {
		*status = StatusSuspended

	}
	return nil
}
