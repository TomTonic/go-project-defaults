package main

import (
	"bytes"
	"testing"
)

// TestGreetingReturnsHelloWorld verifies that users see the expected hello-world
// message when they use the default CLI output.
//
// This test covers the main package entry behavior for the example command in
// this repository.
//
// It verifies that the greeting helper returns the exact text printed by the
// command.
func TestGreetingReturnsHelloWorld(t *testing.T) {
	if got := greeting(); got != "Hello, world!" {
		t.Fatalf("greeting() = %q, want %q", got, "Hello, world!")
	}
}

// TestRunWritesHelloWorld verifies that users receive the hello-world line on
// standard output when they execute the command.
//
// This test covers the CLI output path in the repository's example command.
//
// It verifies that the command writes the expected line including the trailing
// newline required for terminal-friendly output.
func TestRunWritesHelloWorld(t *testing.T) {
	var output bytes.Buffer
	run(&output)

	if got := output.String(); got != "Hello, world!\n" {
		t.Fatalf("run() output = %q, want %q", got, "Hello, world!\n")
	}
}

// TestMainWritesHelloWorld verifies that users receive the expected greeting
// when they invoke the compiled command entry point.
//
// This test covers the top-level CLI bootstrap in the repository's example
// application.
//
// It verifies that main routes its output through the configured writer so the
// command prints the same hello-world line as the lower-level helper.
func TestMainWritesHelloWorld(t *testing.T) {
	var captured bytes.Buffer
	originalOutput := output
	output = &captured
	t.Cleanup(func() {
		output = originalOutput
	})

	main()

	if got := captured.String(); got != "Hello, world!\n" {
		t.Fatalf("main() output = %q, want %q", got, "Hello, world!\n")
	}
}
