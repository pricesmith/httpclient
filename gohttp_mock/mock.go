package gohttp_mock

import (
	"fmt"
	"net/http"

	"github.com/opnscty/go-httpclient/core"
)

// Mock provides a clean way to configure http mocks based on a combination
// of a Request Method, URL, and Request Body.
type Mock struct {
	Method      string
	URL         string
	RequestBody string
	// [Todo] Headers?

	ResponseStatusCode int
	ResponseBody       string
	Error              error
}

// GetResponse returns a Response object based on the Mock configuration.
func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}

	return &response, nil
}
