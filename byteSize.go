package ahsay

import (
	"encoding/xml"
	"fmt"
	"math"
	"strconv"
)

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
