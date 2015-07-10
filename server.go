package ahsay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Response is a way to communicate via channels both objects and possible errors
type response struct {
	Body []byte
	Err  error
}

// Server defines the properties a request need to talk to a specific server
type Server interface {
	Host() string // entire part before "path" - i.e. http://www.google:com:8080
	Username() string
	Password() string
}

func request(s Server, data map[string]string, url string) <-chan response {
	c := make(chan response)

	go func() {
		values := createValues(s, data)
		body, err := callEndpoint(url, values)
		c <- response{Body: body, Err: err}
	}()

	return c
}

func createURL(s Server, ep string) string {
	return fmt.Sprintf("%s/obs/api/%s", s.Host(), ep)
}

func createValues(s Server, data map[string]string) url.Values {
	values := make(url.Values)
	for k, v := range data {
		values.Add(k, v)
	}
	values.Set("SysUser", s.Username())
	values.Set("SysPwd", s.Password())
	return values
}

func callEndpoint(url string, values url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
