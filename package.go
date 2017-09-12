package pusherfeed

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	//
	"github.com/dghubble/sling"
)

const (
	CONST_PUBLISH_ENDPOINT = "https://us1.pusherplatform.io/services/feeds/v1/%s/feeds/%s/items"
)

func buildUrl(instanceId, feedId string) string {
	return fmt.Sprintf(CONST_PUBLISH_ENDPOINT, instanceId, feedId)
}

type Payload struct {
	Items []interface{} `json:"items"`
}

type Client struct {
	httpClient *http.Client
	instance string
	keyId string
	keySecret string
	feedId string
}

func NewClient(instance, keyId, keySecret, feedId string) *Client {

	return &Client{
		httpClient: &http.Client{},
		instance: instance,
		keyId: keyId,
		keySecret: keySecret,
		feedId: feedId,
	}
}

func (client *Client) post(url string, msg interface{}) (map[string]interface{}, error) {

	request, err := sling.New().Post(url).BodyJSON(msg).Request()

	request.Header.Add("Authorization", "Bearer " + client.NewToken())

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj := make(map[string]interface{})

	if err := json.Unmarshal(b, &obj); err != nil {
		fmt.Println(string(b))
		return nil, err
	}

	return obj, nil
}
