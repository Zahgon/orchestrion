// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package linkdeps

import (
	"bufio"
	"context"
	"io"

	"github.com/DataDog/orchestrion/internal/toolexec/importcfg"
	"github.com/blakesmith/ar"
	"github.com/rs/zerolog"
)

const (
	// Filename is the standard file name for link.deps files.
	Filename = "link.deps"

	headerV1 = "#" + Filename + "@v1"
)

// LinkDeps represents the contents of a [Filename] file. It lists all synthetic
// dependencies added by instrumentation into a Go object archive. These include
// the transitive closure of link-time dependencies, so that it is not necessary
// to perform a full traversal of transitive dependencies to consume. Link-time
// dependencies consist of new dependencies introduced to resolve go:linkname
// directives as well as new import-level directives that the Go toolchain is
// not normally aware of.
type LinkDeps struct {
	deps map[string]struct{}
}

// FromImportConfig aggregates entries from all [Filename] found in the
// archives listed in [importcfg.ImportConfig].
func FromImportConfig(ctx context.Context, importcfg *importcfg.ImportConfig) (LinkDeps, error) {
	_ = "STUB: not implemented"
	return *new(LinkDeps), nil
}

// This transitive link-time dependency is already satisfied at
// compile-time, so we don't need to carry it over.

// FromArchive reads a [Filename] file from the provided Go archive file.
// Returns an empty [LinkDeps] if the archive does not contain a [Filename]
// file.
func FromArchive(ctx context.Context, archive string) (res LinkDeps, err error) {
	_ = "STUB: not implemented"
	return *new(LinkDeps), nil
}

// ReadFile reads a [Filename] file from the provided filename.
func ReadFile(filename string) (LinkDeps, error) {
	_ = "STUB: not implemented"
	return *new(LinkDeps), nil
}

// Read reads a [Filename] file content from the provided [io.Reader].
func Read(r io.Reader) (l LinkDeps, err error) {
	_ = "STUB: not implemented"
	return *new(LinkDeps), nil
}

// parseV1 parses the contents of V1 [Filename] files.
func parseV1(r *bufio.Reader) (l LinkDeps, err error) {
	_ = "STUB: not implemented"
	return *new(LinkDeps), nil
}

// Contains checks whether a given import path is already represented by this
// [LinkDeps].
func (l *LinkDeps) Contains(importPath string) bool { _ = "STUB: not implemented"; return false }

// Add registers a new import path in this [LinkDeps] instance.
func (l *LinkDeps) Add(importPath string) { _ = "STUB: not implemented"; return }

// Dependencies returns all import paths registered in this [LinkDeps] instance.
func (l *LinkDeps) Dependencies() []string { _ = "STUB: not implemented"; return nil }

// Empty returns true if this [LinkDeps] instance is empty.
func (l *LinkDeps) Empty() bool { _ = "STUB: not implemented"; return false }

// Len returns the number of import paths registered in this [LinkDeps]
// instance.
func (l *LinkDeps) Len() int {
	_ = "STUB: not implemented"

	// Write writes this [LinkDeps] instance to the provided writer.
	return 0
}

func (l *LinkDeps) Write(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// We sort entries to ensure the output is deterministic, since these files
// eventually get embedded in `_pkg_.a` files and we wouldn't want to cause
// unnecessary rebuilds.

// readArchiveData returns the content of the given entry from the provided archive file. If there
// is no such entry in the archive, a nil io.Reader and no error is returned.
func readArchiveData(archive string, entry string) (rc io.ReadCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// If we return no [io.ReadCloser], then we need to close the file ourselves.

var _ zerolog.LogArrayMarshaler = (*LinkDeps)(nil)

func (l *LinkDeps) MarshalZerologArray(a *zerolog.Array) { _ = "STUB: not implemented"; return }

type arReadCloser struct {
	*ar.Reader
	io.Closer
}
