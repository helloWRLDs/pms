package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RequestBuilder struct {
	method  string
	url     string
	headers map[string]string
	query   map[string]string
	body    []byte

	httpClient *http.Client
	req        *http.Request
}

func (rb RequestBuilder) String() string {
	return fmt.Sprintf(
		"\nMethod:\t\t%s\nURL:\t\t%s\nHeaders:\t\t%v\nQuery:\t\t%v\nBody:\t\t%s",
		rb.method, rb.url, rb.headers, rb.query, string(rb.body),
	)
}

func New() *RequestBuilder {
	return &RequestBuilder{
		headers: map[string]string{
			"Content-Type": "application/json",
		},
		query:      make(map[string]string, 0),
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (rb *RequestBuilder) Body(data any) *RequestBuilder {
	switch v := data.(type) {
	case []byte:
		rb.body = v
	case string:
		rb.body = []byte(v)
	default:
		body, err := json.Marshal(v)
		if err == nil {
			rb.body = body
		}
	}
	return rb
}

func (rb *RequestBuilder) Method(method string) *RequestBuilder {
	rb.method = method
	return rb
}

func (rb *RequestBuilder) URL(url string) *RequestBuilder {
	rb.url = url
	return rb
}

func (rb *RequestBuilder) Headers(args ...string) *RequestBuilder {
	for i := 0; i < len(args)-1; i += 2 {
		rb.headers[args[i]] = args[i+1]
	}
	return rb
}

func (rb *RequestBuilder) Query(args ...any) *RequestBuilder {
	for i := 0; i < len(args)-1; i += 2 {
		rb.query[fmt.Sprintf("%v", args[i])] = fmt.Sprintf("%v", args[i+1])
	}
	return rb
}

func (rb *RequestBuilder) Client(client *http.Client) *RequestBuilder {
	rb.httpClient = client
	return rb
}

func (rb *RequestBuilder) Do() (response Response, err error) {
	if rb.req == nil {
		err := rb.build()
		if err != nil {
			return response, err
		}
	}

	res, err := rb.httpClient.Do(rb.req)
	if err != nil {
		return response, fmt.Errorf("failed to build request: %w", err)
	}
	defer res.Body.Close()
	response.Status = res.StatusCode
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("failed reading response body: %w", err)
	}
	response.Data = body
	return response, nil
}

func (rb *RequestBuilder) build() error {
	u, err := url.Parse(rb.url)
	if err != nil {
		return fmt.Errorf("failed parsing URL: %w", err)
	}
	q := u.Query()
	for k, v := range rb.query {
		if k != "" && v != "" {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(rb.method, u.String(), bytes.NewBuffer(rb.body))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	for k, v := range rb.headers {
		if k != "" && v != "" {
			req.Header.Set(k, v)
		}
	}
	rb.req = req
	return nil
}
