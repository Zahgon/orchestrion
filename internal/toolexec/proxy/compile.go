// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package proxy

import (
	gocontext "context"

	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/toolexec/aspect/linkdeps"
)

//go:generate go run github.com/DataDog/orchestrion/internal/toolexec/proxy/generator -command=compile

type compileFlagSet struct {
	Asmhdr      string `ddflag:"-asmhdr"`
	BuildID     string `ddflag:"-buildid"`
	ImportCfg   string `ddflag:"-importcfg"`
	Lang        string `ddflag:"-lang"`
	Output      string `ddflag:"-o"`
	Package     string `ddflag:"-p"`
	ShowVersion bool   `ddflag:"-V"`
}

// CompileCommand represents a go tool `compile` invocation
type CompileCommand struct {
	command
	Files []string
	Flags compileFlagSet
	// WorkDir is the $WORK directory managed by the go toolchain.
	WorkDir string

	// LinkDeps lists all link-time dependencies that must be honored to link a
	// dependent of the built package. If not blank, this is written to disk, then
	// appended to the archive output.
	LinkDeps linkdeps.LinkDeps

	// importPath is the import path of the package being built.
	importPath string
	// finishToken is the token returned by the job server in response to the
	// [nbt.StartRequest] when the operation needs to continue, and that is then
	// forwarded to the [nbt.FinishRequest].
	finishToken string
}

func (*CompileCommand) Type() CommandType { _ = "STUB: not implemented"; return *new(CommandType) }

func (c *CompileCommand) ShowVersion() bool { _ = "STUB: not implemented"; return false }

// TestMain returns true if the compiled package name is "main" and all source
// Go files are rooted in the same directory as the importcfg file. This
// indicates the package being compiled is a synthetic "main" package generated
// by `go test`. For more accurate readings, users should also validate the
// declared package import path ends in `.test`.
func (c *CompileCommand) TestMain() bool { _ = "STUB: not implemented"; return false }

func (cmd *CompileCommand) SetLang(to context.GoLangVersion) error {
	_ = "STUB: not implemented"

	// No minimal language requirement change, nothing to do...
	return nil
}

// No language level was specified, so anything the compiler can do is possible...

// Minimum language requirement from injected code is already met, nothing to do...

// GoFiles returns the list of Go files passed as arguments to cmd
func (cmd *CompileCommand) GoFiles() []string { _ = "STUB: not implemented"; return nil }

// AddFiles adds the provided go files paths to the list of Go files passed
// as arguments to cmd
func (cmd *CompileCommand) AddFiles(files []string) { _ = "STUB: not implemented"; return }

func (cmd *CompileCommand) Close(ctx gocontext.Context, cmdErr error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Success so far, we attach link-time dependencies...

// Notify the job server of the status of the command, and combine with the previous error if any...

func (cmd *CompileCommand) attachLinkDeps(ctx gocontext.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Already failing, not doing anything...

func (cmd *CompileCommand) notifyJobServer(ctx gocontext.Context, cmdErr error) error {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do...
}

// parseCompileCommand parses a [*CompileCommand] from the provided arguments.
// It sends an [nbt.StartRequest] to the job server to determine whether a
// previous execution of the same command has produced re-usable artifacts;
// in which case it copies them into place and returns nil.
func parseCompileCommand(ctx gocontext.Context, importPath string, args []string) (*CompileCommand, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The WorkDir is the parent of the stage directory, which is where the importcfg file is located.

// We place a "reused" marker next to the output file to identify that it was re-used.

// We place a "reused" marker next to the output file to identify that it was re-used.

var _ Command = (*CompileCommand)(nil)
