package ahsay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

// UserType is used to idenify the type of user
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

// ClientType is used to idenify the type of client
type ClientType int

// Obm is the ClientType for Client with OBM version
// Acb is the ClientType for Client with ACB version
const (
	Obm ClientType = iota + 1
	Acb
)

func (t ClientType) String() string {
	switch t {
	case Obm:
		return "OBM"
	case Acb:
		return "ACB"
	default:
		return "Client type not set"
	}
}

// UnmarshalXMLAttr ensures ClientType implements the xml.UnmarshalerAttr interface
func (t *ClientType) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "ACB" {
		*t = Acb
	} else if attr.Value == "OBM" {
		*t = Obm
	}

	return nil
}

// ByteSize is a representation of an amount of bytes
type ByteSize uint64

func (size ByteSize) String() string {
	prefixes := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	format := "%.2f %s"
	value := float64(size)
	var (
		div     float64
		divisor float64
	)
	for index, unit := range prefixes {
		divisor = math.Pow(1024, float64(index))
		div = value / divisor
		if div < 1024 {
			return fmt.Sprintf(format, div, unit)
		}
	}
	return fmt.Sprintf(format, value, prefixes[0])
}

// UnmarshalXMLAttr ensures ByteSize implements the xml.UnmarshalerAttr interface
func (size *ByteSize) UnmarshalXMLAttr(attr xml.Attr) error {
	i, err := strconv.ParseInt(attr.Value, 10, 64)
	if err != nil {
		return err
	}
	*size = ByteSize(i)

	return nil
}

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

// Timestamp is a faux class to enable unmarshalling of time.Time cases
type Timestamp time.Time

func (t Timestamp) String() string {
	s := time.Time(t)
	return s.String()
}

// UnmarshalXMLAttr ensures Timestamp implements the xml.UnmarshalerAttr interface
func (t *Timestamp) UnmarshalXMLAttr(attr xml.Attr) error {
	i, err := strconv.ParseInt(attr.Value, 10, 64)
	if err != nil {
		return err
	}
	sec := i / 1000                // The time we get is in ms
	ns := (i % 1000) * 1000 * 1000 // ms to ns

	time := time.Unix(sec, ns)
	*t = Timestamp(time)
	return nil
}
