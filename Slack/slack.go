package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SlackPayload struct {
	Text interface{} `json:"text"`
}

func PostMessage(url string, data ...interface{}) {

	SlackData := SlackPayload{Text: data[0]}
	j, err := json.Marshal(SlackData)

	if err != nil {
		fmt.Print(err)
	}
	resp, err := http.Post(url, http.DetectContentType(j), bytes.NewBuffer(j))
	defer resp.Body.Close()
	if err != nil {
		log.Panic(err, resp.StatusCode)
	}
}
