//go:generate sh -c "interfacer -for github.com/slack-go/slack.Client -as slackclient.SlackClient | grep -v _search > client_interface.go"
//go:generate mockgen -package mock_slackclient -destination=mock_slackclient/client.go .  SlackClient
package slackclient

import (
	"context"

	"github.com/slack-go/slack"
)

type contextKey string

var (
	clientContextKey contextKey = "slackclient.client"
)

// WithClient returns a new context.Context where calls to ClientFor
// will return the given client, rather than creating a new one.
//
// This can be used for testing, when we want any calls to build a
// client to return a mock, rather than a real implementation.
func WithClient(ctx context.Context, client SlackClient) context.Context {
	return context.WithValue(ctx, clientContextKey, client)
}

// ClientFor will return a SlackClient, either from the client stored in
// the context or by generating a new client for the given organisation.
func ClientFor(ctx context.Context, organisationID string) (SlackClient, error) {
	client, ok := ctx.Value(clientContextKey).(SlackClient)
	if ok {
		return client, nil
	}

	creds, err := getCredentials(ctx, organisationID)
	if err != nil {
		return nil, err
	}

	return slack.New(creds.SlackAccessToken), nil
}

type Credentials struct {
	SlackAccessToken string
}

func getCredentials(ctx context.Context, organisationID string) (*Credentials, error) {
	panic("this should never be called in our tests, as we should have received the mock instead")
}
