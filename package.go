package pusherfeed

import (
	"golang.org/x/net/context"
	//
	"github.com/golangdaddy/tarantula/httpclient"
)

const (
	CONST_ENDPOINT_INSTANCE_FEEDS = "https://us1.pusherplatform.io/services/feeds/v1/%s/feeds"
	CONST_ENDPOINT_PUBLISH = "https://us1.pusherplatform.io/services/feeds/v1/%s/feeds/%s/items"
)

type Payload struct {
	Items []interface{} `json:"items"`
}

type Client struct {
	*httpclient.Client
	isTestClient bool
	instanceLocator string
	keyId string
	secretKey string
}

type Feed struct {
	*Client
	feedId string
}

func NewClient(instanceLocator, keyId, secretKey string) *Client {

	return &Client{
		Client: httpclient.NewClient(),
		instanceLocator: instanceLocator,
		keyId: keyId,
		secretKey: secretKey,
	}
}

func NewUrlfetchClient(ctx context.Context, instanceLocator, keyId, secretKey string) *Client {

	return &Client{
		Client: httpclient.NewUrlfetchClient(ctx),
		instanceLocator: instanceLocator,
		keyId: keyId,
		secretKey: secretKey,
	}
}

func (client *Client) Feed(id string) *Feed {
	return &Feed{
		client,
		id,
	}
}
