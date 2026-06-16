package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/openmcp-project/opencontrolplane-gen/pkg/commands"
	"github.com/openmcp-project/opencontrolplane-gen/pkg/logs"
	"github.com/openmcp-project/opencontrolplane-gen/pkg/runner"
)

func main() {
	dryRun := commands.EvalBoolEnv("DRY_RUN")
	debug := commands.EvalBoolEnv("DEBUG")

	logs.Init(debug)

	filepath := filepath.Join(os.Getenv("PWD"), os.Getenv("GOFILE"))
	logs.Debug("go:generate opencontrolplane-gen", fmt.Sprintf("opencontrolplane-gen called for file: %s", filepath))

	runner := runner.Runner{
		Commands: []commands.Command{
			commands.NewReplaceCommand(),
			commands.NewIfCommand(),
		},
	}
	result := runner.Run(filepath)

	if !dryRun {
		if err := os.WriteFile(filepath, result.Bytes(), 0644); err != nil {
			log.Fatalf("failed to write to file: %v", err)
		}
		log.Printf("%s: saved changes\n", filepath)
		return
	}

	// dry run prints in memory result unless debug is set
	if !debug {
		if _, err := fmt.Fprintf(os.Stdout, "### %s\n", filepath); err != nil {
			log.Fatalf("failed to write to stdout: %v", err)
		}
		if _, err := fmt.Fprint(os.Stdout, result.String()); err != nil {
			log.Fatalf("failed to write to stdout: %v", err)
		}
	}
}
