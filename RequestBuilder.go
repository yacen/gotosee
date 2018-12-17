package gotosee

import (
	"errors"
	"io"
	"net/http"
	urlLib "net/url"
)

type RequestBuilder struct {
	URL       string
	Method    string
	MediaType string
	Headers   http.Header
	Body      *RequestBody
}

func NewRequestBuilder(request ...*http.Request) {
	reqBuilder := &RequestBuilder{}
	if len(request) == 0 {
		reqBuilder.Method = METHOD_GET
		reqBuilder.Headers = make(http.Header)
	} else {
		req := request[0]
		reqBuilder.URL = req.URL.String()
		reqBuilder.Method = req.Method
		reqBuilder.Body = NewRequestBodyWithReader(req.Body)
		reqBuilder.Headers = req.Header
	}
}

/**
 * Sets the url target of the request
 *
 */

func (builder *RequestBuilder) Url(url *urlLib.URL) *RequestBuilder {
	if url == nil {
		panic(errors.New("url == nil"))
	}
	builder.URL = url.String()
	return builder
}

/**
 * Sets the url string target of the request
 *
 */
func (builder *RequestBuilder) UrlStr(url string) *RequestBuilder {
	if url == "" {
		panic(errors.New("url == ``"))
	}
	builder.URL = url
	return builder
}

/**
 * Sets the header named {@code name} to {@code value}. If this request already has any headers
 * with that name, they are all replaced.
 */
func (builder *RequestBuilder) SetHeader(name string, value string) *RequestBuilder {
	builder.Headers.Set(name, value)
	return builder
}

/**
 * Adds a header with {@code name} and {@code value}. Prefer this method for multiply-valued
 * headers like "Cookie".
 *
 * <p>Note that for some headers including {@code Content-Length} and {@code Content-Encoding},
 * OkHttp may replace {@code value} with a header derived from the request body.
 */
func (builder *RequestBuilder) AddHeader(name string, value string) *RequestBuilder {
	builder.Headers.Add(name, value)
	return builder
}

/** Removes all headers named {@code name} on this builder. */
func (builder *RequestBuilder) RemoveHeader(name string) *RequestBuilder {
	builder.Headers.Del(name)
	return builder
}

/** Removes all headers on this builder and adds {@code header}. */
func (builder *RequestBuilder) Header(header http.Header) *RequestBuilder {
	builder.Headers = CloneHeader(header)
	return builder
}

/**
 * Sets this request's {@code Cache-Control} header, replacing any cache control headers already
 * present. If {@code cacheControl} doesn't define any directives, this clears this request's
 * cache-control headers.
 */
func (builder *RequestBuilder) CacheControl() *RequestBuilder {
	return builder
}

func (builder *RequestBuilder) ContentType(contentType string) *RequestBuilder {
	builder.MediaType = contentType
	return builder
}

func (builder *RequestBuilder) Get() *RequestBuilder {
	builder.Method = METHOD_GET
	return builder
}

func (builder *RequestBuilder) Head() *RequestBuilder {
	builder.Method = METHOD_HEAD
	return builder
}

func (builder *RequestBuilder) Post(body *RequestBody) *RequestBuilder {
	builder.Method = METHOD_POST
	builder.Body = body
	return builder
}

func (builder *RequestBuilder) Delete(body ...*RequestBody) *RequestBuilder {
	builder.Method = METHOD_DELETE
	if len(body) > 0 {
		builder.Body = body[0]
	}
	return builder
}

func (builder *RequestBuilder) Put(body *RequestBody) *RequestBuilder {
	builder.Method = METHOD_PUT
	builder.Body = body
	return builder
}

func (builder *RequestBuilder) PATCH(body *RequestBody) *RequestBuilder {
	builder.Method = METHOD_PATCH
	builder.Body = body
	return builder
}

func (builder *RequestBuilder) Build() (*http.Request, error) {
	method := builder.getMethod()
	url, err := builder.getUrl()
	body, err := builder.getBody()
	req, err := http.NewRequest(method, url, body)
	return req, err
}

func (builder *RequestBuilder) getMethod() string {
	method := builder.Method
	if method == "" {
		method = METHOD_GET
	}
	return method
}

func (builder *RequestBuilder) getUrl() (url string, err error) {
	url = builder.URL
	if url == "" {
		err = errors.New("url is empty")
	}
	return
}

func (builder *RequestBuilder) getBody() (body io.Reader, err error) {
	if builder.Method == METHOD_GET || builder.Body == nil {
		return
	}
	body, err = builder.Body.ConvertToReader(builder.MediaType)
	return
}
