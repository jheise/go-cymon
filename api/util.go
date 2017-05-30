package api

import (
	// standard
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
}

func getData(apikey string, endpoint string) ([]byte, error) {
	urlDst := fmt.Sprintf("https://cymon.io/api/nexus/v1/%s", endpoint)
	client := new(http.Client)
	req, err := http.NewRequest("GET", urlDst, nil)
	if err != nil {
		return []byte{}, err
	}

	apitoken := fmt.Sprintf("Token %s", apikey)
	req.Header.Add("Authorization", apitoken)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
