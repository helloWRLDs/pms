package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Details struct {
	Method  string
	URL     string
	Headers map[string]string
	Query   map[string]string
	Body    []byte

	r *http.Request
}

func (d *Details) httpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func (d *Details) Build() error {
	u, err := url.Parse(d.URL)
	if err != nil {
		return fmt.Errorf("failed parsing URL: %w", err)
	}
	q := u.Query()
	for k, v := range d.Query {
		if k != "" && v != "" {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(d.Method, u.String(), bytes.NewBuffer(d.Body))
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	for k, v := range d.Headers {
		if k != "" && v != "" {
			req.Header.Add(k, v)
		}
	}
	d.r = req

	return nil
}

func (d *Details) Make() (int, []byte, error) {
	if d.r == nil {
		err := d.Build()
		if err != nil {
			return -1, nil, err
		}
	}

	res, err := d.httpClient().Do(d.r)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, fmt.Errorf("failed reading response body: %w", err)
	}
	return res.StatusCode, body, nil
}
