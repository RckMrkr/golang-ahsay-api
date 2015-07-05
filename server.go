package ahsay

import (
	"encoding/xml"
	"fmt"
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

func request(s Server, data map[string]string, ep string, obj interface{}) <-chan Response {
	r = make(chan interface{})

	go func() {
		url := createUrl(s, ep)
		values := createValues(s, data)
		body, err := callEndpoint(url, values)
		if err != nil {
			return response{Err: err}
		}

		err = unmarshal(body, &obj)

		return response{Object: obj, Err: err}
	}()

	return r
}

func createUrl(s Server, ep string) string {
	return fmt.Sprintf("%s/obs/api/", s.Host(), ep)
}

func createValues(s Server, data map[string]string) (values url.Values) {
	for k, v := range(data){
		values.Add()
	}
	values.Set("SysUser", s.Username())
	values.Set("SysPwd", s.Password())
	return
}

func callEndpoint(url string, values url.Values) string, error {
	resp, err := http.PostForm(url, data)
	return resp, err
}

func unmarshal(xml string, obj interface{}) interface {
	// Missing
}
