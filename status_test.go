package ahsay

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusString(t *testing.T) {
	assert := assert.New(t)

	// Try each standard case
	s := StatusEnabled
	assert.Equal("Enabled", fmt.Sprintf("%v", s))

	s = StatusSuspended
	assert.Equal("Suspended", fmt.Sprintf("%v", s))

	s = *new(Status)
	assert.Equal("Status not set", fmt.Sprintf("%v", s))
}

func TestStatusUnmarshalEnabled(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Status Status `xml:",attr"`
	}{}

	str := `
		<body Status="ENABLE" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(StatusEnabled, obj.Status)
}

func TestStatusUnmarshalFalse(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Status Status `xml:",attr"`
	}{}

	str := `
		<body Status="SUSPENDED" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(StatusSuspended, obj.Status)
}

func TestStatusUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		Status Status `xml:",attr"`
	}{}

	str := `
		<body Status="Invalid input" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(Status), obj.Status)
}

func TestStatusToBool(t *testing.T) {
	assert := assert.New(t)

	s, err := StatusEnabled.toBool()
	assert.Nil(err)
	assert.True(s)

	s, err = StatusSuspended.toBool()
	assert.Nil(err)
	assert.False(s)

	_, err = (*new(Status)).toBool()
	if assert.NotNil(err) {
		assert.Error(err)
	}
}
