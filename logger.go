package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	publisher "github.com/Saurav-Suman/logging-go/Publisher"
)

//github.com/Saurav-Suman/
type Conf map[string]string

type QueueCategory struct {
	Api      string
	Debug    string
	Info     string
	Warn     string
	Error    string
	Critical string
}

type ApiLoggerFields struct {
	Ip         string
	Url        string
	StatusCode int
	Request    interface{}
	Method     string
	Headers    interface{}
	Response   interface{}
	Timestamp  string
}

type SystemLoggerConfig struct {
	Console     bool
	RabbitmqURL string
	QueuePrefix string
	QueueNames  QueueCategory
}

type SystemLoogerFields struct {
	Message string
}

// Log levels
const (
	levelCrit = iota + 1
	levelErr
	levelWarn
	levelInfo
	levelDebug
)

// Map
var levelStrings = map[int]string{
	levelDebug: "DEBUG",
	levelInfo:  "INFO",
	levelWarn:  "WARN",
	levelErr:   "ERROR",
	levelCrit:  "CRITICAL",
}

func (l *SystemLoggerConfig) Debug(msg SystemLoogerFields) {
	l.LogMe(levelDebug, l.QueueNames.Debug, msg)
}

func (l *SystemLoggerConfig) Info(msg SystemLoogerFields) {
	l.LogMe(levelInfo, l.QueueNames.Info, msg)
}

func (l *SystemLoggerConfig) Warn(msg SystemLoogerFields) {
	l.LogMe(levelWarn, l.QueueNames.Warn, msg)
}

func (l *SystemLoggerConfig) Error(msg SystemLoogerFields) {
	l.LogMe(levelErr, l.QueueNames.Error, msg)
}

func (l *SystemLoggerConfig) Critical(msg SystemLoogerFields) {
	l.LogMe(levelCrit, l.QueueNames.Critical, msg)
}

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

func (s *SystemLoggerConfig) InitLogging() {
	publisher.InitRMQ(s.RabbitmqURL)
}

func (s *SystemLoggerConfig) LogMe(logLevel int, queueName string, data SystemLoogerFields) {
	//currentTime := time.Now()
	data.Timestamp = time.Now().Format(time.RFC3339) //currentTime.Format("2006.01.02 15:04:05")
	var queueToSend strings.Builder

	queueToSend.WriteString(s.QueuePrefix)
	queueToSend.WriteString(".")
	queueToSend.WriteString(queueName)
	if !s.Console {
		publisher.Publish(s.QueuePrefix, queueName, data)
	} else {
		publishData, err := json.Marshal(data)

		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(publishData))
	}

}

func (s *SystemLoggerConfig) Api(data ApiLoggerFields) {
	var queueToSend strings.Builder
	data.Timestamp = time.Now().Format(time.RFC3339) //currentTime.Format("2006.01.02 15:04:05")
	queueToSend.WriteString(s.QueuePrefix)
	queueToSend.WriteString(".")
	queueToSend.WriteString(s.QueueNames.Api)
	if !s.Console {
		publisher.Publish(s.QueuePrefix, s.QueueNames.Api, data)
	} else {
		publishData, err := json.Marshal(data)

		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(string(publishData))
	}

}
