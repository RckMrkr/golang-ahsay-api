package ahsay

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimestampString(t *testing.T) {
	assert := assert.New(t)

	s := time.Now()
	ts := Timestamp(s)

	assert.Equal(fmt.Sprintf("%v", s), fmt.Sprintf("%v", ts))
}

func TestTimestampUnmarshal(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		Timestamp Timestamp
	}{}

	str := `
	<body>
		<Timestamp>1302687743242</Timestamp>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	s := time.Unix(1302687743, 242*1000*1000)
	timeObj := time.Time(obj.Timestamp)
	timesAreEqual := s.Equal(timeObj)
	assert.True(timesAreEqual)
}

func TestTimestampUnmarshalInvalid(t *testing.T) {
	assert := assert.New(t)
	obj := struct {
		Timestamp Timestamp
	}{}

	str := `
	<body>
		<Timestamp>Invalid input</Timestamp>
	</body>
	`
	b := []byte(str)
	xml.Unmarshal(b, &obj)
	assert.Equal(*new(Timestamp), obj.Timestamp)
}
