package ahsay

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

type UserType int

const (
	PAID  UserType = iota + 1
	TRIAL UserType = iota + 1
)

func (t UserType) String() string {
	switch t {
	case PAID:
		return "Paid"
	case TRIAL:
		return "Trial"
	default:
		return "User type not set"
	}
}

func (t *UserType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	if s == "PAID" {
		*t = PAID
	} else if s == "TRIAL" {
		*t = TRIAL
	}

	return nil
}

type ClientType int

const (
	OBM ClientType = iota + 1
	ACB ClientType = iota + 1
)

func (t ClientType) String() string {
	switch t {
	case OBM:
		return "OBM"
	case ACB:
		return "ACB"
	default:
		return "Client type not set"
	}
}

func (t *ClientType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	if s == "ACB" {
		*t = ACB
	} else if s == "OBM" {
		*t = OBM
	}

	return nil
}

type ByteSize int

func (s ByteSize) String() string {
	i := int(s)
	div := i / 1024
	var unit string
	switch div {
	case 0:
		unit = "B"
	case 1:
		unit = "KB"
	case 2:
		unit = "MB"
	case 3:
		unit = "GB"
	case 4:
		unit = "TB"
	case 5:
		unit = "PB"
	case 6:
		unit = "EB" // Think we are save for a while now
	default:
		return fmt.Sprintf("%d%s", i, "B")
	}

	return fmt.Sprintf("%d%s", math.Mod(i, 1024), unit)
}

func (size *ByteSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	i, err = strconv.parseInt(s, 10)
	if err != nil {
		return err
	}
	*size = ByteSize(i)

	return nil
}

type Boolean bool

func (b Boolean) String() string {
	t = b.(bool)
	if t {
		return "True"
	} else {
		return "False"
	}
}

func (size *ByteSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	if s == "Y" {
		*size = true
	} else if s == "N" {
		*size = false

	}
	return nil
}

type Status bool

func (b Status) String() string {
	t = b.(bool)
	if t {
		return "Enabled"
	} else {
		return "Suspended"
	}
}

func (size *Status) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	if s == "ENABLE" {
		*size = true
	} else if s == "SUSPENDED" {
		*size = false

	}
	return nil
}

type Timestamp time.Time

func (s Timestamp) String() string {
	s = b.(time.Time)
	return s.String()
}

func (s *Timestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	i, err = strconv.parseInt(s, 10)
	if err != nil {
		return err
	}
	sec := i / 1000                       // The time we get is in ms
	ns := math.Mod(i, 1000) * 1000 * 1000 // ms to ns

	*s = time.Unix(sec, ns)
	return nil
}
