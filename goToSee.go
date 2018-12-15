package gotosee

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GoToSee struct {
	client *http.Client
}

func NewGoToSee() *GoToSee {
	return &GoToSee{client: &http.Client{}}
}

func (this *GoToSee) GetJson(url string, v interface{}) (err error) {
	body, err := this.GetBytes(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, v)
	return
}

func (this *GoToSee) GetBytes(url string) (body []byte, err error) {
	resp, err := this.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func (this *GoToSee) Get(url string) (resp *http.Response, err error) {
	return this.client.Get(url)
}

func (this *GoToSee) Head(url string, options interface{}) (resp *http.Response, err error) {
	return this.client.Head(url)
}

func (this *GoToSee) Options(url string, options interface{}) (resp *http.Response, err error) {
	return
}

func (this *GoToSee) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return this.client.Post(url, contentType, body)
}

func (this *GoToSee) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return this.client.PostForm(url, data)
}

func (this *GoToSee) Patch(url string, options interface{}) (resp *http.Response, err error) {
	return
}

func (this *GoToSee) Delete(url string, options interface{}) (resp *http.Response, err error) {

	return
}
