package pusherfeed

import (
	"fmt"
)

func (client *Client) Feeds(msg interface{}) ([]interface{}, error) {

	url := fmt.Sprintf(CONST_ENDPOINT_INSTANCE_FEEDS, client.instanceLocator)

	h := map[string]string{
		"Authorization": "Bearer "+client.NewToken(),
	}

	m := []interface{}{}
	_, err := client.Get(url, &m, h)

	return m, err
}

func (feed *Feed) Publish(msgs ...interface{}) (map[string]interface{}, error) {

	url := fmt.Sprintf(CONST_ENDPOINT_PUBLISH, feed.instanceLocator, feed.feedId)

	h := map[string]string{
		"Authorization": "Bearer "+feed.NewToken(),
	}

	items := map[string]interface{}{
		"items": append(
			[]interface{}{},
			msgs...,
		),
	}

	m := map[string]interface{}{}
	_, err := feed.Post(url, items, &m, h)

	return m, err
}
