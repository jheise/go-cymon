package api

import (
	"encoding/json"
	"fmt"
)

type IPAddrLookup struct {
	Addr    string   `json:"addr"`
	Created string   `json:"created"`
	Updated string   `json:"updated"`
	Sources []string `json:"sources"`
	Events  string   `json:"events"`
	Domains string   `json:"domains"`
	Urls    string   `json:"urls"`
}

type IPAddrEvent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DetailsUrl  string `json:"details_url"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Tag         string `json:"tag"`
}

type IPAddrEvents struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []IPAddrEvent `json:"results"`
}

type IPAddrDomain struct {
	Name    string `json:"name"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

type IPAddrDomains struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []IPAddrDomain `json:"results"`
}

type IPAddrUrl struct {
	Location string `json:"location"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type IPAddrUrls struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []IPAddrUrl `json:"results"`
}

func (self *Client) GetIPAddrLookup(ipaddr string) (IPAddrLookup, error) {
	var data IPAddrLookup
	apipath := fmt.Sprintf("ip/%s/", ipaddr)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPAddrEvents(ipaddr string) (IPAddrEvents, error) {
	var data IPAddrEvents
	apipath := fmt.Sprintf("ip/%s/events/", ipaddr)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPAddrDomains(ipaddr string) (IPAddrDomains, error) {
	var data IPAddrDomains
	apipath := fmt.Sprintf("ip/%s/domains/", ipaddr)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}

func (self *Client) GetIPAddrUrls(ipaddr string) (IPAddrUrls, error) {
	var data IPAddrUrls
	apipath := fmt.Sprintf("ip/%s/urls/", ipaddr)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	json.Unmarshal(body, &data)
	return data, nil
}
