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

	"github.com/EricFrancis12/url-checker-lambda/pkg"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	client    *http.Client
	hostnames []string = pkg.Dedupe(strings.Split(os.Getenv(pkg.EnvHostnames), "\\n"))
	targetUrl string   = os.Getenv(pkg.EnvTargetUrl)
)

func init() {
	if len(hostnames) < 2 {
		log.Fatal(fmt.Errorf("at least two hostnames are required"))
	}

	if _, err := url.Parse(targetUrl); err != nil {
		log.Fatal(fmt.Errorf("invalid target url: %s", targetUrl))
	}

	client = &http.Client{}
}

func handleRequest(ctx context.Context, _ json.RawMessage) (pkg.LambdaResp, error) {
	req, err := http.NewRequest(http.MethodGet, targetUrl, nil)
	if err != nil {
		err = fmt.Errorf("error creating http request: %v", err)
		return pkg.LambdaResp{}, err
	}

	authToken := os.Getenv(pkg.EnvLambdaAuthToken)
	if authToken != "" {
		req.Header.Add(pkg.HttpHeaderAuthorization, pkg.BearerHeader(authToken))
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("http GET error: %v", err)
		return pkg.LambdaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non 200 status code (%d)", resp.StatusCode)
		return pkg.LambdaResp{}, err
	}

	data, err := pkg.NewDataFromReader(resp.Body, hostnames)
	if err != nil {
		return pkg.LambdaResp{}, err
	}

	cdata, err := data.Compliment(hostnames)
	if err != nil {
		return pkg.LambdaResp{}, err
	}

	return cdata.Resp(), nil
}

func main() {
	lambda.Start(handleRequest)
}
