package main

import (
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

	//jsonString, _ := json.Marshal(dummydatatosend)

	log := logger.SystemLoggerConfig{
		Console:     true,
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

	log.InitLogging()

	log.Critical(logger.SystemLoogerFields{Message: "Divide By Zero"})

}
