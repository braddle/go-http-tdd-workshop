package consumer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	h string
}

func NewClient(host string) Client {
	return Client{
		h: host,
	}
}

type HealthCheck struct {
	Status string   `json:"status"`
	Errors []string `json:"errors"`
}

func (c Client) HeathCheck() (HealthCheck, error) {
	url := fmt.Sprintf("%s/healthcheck", c.h)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")

	resp, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	h := HealthCheck{}

	json.Unmarshal(body, &h)

	return h, nil
}
