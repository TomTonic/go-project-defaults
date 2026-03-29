// Package main provides a minimal hello-world CLI used to validate the
// repository workflows and release configuration.
package main

import (
	"fmt"
	"io"
	"os"
)

var (
	version           = "dev"
	commit            = "unknown"
	date              = "unknown"
	output  io.Writer = os.Stdout
)

func main() {
	run(os.Args[1:], output)
}

// run writes the CLI output to the provided writer so tests can verify it.
func run(args []string, output io.Writer) {
	if len(args) > 0 && (args[0] == "version" || args[0] == "--version") {
		_, _ = fmt.Fprintln(output, versionString())
		return
	}

	_, _ = fmt.Fprintln(output, greeting())
}

// greeting returns the static hello-world text for the default CLI output.
func greeting() string {
	return "Hello, world!"
}

// versionString returns release metadata injected by GoReleaser ldflags.
func versionString() string {
	return fmt.Sprintf("version=%s commit=%s date=%s", version, commit, date)
}
