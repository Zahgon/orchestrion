// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package gomod

import (
	"context"
	"io"
)

type (
	// File represents some selected entries from the content of a `go.mod` file.
	File struct {
		// Go is the value of the `go` directive.
		Go Version
		// Toolchain is the value of the `toolchain` directive, if present.
		Toolchain Toolchain
		// Require is a list of all `require` directives' contents.
		Require []Require
	}

	// Edit represents an edition that can be made to a `go.mod` file via `go mod edit`.
	Edit interface {
		goModEditFlag() string
	}

	Version   string
	Toolchain string

	// Require represents the target of a `require` directive entry.
	Require struct {
		// Path is the path of the required module.
		Path string
		// Version is the required module's version.
		Version string
	}

	// Replace represents the target of a `replace` directive entry.
	Replace struct {
		// OldPath is the path of the module being replaced.
		OldPath string
		// OldVersion is the version of the module being replaced, if any.
		OldVersion string
		// NewPath is the path of the replacement module.
		NewPath string
		// NewVersion is the version of the replacement module, if any.
		NewVersion string
	}
)

// Parse processes the contents of the designated `go.mod` file using
// `go mod edit -json` and returns the corresponding parsed [goMod].
func Parse(ctx context.Context, modfile string) (File, error) {
	_ = "STUB: not implemented"
	return *new(File), nil
}

// Requires returns true if the `go.mod` file contains a require directive for
// the designated module path.
func (m *File) Requires(path string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// RunGet executes the `go get <modSpecs...>` subcommand with the provided
// module specifications on the designated `go.mod` file.
func RunGet(ctx context.Context, modfile string, modSpecs ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run executes the `go mod <command> <args...>` subcommand with the
// provided arguments on the designated `go.mod` file, sending standard output
// to the provided writer.
func Run(ctx context.Context, command string, modfile string, stdout io.Writer, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// RunEdit makes the specified changes to the `go.mod` file, then runs `go mod tidy` if needed.
// If there is a `vendor` directory, it also runs `go mod vendor` before returning.
func RunEdit(ctx context.Context, modfile string, edits ...Edit) error {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do.
}

//  No `vendor` directory, nothing to do...

func (v Version) goModEditFlag() string { _ = "STUB: not implemented"; return "" }

func (t Toolchain) goModEditFlag() string { _ = "STUB: not implemented"; return "" }

func (r Require) goModEditFlag() string { _ = "STUB: not implemented"; return "" }

func (r Replace) goModEditFlag() string { _ = "STUB: not implemented"; return "" }
