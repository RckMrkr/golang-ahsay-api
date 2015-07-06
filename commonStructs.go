package ahsay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

type UserType int

const (
	PAID UserType = iota + 1
	TRIAL
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
	ACB
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

type ByteSize uint64

func (s ByteSize) String() string {
	prefixes := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	format := "%.2f %s"
	value := float64(s)
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

func (size *ByteSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*size = ByteSize(i)

	return nil
}

type Boolean uint8

const (
	B_TRUE Boolean = iota + 1
	B_FALSE
)

func (b Boolean) String() string {
	switch b {
	case B_TRUE:
		return "True"
	case B_FALSE:
		return "False"
	default:
		return "Not set"
	}
}

func (b Boolean) toBool() (bool, error) {
	switch b {
	case B_TRUE:
		return true, nil
	case B_FALSE:
		return false, nil
	default:
		return false, errors.New("Boolean not set")
	}
}

func (size *Boolean) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	if s == "Y" {
		*size = B_TRUE
	} else if s == "N" {
		*size = B_FALSE

	}
	return nil
}

type Status uint8

const (
	S_ENABLED Status = iota + 1
	S_SUSPENDED
)

func (b Status) String() string {
	switch b {
	case S_ENABLED:
		return "Enabled"
	case S_SUSPENDED:
		return "Suspended"
	default:
		return "Status not set"
	}
}

func (b Status) toBool() (bool, error) {
	switch b {
	case S_ENABLED:
		return true, nil
	case S_SUSPENDED:
		return false, nil
	default:
		return false, errors.New("Status not set")
	}
}

func (size *Status) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	if s == "ENABLE" {
		*size = S_ENABLED
	} else if s == "SUSPENDED" {
		*size = S_SUSPENDED

	}
	return nil
}

type Timestamp time.Time

func (t Timestamp) String() string {
	s := time.Time(t)
	return s.String()
}

func (t *Timestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	d.DecodeElement(&s, &start)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	sec := i / 1000                // The time we get is in ms
	ns := (i % 1000) * 1000 * 1000 // ms to ns

	time := time.Unix(sec, ns)
	*t = Timestamp(time)
	return nil
}
