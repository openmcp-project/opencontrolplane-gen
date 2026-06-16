package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/christophrj/opencontrolplane-gen/logs"
)

const ociReplace = "opencontrolplane-gen:replace"

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
	if Prefix(loc, ociReplace) {
		argAssignments := assignments(loc, ociReplace)
		if len(argAssignments) < 1 {
			logs.Debug(fmt.Sprintf("(%s) failed to parse (%s): invalid number of assignments", os.Getenv("GOFILE"), loc))
			return loc
		}
		r.arguments = []searchAndReplace{}
		for _, a := range argAssignments {
			replace, ok := os.LookupEnv(a.right)
			if !ok {
				logs.Debug(fmt.Sprintf("(%s) failed to lookup env (%s) of (%s)", os.Getenv("GOFILE"), a.right, loc))
			}
			r.arguments = append(r.arguments, searchAndReplace{search: a.left, replace: replace})
		}
		r.active = true
		logs.Debug(fmt.Sprintf("removed line: %s", loc))
		// remove the opencontrolplane-gen comment as part of the processing
		return ""
	}
	if r.active {
		original := loc
		for _, arg := range r.arguments {
			loc = strings.ReplaceAll(loc, arg.search, arg.replace)
		}
		logs.Debug(fmt.Sprintf("(%s) replaced (%s) with (%s)", os.Getenv("GOFILE"), original, loc))
		// replace is a one line command that instantly deactivates itself after processing a line of code
		r.active = false
	}
	return loc
}
