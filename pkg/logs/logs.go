package logs

import (
	"log"
)

var debugLogs bool

// Init set global debug flag
func Init(debug bool) {
	debugLogs = debug
}

// Debug logs based on global debug flag setting
func Debug(v ...any) {
	if debugLogs {
		log.Println(v...)
	}
}
