package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type UrlLookup struct {
	Location string   `json:"location"`
	Created  string   `json:"created"`
	Updated  string   `json:"updated"`
	Sources  []string `json:"sources"`
	Ips      []string `json:"ips"`
	Domain   string   `json:"domain"`
}

func (self *Client) GetUrlLookup(location string) (UrlLookup, error) {
	fmt.Printf("looking upd %s\n", location)
	var data UrlLookup
	encoded := url.PathEscape(location)
	fmt.Printf("encoded as %s\n", encoded)
	apipath := fmt.Sprintf("url/%s", encoded)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}
