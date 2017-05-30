package api

import (
	"encoding/json"
	"fmt"
)

type BlacklistEntry struct {
	Addr string `json:"addr"`
	Url  string `json:"url"`
}

type Blacklist struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Results  []BlacklistEntry `json:"results"`
}

func (self *Client) GetIPAddrBlacklist(tag string, days int) (Blacklist, error) {
	var data Blacklist
	apipath := fmt.Sprintf("blacklist/ip/%s/", tag)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetDomainBlacklist(tag string, days int) (Blacklist, error) {
	var data Blacklist
	apipath := fmt.Sprintf("blacklist/domain/%s/", tag)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}
