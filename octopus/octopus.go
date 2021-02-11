package octopus

import (
	"fmt"
	"net/http"
	"os"
)

type Octopus struct {
	client  http.Client
	apiKey  string
	baseURL string
	Config  *Server
}

func NewOctopus(config *Server) Octopus {

	if config.Key == "" {
		fmt.Println("API Key is Required")
		os.Exit(1)
	}
	if config.URL == "" {
		fmt.Println("baseURL is Required")
		os.Exit(1)
	}
	o := Octopus{
		client:  http.Client{},
		apiKey:  config.Key,
		baseURL: config.URL,
		Config:  config,
	}
	return o
}

func (octopus *Octopus) doRequest(path string, method string) *http.Response {
	if method == "" {
		method = "GET"
	}
	req, _ := http.NewRequest("GET", octopus.baseURL+"/api/"+path, nil)
	req.Header.Add("X-Octopus-ApiKey", octopus.apiKey)
	resp, err := octopus.client.Do(req)
	if err != nil {
		fmt.Println("Error making request")
		os.Exit(2)
	}
	return resp
}
