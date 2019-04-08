package main

import (
	logger "github.com/Saurav-Suman/logging-go"
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

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	loggerConf := logger.EnableLogging(logger.Conf{
		"LoggerTimeFormat": "time.RFC3339",
		"RabbitmqURL":      "amqp://guest:guest@127.0.0.1:5672/",
		"RabbitmqQueue":    "test",
	})

	//m := Message{"Alice", "Hello", 1294706395881547000}
	/*
		 logger.Fields{
			"animal": "walrus",
			"size":   10,
		}
	*/

	logger.Critical(loggerConf, "publish", "Divide by zero")

}
