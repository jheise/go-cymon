package api

import (
	"encoding/json"
	"fmt"
)

type DomainLookup struct {
	Name    string   `json:"name"`
	Created string   `json:"created"`
	Updated string   `json:"updated"`
	Sources []string `json:"sources"`
	Ips     []string `json:"ips"`
	Urls    []string `json:"urls"`
}

func (self *Client) GetDomainLookup(domain string) (DomainLookup, error) {
	var data DomainLookup
	apipath := fmt.Sprintf("domain/%s", domain)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}
