package examples

import (
	"fmt"
)

type Endpoints struct {
	CurrentUser       string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryURL     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com")
	if err != nil {
		// Deal with error as needed
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status))
	fmt.Println(fmt.Sprintf("Response Body: %s\n", response.String()))

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		// Deal with unmarshal error as needed
		return nil, err
	}

	fmt.Println((fmt.Sprintf("Repository URL: %s", endpoints.RepositoryURL)))
	return &endpoints, nil
}
