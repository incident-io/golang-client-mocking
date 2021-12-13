//go:generate sh -c "interfacer -for github.com/slack-go/slack.Client -as slackclient.SlackClient | grep -v _search > client_interface.go"
//go:generate mockgen -package mock_slackclient -destination=mock_slackclient/client.go .  SlackClient
package slackclient

import (
	"context"

	"github.com/slack-go/slack"
)

func ClientFor(ctx context.Context, organisationID string) (*slack.Client, error) {
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
	panic("unimplemented")
}
