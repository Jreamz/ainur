package internal

import (
	"context"
	"log"

	authentik "goauthentik.io/api/v3"
)

type AuthentikClient struct {
	client *authentik.APIClient
}

func NewAuthentikClient(endpoint, token string) *AuthentikClient {
	config := authentik.NewConfiguration()
	config.Host = endpoint
	config.Scheme = "http"
	config.AddDefaultHeader("Authorization", "Bearer "+token)

	return &AuthentikClient{
		client: authentik.NewAPIClient(config),
	}
}

func (c *AuthentikClient) CreateUserRequest(pr *ProvisionRequest) (*authentik.User, error) {
	userRequest := *authentik.NewUserRequest(pr.Email, pr.FirstName+" "+pr.LastName)
	user, _, err := c.client.CoreApi.CoreUsersCreate(context.Background()).UserRequest(userRequest).Execute()
	log.Printf("authentik error: %v", err)
	return user, err
}

func (c *AuthentikClient) SearchUsersList(query string) (any, error) {
	users, _, err := c.client.CoreApi.CoreUsersList(
		context.Background(),
	).Search(query).IncludeGroups(true).IncludeRoles(true).Execute()

	return users.Results, err
}
