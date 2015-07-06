package ahsay

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteSizeString(t *testing.T) {
	assert := assert.New(t)

	// Try each standard case
	s := ByteSize(943)
	assert.Equal("943.00 B", fmt.Sprintf("%v", s))

	s = ByteSize(9433)
	assert.Equal("9.21 KB", fmt.Sprintf("%v", s))

	s = ByteSize(94331234)
	assert.Equal("89.96 MB", fmt.Sprintf("%v", s))

	s = ByteSize(943312341234)
	assert.Equal("878.53 GB", fmt.Sprintf("%v", s))

	s = ByteSize(1044543312341234)
	assert.Equal("950.01 TB", fmt.Sprintf("%v", s))

	s = ByteSize(100004543312341234)
	assert.Equal("88.82 PB", fmt.Sprintf("%v", s))

	s = ByteSize(2305843009213693952)
	assert.Equal("2305843009213693952.00 B", fmt.Sprintf("%v", s))

	// Edge cases:
	s = ByteSize(1024)
	assert.Equal("1.00 KB", fmt.Sprintf("%v", s))

	s = ByteSize(0)
	assert.Equal("0.00 B", fmt.Sprintf("%v", s))
}

func TestByteSizeUnmarshalBytes(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ByteSize ByteSize
	}{}

	str := `
	<body>
		<ByteSize>425</ByteSize>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(ByteSize(425), obj.ByteSize)
}

func TestByteSizeUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		ByteSize ByteSize
	}{}

	str := `
	<body>
		<ByteSize>fdsa</ByteSize>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(ByteSize), obj.ByteSize)
}
