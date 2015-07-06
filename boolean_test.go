package ahsay

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooleanString(t *testing.T) {
	assert := assert.New(t)

	// Try each standard case
	s := B_TRUE
	assert.Equal("True", fmt.Sprintf("%v", s))

	s = B_FALSE
	assert.Equal("False", fmt.Sprintf("%v", s))

	s = *new(Boolean)
	assert.Equal("Not set", fmt.Sprintf("%v", s))
}

func TestBooleanUnmarshalTrue(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Boolean Boolean
	}{}

	str := `
	<body>
		<Boolean>Y</Boolean>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(B_TRUE, obj.Boolean)
}

func TestBooleanUnmarshalFalse(t *testing.T) {
	assert := assert.New(t)

	obj := struct {
		Boolean Boolean
	}{}

	str := `
	<body>
		<Boolean>N</Boolean>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(B_FALSE, obj.Boolean)
}

func TestBooleanUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		Boolean Boolean
	}{}

	str := `
	<body>
		<Boolean>Invalid input</Boolean>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(Boolean), obj.Boolean)
}

func TestBooleanToBool(t *testing.T) {
	assert := assert.New(t)

	b, err := B_TRUE.toBool()
	assert.Nil(err)
	assert.True(b)

	b, err = B_FALSE.toBool()
	assert.Nil(err)
	assert.False(b)

	_, err = (*new(Boolean)).toBool()
	if assert.NotNil(err) {
		assert.Error(err)
	}
}
