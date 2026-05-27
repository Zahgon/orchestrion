// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Package goflags allows parsing go command invocations and storing their flags in a
// CommandFlags structure. It also provides utilities to backtrack through the process stack to
// find and parse the flags of the first parent go command found in the process hierarchy.
package goflags

import (
	"context"
	"sync"
)

// CommandFlags represents the flags provided to a go command invocation
type CommandFlags struct {
	Long    map[string]string
	Short   map[string]struct{}
	Unknown []string // flags we don't process but store anyway
}

var (
	shortFlags = map[string]struct{}{
		"-a":          {}, // Rebuild everything, ignoring cached artifacts
		"-asan":       {}, // Enables address sanitizer
		"-cover":      {}, // Enables coverage collection
		"-linkshared": {}, // Build code that links against shared libraries
		"-modcacherw": {}, // Keep module cache files read-write
		"-msan":       {}, // Enable memory sanitizer
		"-race":       {}, // Enable data race detection
		"-trimpath":   {}, // Remove all file system paths from the resulting executable
		"-work":       {}, // Keep working temporary directory instead of deleting it
	}
	longFlags = map[string]struct{}{
		"-asmflags":   {}, // Flags passed through to the assembly
		"-buildmode":  {}, // Set build mode
		"-buildvcs":   {}, // Whether to stamp binaries with version control information
		"-compiler":   {}, // Select what compiler to use
		"-covermode":  {}, // Set coverage mode
		"-coverpkg":   {}, // Set list of packages to collect coverage for
		"-gccgoflags": {}, // Flags passed through to the gccgo compiler
		"-gcflags":    {}, // Flags passed through to the gc compiler
		"-ldflags":    {}, // Flags passed through to the linker
		"-mod":        {}, // Set module download mode
		"-modfile":    {}, // Set module file
		"-overlay":    {}, // Set overlay file
		"-pgo":        {}, // Set profile-guided optimization profile file
		"-pkgdir":     {}, // Set package install & load directory
		"-tags":       {}, // Set build tags
		"-toolexec":   {}, // Set the command to run around tool execution
	}
)

// Get returns the value of the specified long-form flag if present. The name is
// provided including the leading hyphen, e.g: "-tags".
func (f CommandFlags) Get(flag string) (val string, found bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Except returns a copy of this CommandFlags with the specified flags removed.
// The [CommandFlags.Unknown] field is not modified, even if it is in the list
// of flags to be removed.
func (f CommandFlags) Except(remove ...string) CommandFlags {
	_ = "STUB: not implemented"
	return *new(CommandFlags)
}

// Slice returns the command flags as a string slice
// - long flags are returned as a string of the form '-flagName="flagVal"'
// - short flags are returned as a string of the form '-flagName'
// - unknown flags and values are ignored
func (f CommandFlags) Slice() []string { _ = "STUB: not implemented"; return nil }

// ParseCommandFlags parses a slice representing a go command invocation
// and returns its flags. Direct arguments to the command are ignored. The value
// of $GOFLAGS is also included in the returned flags.
func ParseCommandFlags(ctx context.Context, wd string, args []string) (CommandFlags, error) {
	_ = "STUB: not implemented"
	return *new(CommandFlags), nil
}

// Remove any `-C` flag provided on the command line. This is required to immediately follow the `go` command, and
// can be present only once.

// ["-C", "directory", ...]

// ["-C=directory", ...]

// The next argument after a `-C` (if present) would be the go command name ("run", "test", "list", etc...). This is
// not interesting for our purposes, so we skip it.

// Compose the complete list of arguments: those from GOFLAGS, and the rest of the command line so far; in this order
// as the CLI arguments have precedence over those from GOFLAGS.

// Any argument after "--" is a positional argument, so we are done parsing.

// Any argument without a leading "-" is a positional argument (until proven otherwise).

// The Go CLI accepts flags with two hyphens instead of one, but we want
// to normalize to a single hyphen here...

// Intentionally the un-normalized variant in Unknown flags.

// Intentionally the un-normalized variant in Unknown flags.

// If there's more args, and the next one does not have a leading -, we'll assume this is the value of this
// unknown flag and consume it.

// inferCoverpkg will add the necessary `-coverpkg` argument if the `-cover` flags is present and
// `-coverpkg` is not, as otherwise, sub-commands triggered with these flags will not apply coverage
// to the intended packages.
// If `-coverpkg` is present, it will expand any relative paths (recognized by a `./` prefix) into
// absolute package names, so that child builds do not interpret these relative to a different
// package root.
func (f *CommandFlags) inferCoverpkg(ctx context.Context, wd string, positionalArgs []string) error {
	_ = "STUB: not implemented"
	return nil

	// Make sure we satisfy the same build constraints; but don't run -toolexec
}

// Blank specified, not trying to expand it...

// We have patterns, we need to make sure they are expressed in absolute terms.

// If the pattern is not relative, so we're good.

// -covermode implies -cover

// Flags return the top level go command flags
func Flags(ctx context.Context) (CommandFlags, error) {
	_ = "STUB: not implemented"
	return *new(CommandFlags), nil
}

// SetFlagsFromPid sets the top level go command flags by looking up the process
// tree from the specified PID. This is used by the job server when it is
// started as a daemon (and hence cannot crawl it's own process tree to find
// this information).
func SetFlagsFromPid(ctx context.Context, pid int) error { _ = "STUB: not implemented"; return nil }

// SetFlags sets the flags for this process to those parsed from the provided
// slice. Does nothing if SetFlags or Flags has already been called once.
func SetFlags(ctx context.Context, wd string, args []string) { _ = "STUB: not implemented"; return }

func isLong(str string) bool { _ = "STUB: not implemented"; return false }

func isShort(str string) bool { _ = "STUB: not implemented"; return false }

// parentGoCommandFlags backtracks through the process tree
// to find a parent go command invocation and returns its arguments
func parentGoCommandFlags(ctx context.Context, pid int) (flags CommandFlags, err error) {
	_ = "STUB: not implemented"
	return *new(CommandFlags), nil
}

// Backtrack through the process stack until we find the parent Go command

// When running in containers using on macOS VZ+rosetta, the reported command line may be led by
// the registered rosetta binfmt handler. In such cases, the argv0 has a leaf name of "rosetta"
// and is not present within the container itself (it's only on the hypervisor). In such cases,
// we try to resolve argv[1] instead. This can only manifest itself on amd64 + linux.

// The fallback was successful, we no longer have an error!

// Found the go command process, break out of backtracking

var (
	flags    CommandFlags
	flagsErr error
	once     sync.Once
)
