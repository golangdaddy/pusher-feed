package pusherfeed

import (
	"fmt"
	"flag"
	"testing"
)

var client *Client

func init() {

	instance := flag.String("instance", "v1:us1:c2249ba7-ade0-0000-8864-000000dad3f6", "is a string")
	keyId := flag.String("keyId", "87f87a37-5aba-412e-b2ee-00000c000000", "is a string")
	keySecret := flag.String("keySecret", "6D9/ZfCKLMkgG8d1fopQIrG+OxEUhzEgS/1zvhtkUL8=", "is a string")
	feedId := flag.String("feedId", "myFeed", "is a string")

	flag.Parse()

	client = NewClient(
		*instance,
		*keyId,
		*keySecret,
		*feedId,
	)

	client.isTestClient = true
}

func TestPublish(t *testing.T) {

	msg := map[string]interface{}{
		"hello": "world",
	}

	_, err := client.Publish(msg)
	if err != nil {
		panic(err)
	}
}

func TestNewToken(t *testing.T) {

	jwt := client.NewToken()

	fmt.Println(jwt)

	if len(jwt) == 0 {
		t.Fail()
	}
}
