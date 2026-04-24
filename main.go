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
		//
		// NOTE (personal): upstream uses os.Exit(1) here too — filed an issue
		// suggesting they expose a typed ExitError so callers can distinguish
		// usage errors (2) from runtime errors (1). Tracking in FORK_NOTES.md.
		//
		// UPDATE: checked upstream issue tracker — looks like they're planning
		// to add ExitError in v0.5.0. Will revisit this TODO then.
		os.Exit(1)
	}
}
