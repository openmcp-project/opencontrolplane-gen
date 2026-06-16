package logs

import (
	"fmt"
	"log"
	"os"
)

var debugLogs bool

// Init set global debug flag
func Init(debug bool) {
	debugLogs = debug
}

// Debug logs based on global debug flag setting and prefixes any log entry with filename + command identifier.
func Debug(commandIdentifier string, v ...any) {
	if debugLogs {
		log.Printf("%s %s: %s\n", os.Getenv("GOFILE"), commandIdentifier, fmt.Sprint(v...))
	}
}
