package slackclient_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/lawrencejones/scratch/golang-client-mocking/slackclient"
	mock_slackclient "github.com/lawrencejones/scratch/golang-client-mocking/slackclient/mock_slackclient"
	"github.com/slack-go/slack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClientFor", func() {
	var (
		ctx  context.Context
		sc   *mock_slackclient.MockSlackClient
		ctrl *gomock.Controller
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		sc = mock_slackclient.NewMockSlackClient(ctrl)
		ctx = slackclient.WithClient(context.Background(), sc)
	})

	Describe("mocking a channel call", func() {
		BeforeEach(func() {
			sc.EXPECT().GetConversationInfoContext(gomock.Any(), "CH123", false).
				Return(&slack.Channel{
					GroupConversation: slack.GroupConversation{
						Name: "my-channel",
						Conversation: slack.Conversation{
							NameNormalized: "my-channel",
							ID:             "CH123",
						},
					},
				}, nil).Times(1)
		})

		Specify("returns a client that responds with the mock", func() {
			client, err := slackclient.ClientFor(ctx, "OR123")
			Expect(err).NotTo(HaveOccurred(), "Slack client should have built with no error")

			channel, err := client.GetConversationInfoContext(ctx, "CH123", false)
			Expect(err).NotTo(HaveOccurred())

			Expect(channel.NameNormalized).To(Equal("my-channel"))
		})
	})
})

// MockSlackClient is used to generate a mock Slack client and stash it
// into the test context, ensuring engine code will use the mock client
// instead of reaching out into the real world.
func MockSlackClient(ctxPtr *context.Context, scPtr **mock_slackclient.MockSlackClient, ctrlPtr **gomock.Controller) {
	var ctrl *gomock.Controller
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		if ctrlPtr != nil {
			*ctrlPtr = ctrl
		}

		*scPtr = mock_slackclient.NewMockSlackClient(ctrl)
		*ctxPtr = slackclient.WithClient(*ctxPtr, *scPtr)
	})
}
