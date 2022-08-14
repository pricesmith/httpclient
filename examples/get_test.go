package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/opnscty/go-httpclient/gohttp_mock"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting test cases for package 'examples'...")

	// Tell HTTP library to mock all following requests
	gohttp_mock.MockupServer.Start()

	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {

		// Initialization:
		gohttp_mock.MockupServer.DeleteMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method: http.MethodGet,
			URL:    "https://api.github.com",
			Error:  errors.New("Request Timeout: Could not fetch Github endpoints"),
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation:
		if endpoints != nil {
			t.Error("No endpoints expected.")
		}

		if err == nil {
			t.Error("An error was expected.")
		}

		if err.Error() != "Request Timeout: Could not fetch Github endpoints" {
			t.Error("Invalid error message received.")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttp_mock.MockupServer.DeleteMocks()

		// Initialization:
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			URL:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation:
		if endpoints != nil {
			t.Error("No endpoints expected.")
		}

		if err == nil {
			t.Error("An error was expected.")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("Invalid error message received.")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp_mock.MockupServer.DeleteMocks()

		// Initialization:
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			URL:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "http://api.github.com/user"}`,
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation:
		if err != nil {
			t.Error(fmt.Sprintf("No error was expected. Received: '%s'", err.Error()))
		}

		if endpoints == nil {
			t.Error("Endpoints were expected. Nil was returned.")
		}

		if endpoints.CurrentUser != "http://api.github.com/user" {
			t.Error("Invalid current user url.")
		}
	})
}
