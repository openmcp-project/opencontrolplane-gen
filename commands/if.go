package commands

import (
	"fmt"
	"os"

	"github.com/christophrj/opencontrolplane-gen/logs"
)

const (
	ocpIf = "opencontrolplane-gen:if"
	ocpFi = "opencontrolplane-gen:fi"
)

var _ Command = &ifCommand{}

// Activates on `//opencontrolplane-gen:if`.
// Expects an env variable as parameter that holds a bool value.
// e.g. `//opencontrolplane-gen:if ENV_FEATURE
// where ENV_FEATURE = 'false' will result in any following line to be removed
// until `//opencontrolplane-gen-fi` deactivates the command again.
type ifCommand struct {
	active    bool
	condition bool
}

// NewIfCommand returns an ifCommand
func NewIfCommand() Command {
	return &ifCommand{}
}

// Execute implements [Command].
func (r *ifCommand) Execute(loc string) string {
	if Prefix(loc, ocpIf) {
		argAssignments := assignments(loc, ocpIf)
		if len(argAssignments) != 1 {
			logs.Debug(fmt.Sprintf("(%s) failed to parse (%s): invalid number of assignments", os.Getenv("GOFILE"), loc))
			return loc
		}
		r.active = true
		r.condition = (os.Getenv(argAssignments[0].left) == argAssignments[0].right)
		logs.Debug(fmt.Sprintf("ifCommand condition = %v", r.condition))
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the opencontrolplane-gen comment as part of the processing
		return ""
	}
	if Prefix(loc, ocpFi) {
		r.active = false
		r.condition = false
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the opencontrolplane-gen comment as part of the processing
		return ""
	}
	if r.active && !r.condition {
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		return ""
	}
	return loc
}
