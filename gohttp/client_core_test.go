package gohttp

import (
	"testing"
)

func TestGetRequestBody(t *testing.T) {

	// Initialization
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("No error expected when passing nil body.")
		}

		if body != nil {
			t.Error("No body expected when passing nil body.")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		// Validation
		if err != nil {
			t.Error("No error expected when marshalling slice as json.")
		}

		if string(body) != `["one","two"]` {
			t.Error(("Invalid json body obtained."))
		}

	})

	// [Todo] - finish test cases
	// t.Run("BodyWithXml", func(t *testing.T) {
	// 	// Execution
	// 	body, err := client.getRequestBody("Content-Type", "application/xml")
	// })

	// t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
	// 	// Execution
	// 	body, err := client.getRequestBody("Content-Type", "application/json")
	// })
}
