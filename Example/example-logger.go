package main

import (
	"logging-go"
	"os"
)

func main() {
	os.Setenv("slackURL", "https://hooks.slack.com/services/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")
	os.Setenv("loggerTimeFormat", "time.RFC3339")
	os.Setenv("loggerStream", "slack")
	logger.Critical("data", "suman")

}
