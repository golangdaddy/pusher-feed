package pusherfeed

import (
	"fmt"
	"flag"
	"testing"
)

var client *Client
var feed *Feed

func init() {

	instanceLocator := flag.String("instance", "", "is a string")
	keyId := flag.String("keyId", "", "is a string")
	secretKey := flag.String("keySecret", "", "is a string")
	feedId := flag.String("feedId", "myFeed", "is a string")

	flag.Parse()

	client = NewClient(
		*instanceLocator,
		*keyId,
		*secretKey,
	)
	feed = client.Feed(
		*feedId,
	)

	feed.isTestClient = true
}

func TestFeeds(t *testing.T) {

	msg := map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{
				"hello": "world",
			},
		},
	}

	list, err := client.Feeds(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(list)
}

func TestPublish(t *testing.T) {

	msg := map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{
				"hello": "world",
			},
		},
	}

	m, err := feed.Publish(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}
