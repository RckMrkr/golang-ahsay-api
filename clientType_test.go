package ahsay

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientTypeString(t *testing.T) {
	assert := assert.New(t)

	c := *new(ClientType)
	assert.Equal("Client type not set", fmt.Sprintf("%v", c))

	c = Obm
	assert.Equal("OBM", fmt.Sprintf("%v", c))

	c = Acb
	assert.Equal("ACB", fmt.Sprintf("%v", c))
}

func TestClientTypeUnmarshalObm(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ClientType ClientType `xml:",attr"`
	}{}

	str := `
		<body ClientType="OBM" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(Obm, obj.ClientType)
}

func TestClientTypeUnmarshalAcb(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ClientType ClientType `xml:",attr"`
	}{}

	str := `
		<body ClientType="ACB" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(Acb, obj.ClientType)
}

func TestClientTypeUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ClientType ClientType `xml:",attr"`
	}{}

	str := `
		<body ClientType="Invalid value" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(ClientType), obj.ClientType)
}
