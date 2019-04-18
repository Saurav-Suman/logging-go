package main

import (
	"encoding/json"
	logger "logging-go"
)

func main() {

	dummydatatosend := map[string]interface{}{
		"Name": "Ayopop",
		"Age":  3,
		"Parents": []interface{}{
			"Chiragh",
			"Jakob",
		},
	}

	jsonString, _ := json.Marshal(dummydatatosend)

	log := logger.SystemLoggerConfig{
		RabbitmqURL: "amqp://guest:guest@127.0.0.1:5672/",
		QueuePrefix: "ayopop",
		QueueNames: logger.QueueCategory{
			Api:      "api",
			Debug:    "debug",
			Info:     "info",
			Warn:     "warning",
			Error:    "error",
			Critical: "critical",
		},
	}

	log.Critical(logger.SystemLoogerFields{Source: "DoRecharge", Message: "Divide By Zero",
		Request:  string(jsonString),
		Response: string(jsonString),
	})

	log.Api(logger.ApiLoggerFields{Ip: "192.168.0.1",
		Url:        "getRechargeData",
		StatusCode: 200,
		Request:    string(jsonString),
		Method:     "POST",
		Headers:    "sdsdsdsd",
		Response:   string(jsonString),
	})

}
