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
	run(output)
}

// run writes the CLI output to the provided writer so tests can verify it.
func run(output io.Writer) {
	_, _ = fmt.Fprintln(output, greeting())
}

// greeting returns the static hello-world text for the default CLI output.
func greeting() string {
	return "Hello, world!"
}
