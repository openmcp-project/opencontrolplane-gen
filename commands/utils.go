package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/christophrj/opencontrolplane-gen/logs"
)

// Prefix returns true if loc contains a comment starting with commandIdentifier
// e.g. '// opencontrolplane-gen:if a=b' returns true for commandIdentifier='opencontrolplane-gen:if'
func Prefix(loc, commandIdentifier string) bool {
	line := strings.TrimSpace(loc)
	if !strings.HasPrefix(line, "//") {
		return false
	}
	uncommentedCommand := uncomment(line)
	return strings.HasPrefix(uncommentedCommand, commandIdentifier)
}

// uncomment removes the go comment part
// e.g. '// opencontrolplane-gen:if a=b' returns 'opencontrolplane-gen:if a=b'
func uncomment(line string) string {
	return strings.TrimSpace(strings.TrimPrefix(line, "//"))
}

// trimCommand trims a command line to its arguments
// e.g. '// opencontrolplane-gen:if a=b' returns 'a=b'
func trimCommand(line, commandIdentifier string) string {
	return strings.TrimSpace(strings.TrimPrefix(uncomment(strings.TrimSpace(line)), commandIdentifier))
}

// arguments retrieves the raw arguments
// e.g. '// opencontrolplane-gen:if a=b c=d' returns '{"a=b", "c=d"}'
func arguments(loc, commandIdentifier string) []string {
	args := trimCommand(loc, commandIdentifier)
	return strings.Split(args, " ")
}

// assignments retrieves the arguments as structured assigments
// e.g. '// opencontrolplane-gen:if a=b c=d' returns '{{"left: a, right: b"}, {"left: c, right: d"}}'
func assignments(loc, commandIdentifier string) []assignment {
	args := arguments(loc, commandIdentifier)
	assignments := []assignment{}
	for _, a := range args {
		pair := strings.SplitN(a, "=", 2)
		if len(pair) != 2 {
			logs.Debug(fmt.Sprintf("(%s) failed to parse (%s): invalid argument assignment", os.Getenv("GOFILE"), loc))
			return nil
		}
		assignments = append(assignments, assignment{left: pair[0], right: pair[1]})
	}
	return assignments
}

func EvalBoolEnv(envVar string) bool {
	v := strings.ToLower(os.Getenv(envVar))
	return v == "1" || v == "true"
}

type assignment struct {
	left  string
	right string
}
