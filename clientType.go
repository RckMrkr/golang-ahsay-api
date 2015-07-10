package ahsay

import "encoding/xml"

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
