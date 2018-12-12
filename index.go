package gotosee

import (
	"io"
	"net/http"
	"net/url"
)

func Get(url string, options interface{}) (resp *http.Response, err error) {
	return http.Get(url)
}

func Head(url string, options interface{}) (resp *http.Response, err error) {
	return http.Head(url)
}

func Options(url string, options interface{}) (resp *http.Response, err error) {
	return
}

func Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return http.Post(url, contentType, body)
}

func PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return http.PostForm(url, data)
}

func Patch(url string, options interface{}) (resp *http.Response, err error) {
	return
}

func Delete(url string, options interface{}) (resp *http.Response, err error) {

	return
}
