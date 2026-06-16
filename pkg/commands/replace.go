package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/openmcp-project/opencontrolplane-gen/pkg/logs"
)

const ocpReplace = "opencontrolplane-gen:replace"

var _ Command = &replaceCommand{}

// Activates on `//opencontrolplane-gen:replace`.
// Expects a variable number of search and replace pairs separated by '='.
// E.g. `//opencontrolplane-gen:replace search-a=ENV_A search-b=ENV_B`
// where ENV_* will be replaced by the result of the env variable lookup.
type replaceCommand struct {
	active    bool
	arguments []searchAndReplace
}

type searchAndReplace struct {
	search  string
	replace string
}

// NewReplaceCommand returns a replaceCommand
func NewReplaceCommand() Command {
	return &replaceCommand{}
}

// Execute implements [Command].
func (r *replaceCommand) Execute(loc string) string {
	// enable command
	if Prefix(loc, ocpReplace) {
		argAssignments := assignments(loc, ocpReplace)
		if len(argAssignments) < 1 {
			logs.Debug(ocpReplace, fmt.Sprintf("failed to parse (%s): invalid number of assignments", loc))
			return loc
		}
		r.arguments = []searchAndReplace{}
		for _, a := range argAssignments {
			replace, ok := os.LookupEnv(a.right)
			if !ok {
				logs.Debug(ocpReplace, fmt.Sprintf("failed to lookup env (%s) of (%s)", a.right, loc))
				r.arguments = nil
				return loc
			}
			r.arguments = append(r.arguments, searchAndReplace{search: a.left, replace: replace})
		}
		r.active = true
		logs.Debug(ocpReplace, fmt.Sprintf("removed line: %s", loc))
		// remove the opencontrolplane-gen comment as part of the processing
		return ""
	}
	// replace and disable command
	if r.active {
		original := loc
		for _, arg := range r.arguments {
			loc = strings.ReplaceAll(loc, arg.search, arg.replace)
		}
		logs.Debug(ocpReplace, fmt.Sprintf("replaced \"%s\" with \"%s\"", original, loc))
		// replace is a one line command that instantly deactivates itself after processing a line of code
		r.active = false
		r.arguments = nil
	}
	return loc
}
