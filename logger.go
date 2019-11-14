package logger

import (
	"encoding/json"
	"fmt"
	publisher "github.com/Saurav-Suman/logging-go/Publisher"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Conf map[string]string
type QueueCategory struct {
	Api      string
	Debug    string
	Info     string
	Warn     string
	Error    string
	Critical string
	IO       string
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

// Log levels
const (
	levelCrit = iota + 1
	levelErr
	levelWarn
	levelInfo
	levelDebug
	levelIO
)

// Map
var levelStrings = map[int]string{
	levelDebug: "DEBUG",
	levelInfo:  "INFO",
	levelWarn:  "WARN",
	levelErr:   "ERROR",
	levelCrit:  "CRITICAL",
	levelIO:    "IO",
}

type SystemLoogerFields struct {
	Source    string
	Message   interface{}
	Timestamp string
}

func ConvertDataToString(w io.Writer, input ...interface{}) string {
	convertedString := ""
	for _, arg := range input {
		jsonString, _ := json.Marshal(arg)
		convertedString += string(jsonString)
	}
	return convertedString
}

func FeedDataForConversion(a ...interface{}) string {
	return ConvertDataToString(os.Stdout, a...)
}

func (l *SystemLoggerConfig) Debug(msg ...interface{}) {
	l.LogMe(levelDebug, l.QueueNames.Debug, msg)
}

func (l *SystemLoggerConfig) Info(msg ...interface{}) {
	l.LogMe(levelInfo, l.QueueNames.Info, msg)
}

func (l *SystemLoggerConfig) Warn(msg ...interface{}) {
	l.LogMe(levelWarn, l.QueueNames.Warn, msg)
}

func (l *SystemLoggerConfig) Error(msg ...interface{}) {
	l.LogMe(levelErr, l.QueueNames.Error, msg)
}

func (l *SystemLoggerConfig) Critical(msg ...interface{}) {
	l.LogMe(levelCrit, l.QueueNames.Critical, msg)
}

func (l *SystemLoggerConfig) IO(msg ...interface{}) {
	l.LogConsole(levelIO, l.QueueNames.IO, msg)
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

func (s *SystemLoggerConfig) LogMe(logLevel int, queueName string, msg ...interface{}) {

	if os.Getenv("app") == "" {
		log.Fatalf("%s", "Environment variable `app` missing in env file")
	}

	var msgData = FeedDataForConversion(msg)
	data := SystemLoogerFields{
		Source:    os.Getenv("app"),
		Message:   msgData,
		Timestamp: time.Now().Format(time.RFC3339), //currentTime.Format("2006.01.02 15:04:05")
	}

	//currentTime := time.Now()
	var queueToSend strings.Builder

	queueToSend.WriteString(s.QueuePrefix)
	queueToSend.WriteString(".")
	queueToSend.WriteString(queueName)
	if !s.Console {
		publisher.Publish(s.QueuePrefix, queueName, data)
	} else {
		publisher.Publish(s.QueuePrefix, queueName, msg)
	}
	return
}

func (s *SystemLoggerConfig) LogConsole(logLevel int, queueName string, msg ...interface{}) {

	if os.Getenv("app") == "" {
		log.Fatalf("%s", "Environment variable `app` missing in env file")
	}

	var msgData = FeedDataForConversion(msg)
	data := SystemLoogerFields{
		Source:    os.Getenv("app"),
		Message:   msgData,
		Timestamp: time.Now().Format(time.RFC3339), //currentTime.Format("2006.01.02 15:04:05")
	}

	var queueToSend strings.Builder

	queueToSend.WriteString(s.QueuePrefix)
	queueToSend.WriteString(".")
	queueToSend.WriteString(queueName)

	if s.Console {
		publisher.Publish(s.QueuePrefix, queueName, data)
	}
	return
}

func (s *SystemLoggerConfig) Api(data ApiLoggerFields) {
	var queueToSend strings.Builder
	//data.Timestamp = time.Now().Format(time.RFC3339) //currentTime.Format("2006.01.02 15:04:05")
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
		publisher.Publish(s.QueuePrefix, s.QueueNames.Api, string(publishData))

	}

}
