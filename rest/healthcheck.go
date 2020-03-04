package rest

import (
	"encoding/json"
	"net/http"
)

type HealthCheck struct {
}

type Content struct {
	S string   `json:"status"`
	E []string `json:"errors"`
}

func (h HealthCheck) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	c := Content{S: "OK", E: []string{}}

	b, _ := json.Marshal(c)

	resp.Write(b)
}
