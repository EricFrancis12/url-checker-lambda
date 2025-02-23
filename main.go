package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event json.RawMessage) json.RawMessage {
	msg, err := EventMsgFromBytes(event)
	if err != nil {
		return NewErrResp(err).ToBytes()
	}

	resp, err := http.Get(msg.Url)
	if err != nil {
		err = fmt.Errorf("http GET error: %v", err)
		return NewErrResp(err).ToBytes()
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("received non 200 status code (%d)", resp.StatusCode)
		return NewErrResp(err).ToBytes()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("error reading response body: %v", err)
		return NewErrResp(err).ToBytes()
	}

	return NewLambdaResp(body, nil).ToBytes()
}

func main() {
	lambda.Start(handleRequest)
}
