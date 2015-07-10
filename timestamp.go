package ahsay

import (
	"encoding/xml"
	"strconv"
	"time"
)

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
