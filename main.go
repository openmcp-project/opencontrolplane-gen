package main

import (
	"fmt"
	"log"
	"os"

	"github.com/christophrj/opencontrolplane-gen/commands"
	"github.com/christophrj/opencontrolplane-gen/logs"
	"github.com/christophrj/opencontrolplane-gen/runner"
)

func main() {
	dryRun := commands.EvalBoolEnv("DRY_RUN")
	debug := commands.EvalBoolEnv("DEBUG")

	logs.Init(debug)

	filepath := os.Getenv("PWD") + "/" + os.Getenv("GOFILE")
	logs.Debug(fmt.Sprintf("opencontrolplan-gen called for file: %s", filepath))

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
