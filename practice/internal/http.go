package internal

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"
)

// requestData wraps all the data needed to perform a request
type requestData struct {
	Verb        string            `mapstructure:"verb"`
	URL         string            `mapstructure:"url"`
	Body        io.Reader         `mapstructure:"body"`
	QueryParams map[string]string `mapstructure:"queryParams"`
	Headers     map[string]string `mapstructure:"headers"`
}

type HttpClient struct {
	cli *http.Client
}

// NewHttpClient instance for the httpclient
func NewHttpClient(t time.Duration) *HttpClient {
	return &HttpClient{
		cli: &http.Client{
			Timeout: t * time.Second,
		},
	}
}

func (hc *HttpClient) PerformRequest(rdmap map[string]interface{}) (*http.Response, error) {
	var rd requestData
	if err := mapstructure.Decode(rdmap, &rd); err != nil {
		return nil, err
	}

	return hc.doPerformRequest(&rd)
}

// doPerformRequest validates and triggers the HTTP call as per requestData parameter
func (hc *HttpClient) doPerformRequest(rd *requestData) (*http.Response, error) {
	if err := hc.validate(rd); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(rd.Verb, rd.URL, rd.Body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range rd.Headers {
		req.Header.Set(k, v)
	}

	// If querystring was provided
	if len(rd.QueryParams) > 0 {
		q := req.URL.Query()
		for k, v := range rd.QueryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	res, err := hc.cli.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := res.Body.Close(); err != nil {
		return nil, err
	}

	res.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return res, nil
}

// validate checks for the minimum mandatory request composition
func (hc *HttpClient) validate(rd *requestData) error {
	if rd.Verb == "" || rd.URL == "" {
		return errors.New("neither http Verb nor url must not be string zero value")
	}

	//if err := isValidUrl(rd.URL); err != nil {
	//	return fmt.Errorf("invalid provided url %v: %w", rd.URL, err)
	//}

	return nil
}

// isValidUrl tests a string to determine if it is a well-structured url or not.
func isValidUrl(URL string) error {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	u, err := url.Parse(URL)
	if err != nil {
		return err
	}

	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("there is an error within the URL composition")
	}

	return nil
}
