package gotosee

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

const (
	TYPE_TEXT = iota
	TYPE_OBJECT
	TYPE_READER
)

type RequestBody struct {
	bodyType int
	object   interface{}
	text     string
	reader   io.ReadCloser
}

func NewRequestBodyWithText(text string) *RequestBody {
	return &RequestBody{bodyType: TYPE_TEXT, text: text}
}

func NewRequestBodyWithObject(obj interface{}) *RequestBody {
	return &RequestBody{bodyType: TYPE_OBJECT, object: obj}
}

func NewRequestBodyWithReader(reader io.ReadCloser) *RequestBody {
	return &RequestBody{bodyType: TYPE_READER, reader: reader}
}

func (this *RequestBody) ConvertToReader(contentType string) (body io.Reader, err error) {
	switch this.bodyType {
	case TYPE_OBJECT:
		if this.object == nil {
			return
		} else {
			switch contentType {
			case CONTENT_TYPE_APPLICATION_JSON:
				data, err := json.Marshal(this.object)
				if err != nil {
					break
				}
				body = bytes.NewReader(data)
			case CONTENT_TYPE_APPLICATION_FORM_URLENCODED:

			}

		}
	case TYPE_READER:
		return this.reader, nil
	case TYPE_TEXT:
		return strings.NewReader(this.text), nil
	}
	return
}
