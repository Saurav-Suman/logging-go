package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type WebHookMessage struct {
	Text string
}

func PostMessage(text string) {
	log.Println("sending message to web hook URL:", os.Getenv("LOGGER_SLACK_URL"))

	msg := WebHookMessage{Text: text}
	jsonContent, _ := json.Marshal(msg)
	resp, err := http.Post(os.Getenv("LOGGER_SLACK_URL"), http.DetectContentType(jsonContent), bytes.NewBuffer(jsonContent))

	log.Print("web hook status: ", resp.Status)
	if err != nil {
		log.Panic(err)
	}
}
