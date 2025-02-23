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

	if !includes(hostnames, data.Hostname) {
		return Data{}, fmt.Errorf("unknown hostname: %s", data.Hostname)
	}

	return data, nil
}

func (d Data) Compliment() (Data, error) {
	var filteredHostnames = filter(hostnames, func(hn string) bool {
		return hn != d.Hostname
	})

	if len(filteredHostnames) == 0 {
		return Data{}, fmt.Errorf("no hostname compliments available")
	}

	return Data{
		Hostname: mustGetRand(filteredHostnames),
	}, nil
}

func (d Data) Resp() LambdaResp {
	return NewLambdaResp(d.Hostname)
}
