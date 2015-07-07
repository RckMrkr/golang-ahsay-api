package ahsay

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type response struct {
	Object interface{}
	Err    error
}

type Server interface {
	Host() string // entire part before "path" - i.e. http://www.google:com:8080
	Username() string
	Password() string
}

func request(s Server, data map[string]string, ep string, obj interface{}) <-chan response {
	c := make(chan response)

	go func() {
		url := createUrl(s, ep)
		values := createValues(s, data)
		body, err := callEndpoint(url, values)
		if err != nil {
			c <- response{Err: err}
		}
		xml.Unmarshal(body, &obj)

		c <- response{Object: obj, Err: err}
	}()

	return c
}

func createUrl(s Server, ep string) string {
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
