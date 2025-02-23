package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type EventMsg struct {
	Url string `json:"url"`
}

func EventMsgFromBytes(b []byte) (EventMsg, error) {
	var em EventMsg
	if err := json.Unmarshal(b, &em); err != nil {
		return EventMsg{}, err
	}

	_, err := url.Parse(em.Url)
	if err != nil {
		return EventMsg{}, fmt.Errorf("invalid url: %s", em.Url)
	}

	return em, nil
}
