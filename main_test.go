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
	run(nil, &output)

	if got := output.String(); got != "Hello, world!\n" {
		t.Fatalf("run() output = %q, want %q", got, "Hello, world!\n")
	}
}

// TestRunWritesVersionMetadata verifies that users can inspect the injected
// release metadata when they ask the command for its version.
//
// This test covers the release-information path in the repository's example
// CLI application.
//
// It verifies that the command emits the version, commit, and build date values
// provided by the current build.
func TestRunWritesVersionMetadata(t *testing.T) {
	originalVersion := version
	originalCommit := commit
	originalDate := date
	version = "1.2.3"
	commit = "abc1234"
	date = "2026-03-29T12:00:00Z"
	t.Cleanup(func() {
		version = originalVersion
		commit = originalCommit
		date = originalDate
	})

	var output bytes.Buffer
	run([]string{"version"}, &output)

	want := "version=1.2.3 commit=abc1234 date=2026-03-29T12:00:00Z\n"
	if got := output.String(); got != want {
		t.Fatalf("run(version) output = %q, want %q", got, want)
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
