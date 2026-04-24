// Package main is the entry point for the entireio/cli fork.
// This CLI tool provides a streamlined interface for interacting with
// the Entire platform, including search, analysis, and development workflows.
//
// Personal fork notes:
// - Customized for my own learning and experimentation
// - See FORK_NOTES.md for changes from upstream
package main

import (
	"fmt"
	"os"

	"github.com/entireio/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		// Use exit code 2 for usage errors vs 1 for general errors;
		// keeping it simple here with a single non-zero code for now.
		// TODO: differentiate exit codes once cmd package exposes error types.
		os.Exit(1)
	}
}
