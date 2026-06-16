package commands

import (
	"fmt"
	"os"

	"github.com/christophrj/opencontrolplane-gen/pkg/logs"
)

const (
	ocpIf = "opencontrolplane-gen:if"
	ocpFi = "opencontrolplane-gen:fi"
)

var _ Command = &ifCommand{}

// Activates on `//opencontrolplane-gen:if`.
// Expects an assignment where an env variable value is compared with a string literal
// e.g. `//opencontrolplane-gen:if ENV_VAR=somevalue
// if the ENV_VAR lookup does not match the string literal, then any following line of code gets removed
// until `//opencontrolplane-gen-fi` deactivates the command again.
type ifCommand struct {
	active      bool
	includeLine bool
}

// NewIfCommand returns an ifCommand
func NewIfCommand() Command {
	return &ifCommand{}
}

// Execute implements [Command].
func (r *ifCommand) Execute(loc string) string {
	// enable command
	if Prefix(loc, ocpIf) {
		argAssignments := assignments(loc, ocpIf)
		// exactly one assignment is expected
		if len(argAssignments) != 1 {
			logs.Debug(fmt.Sprintf("(%s) failed to parse (%s): invalid number of assignments", os.Getenv("GOFILE"), loc))
			r.includeLine = false
			return loc
		}
		value, ok := os.LookupEnv(argAssignments[0].left)
		if !ok {
			logs.Debug(fmt.Sprintf("(%s) failed to lookup env (%s) of (%s)", os.Getenv("GOFILE"), argAssignments[0].left, loc))
			r.includeLine = false
			return loc
		}
		r.active = true
		r.includeLine = (value == argAssignments[0].right)
		logs.Debug(fmt.Sprintf("ifCommand includeLine = %v", r.includeLine))
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the opencontrolplane-gen comment as part of the processing
		return ""
	}
	// disable command
	if Prefix(loc, ocpFi) && r.active {
		r.active = false
		r.includeLine = false
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the opencontrolplane-gen comment as part of the processing
		return ""
	}
	if r.active && !r.includeLine {
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		return ""
	}
	return loc
}
