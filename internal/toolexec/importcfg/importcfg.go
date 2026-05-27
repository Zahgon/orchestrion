// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Package importcfg provides utilities to deal with files accepted by the Go toolchain commands as
// the `-importcfg` flag value.
package importcfg

import (
	"context"
	"io"
)

// ImportConfig represents the parsed out contents of an `importcfg` (or `importcfg.link`) file,
// usually passed to the Go compiler and linker via the `-importcfg` flag.
type ImportConfig struct {
	// PackageFile maps package dependencies fully-qualified import paths to their build archive
	// location
	PackageFile map[string]string
	// ImportMap maps package dependencies import paths to their fully-qualified version
	ImportMap map[string]string
	// Extras is data read from an `importcfg` file that is not semantically parsed by this data
	// structure, which is stored only so it can be written back with an updated `importcfg` file if
	// necessary.
	Extras []string
}

// ParseFile parses the contents of the provided `importcfg` (or `importcfg.link`) file.
func ParseFile(ctx context.Context, filename string) (ImportConfig, error) {
	_ = "STUB: not implemented"
	return *new(ImportConfig), nil
}

// ParseFile parses the `importcfg` (or `importcfg.link`) data from the provided reader.
func parse(r io.Reader) (reg ImportConfig, err error) {
	_ = "STUB: not implemented"
	return *new(ImportConfig), nil
}

// CombinePackageFile copies `packagefile` entries from other into the receiver unless it already
// has an entry with the same import path.
func (r *ImportConfig) CombinePackageFile(other *ImportConfig) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

// WriteFile writes the content of the package register to the provided file, in the format expected
// by the standard go toolchain commands.
func (r *ImportConfig) WriteFile(filename string) error { _ = "STUB: not implemented"; return nil }

// write writes the content of the package register to the provided writer, in the format expected
// by the standard go toolchain commands.
func (r *ImportConfig) write(w io.Writer) error { _ = "STUB: not implemented"; return nil }
