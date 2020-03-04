package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type NotFound struct{}

type jsonAPIError struct {
	Status string `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type jsonAPIErrors struct {
	Errors []jsonAPIError `json:"errors"`
}

func (h NotFound) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	e := jsonAPIErrors{
		Errors: []jsonAPIError{
			{
				Status: strconv.Itoa(http.StatusNotFound),
				Title:  "Not Found",
				Detail: fmt.Sprintf("The URI %s did not match any resources", req.URL.Path),
			},
		},
	}

	j, _ := json.Marshal(e)

	resp.WriteHeader(http.StatusNotFound)
	resp.Write(j)
}
