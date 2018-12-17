package gotosee

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// ******************************** GET ********************************
func (this *GoToSee) Get(url string) (resp *http.Response, err error) {
	return this.client.Get(url)
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

func (this *GoToSee) GetObject(url string, res interface{}) (err error) {
	body, err := this.GetBytes(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, res)
	return
}

// ******************************** POST ********************************
func (this *GoToSee) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return this.client.Post(url, contentType, body)
}

func (this *GoToSee) PostJSON(url string, req interface{}) (resp *http.Response, err error) {
	body, err := json.Marshal(req)
	if err != nil {
		return
	}
	resp, err = this.Post(url, CONTENT_TYPE_APPLICATION_JSON, bytes.NewReader(body))
	return
}

func (this *GoToSee) PostObjectGetBytes(url string, req interface{}) (body []byte, err error) {
	resp, err := this.PostJSON(url, req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println("> ", string(body))
	return
}

func (this *GoToSee) PostObjectGetObject(url string, req interface{}, res interface{}) (err error) {
	body, err := this.PostObjectGetBytes(url, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, res)
	return
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

func (this *GoToSee) Head(url string, options interface{}) (resp *http.Response, err error) {
	return this.client.Head(url)
}

func (this *GoToSee) Options(url string, options interface{}) (resp *http.Response, err error) {
	return
}
