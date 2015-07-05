package ahsay

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientTypeString(t *testing.T) {
	assert := assert.New(t)

	c := *new(ClientType)
	assert.Equal("Client type not set", fmt.Sprintf("%v", c))

	c = OBM
	assert.Equal("OBM", fmt.Sprintf("%v", c))

	c = ACB
	assert.Equal("ACB", fmt.Sprintf("%v", c))
}

func TestClientTypeUnmarshalObm(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ClientType ClientType
	}{}

	str := `
	<body>
		<ClientType>OBM</ClientType>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(OBM, obj.ClientType)
}

func TestClientTypeUnmarshalAcb(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ClientType ClientType
	}{}

	str := `
	<body>
		<ClientType>ACB</ClientType>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(ACB, obj.ClientType)
}

func TestClientTypeUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ClientType ClientType
	}{}

	str := `
	<body>
		<ClientType>INVALID VALUE</ClientType>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(ClientType), obj.ClientType)
}
