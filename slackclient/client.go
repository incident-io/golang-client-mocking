//go:generate sh -c "interfacer -for github.com/slack-go/slack.Client -as slackclient.SlackClient | grep -v _search > client_interface.go"
//go:generate mockgen -package mock_slackclient -destination=mock_slackclient/client.go .  SlackClient
package slackclient

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/incident-io/golang-client-mocking/slackclient/mock_slackclient"
	"github.com/slack-go/slack"

	. "github.com/onsi/ginkgo"
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

// MockSlackClient is used in tests to generate a mock client and stash
// it into a context, ensuring all code will use the mock client instead
// of reaching out into the real world.
//
// Example is:
//
// Describe("subject", func() {
//   slackclient.MockSlackClient(&ctx, &sc, nil)
// })
func MockSlackClient(ctxPtr *context.Context, scPtr **mock_slackclient.MockSlackClient, ctrlPtr **gomock.Controller) {
	var ctrl *gomock.Controller
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		if ctrlPtr != nil {
			*ctrlPtr = ctrl
		}

		*scPtr = mock_slackclient.NewMockSlackClient(ctrl)
		*ctxPtr = WithClient(*ctxPtr, *scPtr)
	})
}
