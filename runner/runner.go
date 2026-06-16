package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/christophrj/opencontrolplane-gen/commands"
)

const ocpgenIdentifier = "go:generate opencontrolplane-gen"

// Runner takes a set of Commands
type Runner struct {
	Commands []commands.Command
}

// Run executes the runner commands on the passed in file
func (r *Runner) Run(fpath string) (result bytes.Buffer) {
	// iterate over each file
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		original := fileScanner.Text()
		loc := original
		for _, c := range r.Commands {
			loc = c.Execute(loc)
		}
		if commands.Prefix(loc, ocpgenIdentifier) {
			// remove go:generate opencontrolplane-gen lines
			loc = ""
		}
		// add newline unless line should be removed
		if original == loc || loc != "" {
			if _, err := fmt.Fprintln(&result, loc); err != nil {
				log.Fatalf("failed to write buffer: %v", err)
			}
		}
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("failed to scan source file: %v", err)
	}
	return result
}
