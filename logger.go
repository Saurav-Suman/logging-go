package logger

import (
	"fmt"
	"time"
)

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

func Debug(a ...interface{}) {
	LogMe(levelDebug, "", a...)
}

func Debugf(format string, a ...interface{}) {
	LogMe(levelDebug, format, a...)
}

func Info(a ...interface{}) {
	LogMe(levelInfo, "", a...)
}

func Infof(format string, a ...interface{}) {
	LogMe(levelInfo, format, a...)
}

func Println(a ...interface{}) {
	LogMe(levelInfo, "", a...)
}

func Printf(format string, a ...interface{}) {
	LogMe(levelInfo, format, a...)
}

func Warn(a ...interface{}) {
	LogMe(levelWarn, "", a...)
}

func Warnf(format string, a ...interface{}) {
	LogMe(levelWarn, format, a...)
}

func Error(a ...interface{}) {
	LogMe(levelErr, "", a...)
}

func Errorf(format string, a ...interface{}) {
	LogMe(levelErr, format, a...)
}

func Critical(a ...interface{}) {
	LogMe(levelCrit, "", a...)
}

func Criticalf(format string, a ...interface{}) {
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

func LogMe(logLevel int, format string, a ...interface{}) {
	now := time.Now()
	levelDecoration := levelStrings[logLevel]
	var msg string
	if format != "" {
		msg = fmt.Sprintf(format, a...)
	} else {
		msg = fmt.Sprintln(a...)
	}
	fmt.Printf("%s: %s %s", now.Format(time.RFC3339), levelDecoration, msg)

}
