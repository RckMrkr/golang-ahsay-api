package ahsay

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserTypeString(t *testing.T) {
	assert := assert.New(t)

	u := *new(UserType)
	assert.Equal("User type not set", fmt.Sprintf("%v", u))

	u = PAID
	assert.Equal("Paid", fmt.Sprintf("%v", u))

	u = TRIAL
	assert.Equal("Trial", fmt.Sprintf("%v", u))
}

func TestUserTypeUnmarshalPaid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		UserType UserType
	}{}

	str := `
	<body>
		<UserType>PAID</UserType>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(PAID, obj.UserType)
}

func TestUserTypeUnmarshalTrial(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		UserType UserType
	}{}

	str := `
	<body>
		<UserType>TRIAL</UserType>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(TRIAL, obj.UserType)
}

func TestUserTypeUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		UserType UserType
	}{}

	str := `
	<body>
		<UserType>INVALID VALUE</UserType>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(UserType), obj.UserType)
}
