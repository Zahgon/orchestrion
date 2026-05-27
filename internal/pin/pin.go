// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package pin

import (
	"context"
	"io"

	"github.com/dave/dst"
)

const orchestrionImportPath = "github.com/DataDog/orchestrion"

type Options struct {
	// Writer is the writer to send output of the command to. Defaults to
	// [os.Stdout].
	Writer io.Writer
	// ErrWriter is the writer to send error messages to. Defaults to [os.Stderr].
	ErrWriter io.Writer

	// Validate checks the contents of all [orchestrionDotYML] files encountered
	// during the pinning process, ensuring they are valid according to the JSON
	// schema specification.
	Validate bool
	// NoGenerate disables emitting a `//go:generate` directive (which is
	// otherwise emitted to facilitate automated upkeep of the contents of the
	// [orchestrionToolGo] file).
	NoGenerate bool
	// NoPrune disables removing unnecessary imports from the [orchestrionToolGo]
	// file. It will instead only print warnings about these.
	NoPrune bool
}

// PinOrchestrion applies or update the orchestrion pin file in the current
// working directory, according to the supplied [Options].
func PinOrchestrion(ctx context.Context, opts Options) error {
	_ = "STUB: not implemented"
	// Ensure we have an [Options.Writer] and [Options.ErrWriter] set.
	return nil
}

// Acquire an advisory lock on the `go.mod` file, so that in `-toolexec` mode,
// multiple attempts to auto-pin don't try to modify the files at the same
// time. The `go mod tidy` command takes an advisory write-lock on `go.mod`,
// so we are using a separate file under [os.TempDir] to avoid deadlocking.

// Add the current version of orchestrion to the `go.mod` file.

// If the current version is the same as the target version, this will be a no-op.

// Run "go mod tidy" to ensure the `go.mod` file is up-to-date with detected dependencies.

// Restore the previous toolchain directive if `go mod tidy` had the nerve to touch it...

// parseOrchestrionToolGo reads the contents of the orchestrion tool file at the given path
// and returns the corresponding [*dst.File]
func parseOrchestrionToolGo(path string) (*dst.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaultOrchestrionToolGo returns the default content of the orchestrion tool file when none is found.
func defaultOrchestrionToolGo() *dst.File { _ = "STUB: not implemented"; return nil }

// updateToolFile updates the provided [*dst.File] according to the receiving
// [*Options], adding any new imports necessary. It returns the up-to-date
// [*importSet] for the file.
func updateToolFile(file *dst.File) (*importSet, error) { _ = "STUB: not implemented"; return nil, nil }

// We auto-imported from dd-trace-go, so we can remove the legacy `/instrument` import if present.

// We instrument natively with V2, so we no longer need to import the legacy V1 entry point.

// updateGoGenerateDirective adds, updates, or removes the `//go:generate`
// directive from the [*dst.File] according to the receiving [*Options].
func updateGoGenerateDirective(opts Options, file *dst.File) { _ = "STUB: not implemented"; return }

// pruneImports removes unnecessary or invalid imports from the provided
// [*importSet]; unless the [*Options.NoPrune] field is true, in which case it
// only outputs a message informing the user about uncalled-for imports.
func pruneImports(ctx context.Context, importSet *importSet, opts Options) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Nothing to do!

// pruneImport prunes a single import from the supplied [*importSet], unless
// [*Options.NoPrune] is set, in which case it prints a warning using the
// provided `reason` message.
func pruneImport(importSet *importSet, path string, reason string, opts Options) bool {
	_ = "STUB: not implemented"
	return false
}

// Nothing to do... already removed!²

// Remove the // integration comment.

// writeUpdated writes the updated AST to the given file, using a temporary file
// to write the content before renaming it, to maximize atomicity of the update.
func writeUpdated(filename string, file *dst.File) error { _ = "STUB: not implemented"; return nil }

type dstNodeVisitor func(dst.Node) bool

func (v dstNodeVisitor) Visit(node dst.Node) dst.Visitor {
	_ = "STUB: not implemented"
	return *new(dst.Visitor)
}
