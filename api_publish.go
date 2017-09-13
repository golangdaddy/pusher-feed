package pusherfeed

import (
	"fmt"
)

func (client *Client) Publish(msg interface{}) (map[string]interface{}, error) {

	url := fmt.Sprintf(CONST_PUBLISH_ENDPOINT, client.instance, client.feedId)

	return client.post(url, msg)
}
