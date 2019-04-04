package logger

import (
	"fmt"
	publisher "logging-go/Publisher"
	slack "logging-go/Slack"
	"time"
)

type Conf map[string]string

type Fields map[string]interface{}

type LoggerConfig struct {
	SlackURL         string
	LoggerTimeFormat string
	RabbitmqURL      string
	RabbitmqQueue    string
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

func Debug(l LoggerConfig, destination string, a ...interface{}) {
	LogMe(l, destination, levelDebug, "", a...)
}

func Debugf(l LoggerConfig, destination string, format string, a ...interface{}) {
	LogMe(l, destination, levelDebug, format, a...)
}

func Info(l LoggerConfig, destination string, a ...interface{}) {
	LogMe(l, destination, levelInfo, "", a...)
}

func Infof(l LoggerConfig, destination string, format string, a ...interface{}) {
	LogMe(l, destination, levelInfo, format, a...)
}

func Println(l LoggerConfig, destination string, a ...interface{}) {
	LogMe(l, destination, levelInfo, "", a...)
}

func Printf(l LoggerConfig, destination string, format string, a ...interface{}) {
	LogMe(l, destination, levelInfo, format, a...)
}

func Warn(l LoggerConfig, destination string, a ...interface{}) {
	LogMe(l, destination, levelWarn, "", a...)
}

func Warnf(l LoggerConfig, destination string, format string, a ...interface{}) {
	LogMe(l, destination, levelWarn, format, a...)
}

func Error(l LoggerConfig, destination string, a ...interface{}) {
	LogMe(l, destination, levelErr, "", a...)
}

func Errorf(l LoggerConfig, destination string, format string, a ...interface{}) {
	LogMe(l, destination, levelErr, format, a...)
}

func Critical(l LoggerConfig, destination string, a ...interface{}) {
	LogMe(l, destination, levelCrit, "", a...)
}

func Criticalf(l LoggerConfig, destination string, format string, a ...interface{}) {
	LogMe(l, destination, levelCrit, format, a...)
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

func LogMe(l LoggerConfig, destination string, logLevel int, format string, data ...interface{}) {
	now := time.Now()
	levelDecoration := levelStrings[logLevel]
	var msg string
	/*if format != "" {
		msg = fmt.Sprintf(format, a...)
	} else {
		msg = fmt.Sprintln(a...)
	}*/
	if destination == "slack" {
		msg = fmt.Sprintf("%s: %s %s", now.Format(time.Stamp), levelDecoration, data[0])
		slack.PostMessage(l.SlackURL, msg)
	}
	if destination == "publish" {
		msg = fmt.Sprintf("%s: %s %s", now.Format(time.Stamp), levelDecoration, data[0])
		publisher.Publish(l.RabbitmqURL, l.RabbitmqQueue, msg)
	}

}

func EnableLogging(cnf Conf) LoggerConfig {

	return LoggerConfig{SlackURL: cnf["SlackURL"], LoggerTimeFormat: cnf["LoggerTimeFormat"],
		RabbitmqURL: cnf["RabbitmqURL"], RabbitmqQueue: cnf["RabbitmqQueue"]}

}
