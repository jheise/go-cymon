package api

import (
	// "encoding/json"
	"fmt"
)

type DgaLookup struct {
	Name    string   `json:"name"`
	Created string   `json:"created"`
	Updated string   `json:"updated"`
	Sources []string `json:"sources"`
	Ips     []string `json:"ips"`
	Urls    []string `json:"urls"`
}

func (self *Client) GetDgaLookup(dga string) (DgaLookup, error) {
	var data DgaLookup
	apipath := fmt.Sprintf("dga/%s", dga)
	body, err := getData(self.key, apipath)
	if err != nil {
		return data, err
	}

	fmt.Println(string(body))
	// json.Unmarshal(body, &data)
	return data, nil
}
