package ahsay

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserTypeString(t *testing.T) {
	assert := assert.New(t)

	u := *new(UserType)
	assert.Equal("User type not set", fmt.Sprintf("%v", u))

	u = Paid
	assert.Equal("Paid", fmt.Sprintf("%v", u))

	u = Trial
	assert.Equal("Trial", fmt.Sprintf("%v", u))
}

func TestUserTypeUnmarshalPaid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		UserType UserType `xml:",attr"`
	}{}

	str := `
	<body UserType="PAID" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(Paid, obj.UserType)
}

func TestUserTypeUnmarshalTrial(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		UserType UserType `xml:",attr"`
	}{}

	str := `
	<body UserType="TRIAL" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(Trial, obj.UserType)
}

func TestUserTypeUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		UserType UserType `xml:",attr"`
	}{}

	str := `
	<body UserType="Invalid value" />
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(UserType), obj.UserType)
}
