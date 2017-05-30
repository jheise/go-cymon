package api

import (
	"io/ioutil"
	"strings"
)

type Client struct {
	key string
}

func NewClient(key string) *Client {
	client := new(Client)
	client.key = key

	return client
}

func NewClientFromFile(keypath string) (*Client, error) {
	client := new(Client)
	keydata, err := ioutil.ReadFile(keypath)
	if err != nil {
		return client, err
	}

	client.key = strings.TrimSpace(string(keydata))
	return client, nil
}
