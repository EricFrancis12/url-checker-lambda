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
	hostnames []string = dedupe(strings.Split(os.Getenv(EnvHostnames), "\\n"))
	targetUrl string   = os.Getenv(EnvTargetUrl)
)

func init() {
	if len(hostnames) < 2 {
		log.Fatal(fmt.Errorf("at least two hostnames are required"))
	}

	if _, err := url.Parse(targetUrl); err != nil {
		log.Fatal(fmt.Errorf("invalid target url: %s", targetUrl))
	}
}

func handleRequest(ctx context.Context, _ json.RawMessage) (LambdaResp, error) {
	resp, err := http.Get(targetUrl)
	if err != nil {
		err = fmt.Errorf("http GET error: %v", err)
		return LambdaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non 200 status code (%d)", resp.StatusCode)
		return LambdaResp{}, err
	}

	data, err := NewDataFromReader(resp.Body)
	if err != nil {
		return LambdaResp{}, err
	}

	cdata, err := data.Compliment()
	if err != nil {
		return LambdaResp{}, err
	}

	return cdata.Resp(), nil
}

func main() {
	lambda.Start(handleRequest)
}
