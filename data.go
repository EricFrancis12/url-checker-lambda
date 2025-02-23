package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Data struct {
	Hostname string `json:"hostname"`
}

func NewDataFromReader(r io.Reader) (Data, error) {
	var data Data
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return Data{}, fmt.Errorf("error reading response body: %v", err)
	}

	if data.Hostname == "" {
		return Data{}, fmt.Errorf("hostname is missing")
	}

	if !sliceIncludes(hostnames, data.Hostname) {
		return Data{}, fmt.Errorf("unknown hostname: %s", data.Hostname)
	}

	return data, nil
}
