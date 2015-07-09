package ahsay

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBooleanString(t *testing.T) {
	assert := assert.New(t)

	// Try each standard case
	s := BooleanTrue
	assert.Equal("True", fmt.Sprintf("%v", s))

	s = BooleanFalse
	assert.Equal("False", fmt.Sprintf("%v", s))

	s = *new(Boolean)
	assert.Equal("Not set", fmt.Sprintf("%v", s))
}

func TestBooleanUnmarshalTrue(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Boolean Boolean `xml:"Boolean,attr"`
	}{}

	str := `
	<body Boolean="Y" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(BooleanTrue, obj.Boolean)
}

func TestBooleanUnmarshalFalse(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Boolean Boolean `xml:"Boolean,attr"`
	}{}

	str := `
	<body Boolean="N">
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(BooleanFalse, obj.Boolean)
}

func TestBooleanUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		Boolean Boolean `xml:"Boolean,attr"`
	}{}

	str := `
	<body Boolean="Invalid input">
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(Boolean), obj.Boolean)
}

func TestBooleanToBool(t *testing.T) {
	assert := assert.New(t)

	b, err := BooleanTrue.ToBool()
	assert.Nil(err)
	assert.True(b)

	b, err = BooleanFalse.ToBool()
	assert.Nil(err)
	assert.False(b)

	_, err = (*new(Boolean)).ToBool()
	if assert.NotNil(err) {
		assert.Error(err)
	}
}
