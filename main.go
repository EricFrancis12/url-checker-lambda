package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	hostnames []string = strings.Split(os.Getenv(EnvHostnames), "\n")
	targetUrl string   = os.Getenv(EnvTargetUrl)
)

func init() {
	if len(hostnames) == 0 {
		log.Fatal(fmt.Errorf("at least one hostname is required"))
	}

	if _, err := url.Parse(targetUrl); err != nil {
		log.Fatal(fmt.Errorf("invalid target url: %s", targetUrl))
	}
}

func handleRequest(ctx context.Context, event json.RawMessage) json.RawMessage {
	resp, err := http.Get(targetUrl)
	if err != nil {
		err = fmt.Errorf("http GET error: %v", err)
		return NewErrResp(err).ToBytes()
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("received non 200 status code (%d)", resp.StatusCode)
		return NewErrResp(err).ToBytes()
	}

	data, err := NewDataFromReader(resp.Body)
	if err != nil {
		return NewErrResp(err).ToBytes()
	}

	return NewLambdaResp(data, nil).ToBytes()
}

func main() {
	lambda.Start(handleRequest)
}
