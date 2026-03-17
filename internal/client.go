package internal

import (
	"context"
	"fmt"

	authentik "goauthentik.io/api/v3"
)

type AuthentikClient struct {
	client *authentik.APIClient
}

func NewAuthentikClient(endpoint, token string) *AuthentikClient {
	config := authentik.NewConfiguration()
	config.Host = endpoint
	config.Scheme = "https"
	config.AddDefaultHeader("Authorization", "Bearer "+token)

	return &AuthentikClient{
		client: authentik.NewAPIClient(config),
	}
}

func (c *AuthentikClient) CreateUserRequest(pr *ProvisionRequest) (*authentik.User, error) {
	userRequest := *authentik.NewUserRequest(pr.Email, pr.FirstName+" "+pr.LastName)
	user, _, err := c.client.CoreApi.CoreUsersCreate(context.Background()).UserRequest(userRequest).Execute()
	return user, err
}

func (c *AuthentikClient) CreateUserGroups(req *ProvisionRequest) (*authentik.APIResponse, error) {
	fmt.Println(req, "implement me")
	return nil, nil
}
