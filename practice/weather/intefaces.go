package weather

import "net/http"

type httpClient interface {
	PerformRequest(rdmap map[string]interface{}) (*http.Response, error)
}
