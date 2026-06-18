package logs

import (
	"fmt"
	"log"
	"os"
)

var debugLogs bool

// Init sets global debug flag and logging to stdout instead of stderr
func Init(debug bool) {
	debugLogs = debug
	log.SetOutput(os.Stdout)
}

// Debug logs based on global debug flag setting and prefixes any log entry with filename + command identifier.
func Debug(commandIdentifier string, v ...any) {
	if debugLogs {
		log.Printf("%s %s: %s\n", os.Getenv("GOFILE"), commandIdentifier, fmt.Sprint(v...))
	}
}
