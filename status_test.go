package ahsay

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatusString(t *testing.T) {
	assert := assert.New(t)

	// Try each standard case
	s := S_ENABLED
	assert.Equal("Enabled", fmt.Sprintf("%v", s))

	s = S_SUSPENDED
	assert.Equal("Suspended", fmt.Sprintf("%v", s))

	s = *new(Status)
	assert.Equal("Status not set", fmt.Sprintf("%v", s))
}

func TestStatusUnmarshalEnabled(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Status Status
	}{}

	str := `
	<body>
		<Status>ENABLE</Status>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(S_ENABLED, obj.Status)
}

func TestStatusUnmarshalFalse(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Status Status
	}{}

	str := `
	<body>
		<Status>SUSPENDED</Status>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(S_SUSPENDED, obj.Status)
}

func TestStatusUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		Status Status
	}{}

	str := `
	<body>
		<Status>Invalid input</Status>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(Status), obj.Status)
}

func TestStatusToBool(t *testing.T) {
	assert := assert.New(t)

	s, err := S_ENABLED.toBool()
	assert.Nil(err)
	assert.True(s)

	s, err = S_SUSPENDED.toBool()
	assert.Nil(err)
	assert.False(s)

	_, err = (*new(Status)).toBool()
	if assert.NotNil(err) {
		assert.Error(err)
	}
}
