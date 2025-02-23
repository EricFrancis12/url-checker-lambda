package pkg

import (
	"encoding/json"
	"fmt"
	"io"
)

type Data struct {
	Hostname string `json:"hostname"`
}

func NewData(hostname string) Data {
	return Data{
		Hostname: hostname,
	}
}

func NewDataFromReader(r io.Reader, hostnames []string) (Data, error) {
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

func (d Data) Compliment(hostnames []string) (Data, error) {
	var filteredHostnames = filter(hostnames, func(hn string) bool {
		return hn != d.Hostname
	})

	if len(filteredHostnames) == 0 {
		return Data{}, fmt.Errorf("no hostname compliments available")
	}

	return NewData(mustGetRand(filteredHostnames)), nil
}

func (d Data) Resp() LambdaResp {
	return NewLambdaResp(d.Hostname)
}

func (d Data) Json() []byte {
	return []byte(fmt.Sprintf(`{"hostname":"%s"}`, d.Hostname))
}
