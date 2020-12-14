package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_PerformRequest_GET(t *testing.T) {
	var expectedBody = []byte(`{"name": "something"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "okHeader", r.Header.Get("testHeader"))
		assert.Equal(t, "okQuery", r.URL.Query().Get("testQueryParam"))
		_, _ = w.Write(expectedBody)
	}))
	defer ts.Close()

	client := NewHttpClient(3 * time.Second)
	rdmap := map[string]interface{}{
		"verb":        http.MethodGet,
		"url":         ts.URL,
		"queryParams": map[string]string{"testQueryParam": "okQuery"},
		"headers":     map[string]string{"testHeader": "okHeader"},
	}

	response, err := client.PerformRequest(rdmap)
	if err != nil {
		t.Errorf("unexpected error occured while making request with data %v", rdmap)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("unexpected error occured while reading the response body")
	}
	assert.Equal(t, expectedBody, responseBody)
}

func Test_PerformRequest_POST(t *testing.T) {
	var postBody = []byte(`{"testField": "testingPOST"}`)

	var expectedBody = []byte(`{"name": "something"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unexpected error when reading the body on mocked server side")
		}

		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, postBody, body)
		assert.Equal(t, "okHeader", r.Header.Get("testHeader"))
		assert.Equal(t, "okQuery", r.URL.Query().Get("testQueryParam"))
		_, _ = w.Write(expectedBody)
	}))
	defer ts.Close()

	client := NewHttpClient(3 * time.Second)
	rdmap := map[string]interface{}{
		"verb":        http.MethodPost,
		"url":         ts.URL,
		"body":        ioutil.NopCloser(bytes.NewBuffer(postBody)),
		"queryParams": map[string]string{"testQueryParam": "okQuery"},
		"headers":     map[string]string{"testHeader": "okHeader"},
	}

	response, err := client.PerformRequest(rdmap)
	if err != nil {
		t.Errorf("unexpected error occured while making request with data %v", rdmap)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("unexpected error occured while reading the response body")
	}
	assert.Equal(t, expectedBody, responseBody)
}

func Test_validate(t *testing.T) {
	type fields struct {
		cli *http.Client
	}
	type args struct {
		rd *requestData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid request data",
			fields: fields{
				cli: &http.Client{},
			},
			args: args{
				rd: &requestData{
					Verb: http.MethodGet,
					URL:  "http://www.google.com",
				},
			},
			wantErr: false,
		},
		//{
		//	name: "invalid request data, invalid URL",
		//	fields: fields{
		//		cli: &http.Client{},
		//	},
		//	args: args{
		//		rd: &requestData{
		//			Verb: http.MethodGet,
		//			URL:  "//www.google",
		//		},
		//	},
		//	wantErr: true,
		//},
		{
			name: "invalid request data, not provided verb",
			fields: fields{
				cli: &http.Client{},
			},
			args: args{
				rd: &requestData{
					URL: "http://www.google.com",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid request data, not provided url",
			fields: fields{
				cli: &http.Client{},
			},
			args: args{
				rd: &requestData{
					Verb: http.MethodHead,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HttpClient{
				cli: tt.fields.cli,
			}
			if err := c.validate(tt.args.rd); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isValidUrl(t *testing.T) {
	tests := []struct {
		name    string
		URL     string
		wantErr bool
	}{
		{
			name:    "valid url",
			URL:     "http://www.google.com.ar",
			wantErr: false,
		},
		{
			name:    "invalid URL, incomplete url",
			URL:     "www.googl",
			wantErr: true,
		},
		{
			name:    "invalid URL, no schema",
			URL:     "www.google.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidUrl(tt.URL); (err != nil) != tt.wantErr {
				t.Errorf("isValidUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
