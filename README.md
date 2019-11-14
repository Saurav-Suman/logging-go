# What is `logging-go`?

`logging-go` supports log levels And you can publish log to rabbitMQ for now.

# _Another_ logging library, Why?

> Logging libraries are like opinions, everyone needs one depends upon the need.


## And how is `logging-go` different?

- RabbitMQ Integration



# Usage/examples:



Publish to RabbitMQ using a logging-go instantiated like so:

```golang

import (
	logger "github.com/Saurav-Suman/logging-go"
)

package main

import (
	logger "logging-go"
)

/*
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
*/

func main() {

	log := logger.SystemLoggerConfig{
		LoggerTimeFormat: "time.RFC3339",
		RabbitmqURL:      "amqp://guest:guest@127.0.0.1:5672/",
		QueuePrefix:      "ayopop",
		QueueNames: logger.QueueCategory{
			Api:      "api",
			Debug:    "debug",
			Info:     "info",
			Warn:     "warning",
			Error:    "error",
			Critical: "critical",
		},
	}

	

	log.Critical("Divide by zero")
	log.Info("Runming")
	log.Infof("error at line %d and error is ", 45, "fatal error ")
	log.Criticalf("Divide by zero %s", "0")

	log.Api(logger.ApiLoggerFields{Ip: "192.168.0.1",
		Url:        "getRechargeData",
		StatusCode: 200,
		Request:    "Name=saurav",
		Method:     "POST",
		Headers:    "",
		Response:   "Result=True",
		Timestamp:  "",
	})

}
```



# Contributing

1. Create an issue, describe the bugfix/feature you wish to implement.
2. Fork the repository
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request


# Note:

1. For using multiple parameters using custom inputs for logging; PLEASE REFER release version = 1.1
2. For using predefined (SystemLoogerData) parameters for logging; PLEASE REFER release version = 1.0
