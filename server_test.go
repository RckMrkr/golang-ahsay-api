package ahsay

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type servertest bool

func (s servertest) Username() string {
	return "username"
}
func (s servertest) Password() string {
	return "password"
}

func (s servertest) Host() string {
	return "http://example.org:8080"
}

func TestCreateValues(t *testing.T) {
	assert := assert.New(t)

	s := servertest(true)
	m := make(map[string]string)
	m["k1"] = "v1"
	m["k2"] = "v2"

	v := url.Values{}
	v.Add("k2", "v2")
	v.Add("k1", "v1")
	v.Add("SysPwd", "password")
	v.Add("SysUser", "username")
	assert.Equal(v, createValues(s, m))
}

func TestCreateUrl(t *testing.T) {
	assert := assert.New(t)

	s := servertest(true)
	assert.Equal("http://example.org:8080/obs/api/listUsers.do", createURL(s, "listUsers.do"))
}

func TestCallEndpoint(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "v1")
		fmt.Fprint(w, "v2")
	}))

	defer ts.Close()
	values := make(url.Values)
	values.Add("k1", "v1")
	values.Add("k2", "v2")
	body, err := callEndpoint(ts.URL, url.Values{})
	assert.Equal([]byte("v1\nv2"), body)
	assert.Nil(err)

	_, err = callEndpoint("invalid url", url.Values{})
	assert.NotNil(err)
}

func TestRequest(t *testing.T) {
	assert := assert.New(t)

	s := "Check"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s)
	}))

	defer ts.Close()
	server := servertest(true)
	r := <-request(server, make(map[string]string), ts.URL)

	assert.Nil(r.Err)
	assert.Equal([]byte(s), r.Body)
}

func TestRequestInvalid(t *testing.T) {
	assert := assert.New(t)

	server := servertest(true)
	r := <-request(server, make(map[string]string), "http://this.must.be.invalid.xaz")

	assert.NotNil(r.Err)
}
