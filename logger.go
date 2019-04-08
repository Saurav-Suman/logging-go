package logger

import (
	"fmt"
	"time"

	publisher "github.com/Saurav-Suman/logging-go/Publisher"
)

type Conf map[string]string

type ApiLoggerConfig struct {
	LoggerTimeFormat string
	RabbitmqURL      string
	Queue            string
}

type SystemLoggerConfig struct {
	LoggerTimeFormat string
	RabbitmqURL      string
	Queue            string
}

type ApiLoogerFields struct {
	Ip         string
	Url        string
	StatusCode string
	Request    interface{}
	Method     string
	Headers    interface{}
	Response   interface{}
	Timestamp  string
}

type SystemLoogerFields struct {
	ErrorMsg  string
	Module    string
	Timestamp string
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

func (l *SystemLoggerConfig) Debug(msg string, a ...interface{}) {
	LogMe(levelDebug, "", a...)
}

func (l *SystemLoggerConfig) Debugf(msg string, a ...interface{}) {
	LogMe(levelDebug, format, a...)
}

func (l *SystemLoggerConfig) Info(msg string, a ...interface{}) {
	LogMe(levelInfo, "", a...)
}

func (l *SystemLoggerConfig) Infof(msg string, a ...interface{}) {
	LogMe(levelInfo, format, a...)
}

func (l *SystemLoggerConfig) Warn(msg string, a ...interface{}) {
	LogMe(levelWarn, "", a...)
}

func (l *SystemLoggerConfig) Warnf(msg string, a ...interface{}) {
	LogMe(levelWarn, format, a...)
}

func (l *SystemLoggerConfig) Error(msg string, a ...interface{}) {
	LogMe(levelErr, "", a...)
}

func (l *SystemLoggerConfig) Errorf(msg string, a ...interface{}) {
	LogMe(levelErr, format, a...)
}

func (l *SystemLoggerConfig) Critical(msg string, a ...interface{}) {
	LogMe(levelCrit, "", a...)
}

func (l *SystemLoggerConfig) Criticalf(msg string, a ...interface{}) {
	LogMe(levelCrit, format, a...)
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

func LogMe(logLevel int, format string, data ...interface{}) {
	now := time.Now()
	levelDecoration := levelStrings[logLevel]
	var msg string
	/*if format != "" {
		msg = fmt.Sprintf(format, a...)
	} else {
		msg = fmt.Sprintln(a...)
	}*/
	msg = fmt.Sprintf("%s: %s %s", now.Format(time.Stamp), levelDecoration, data[0])
	publisher.Publish(l.RabbitmqURL, l.RabbitmqQueue, msg)

}

func EnableLogging(cnf Conf) LoggerConfig {

	return LoggerConfig{LoggerTimeFormat: cnf["LoggerTimeFormat"],
		RabbitmqURL: cnf["RabbitmqURL"], RabbitmqQueue: cnf["RabbitmqQueue"]}

}
