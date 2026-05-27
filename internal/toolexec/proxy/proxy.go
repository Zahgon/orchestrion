// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package proxy

import (
	"context"
	"os/exec"
)

type (
	// CommandType represents a Go toolchain command type, such
	// as "compile", "link", etc...
	CommandType int
	// Command represents a Go compilation command
	Command interface {
		// Close invokes all registered OnClose callbacks and releases any resources associated with the
		// command. If the command failed, the error is provided as an argument.
		Close(context.Context, error) error

		// Args are all the command arguments, starting from the Go tool command
		Args() []string
		ReplaceParam(param string, val string) error

		// Type represents the go tool command type (compile, link, asm, etc.)
		Type() CommandType

		// ShowVersion returns true if the command received the `-V=full` argument, signaling it should
		// print its full version information and exit. This feature is used by the go toolchain to
		// create build cache keys, and allows invalidating all build cache when the tooling changes.
		ShowVersion() bool
	}

	// CommandProcessor is a function that takes a command as input and is allowed to modify it or
	// read its data. If it returns an error, the processing chain immediately stops and no further
	// processors will be invoked.
	CommandProcessor[T Command] func(context.Context, T) error

	// command is the default unknown command type
	// Can be used to compose specific Command implementations
	command struct {
		args []string
		// paramPos is the index in args of the *value* provided for the parameter stored in the key
		paramPos map[string]int
	}
)

const (
	CommandTypeOther CommandType = iota
	CommandTypeCompile
	CommandTypeLink
)

func (t CommandType) String() string { _ = "STUB: not implemented"; return "" }

// ProcessCommand applies a processor on a command if said command matches
// the input type of said input processor. Nothing happens if the processor does
// not correspond to the provided command type.
func ProcessCommand[T Command](ctx context.Context, cmd Command, p CommandProcessor[T]) error {
	_ = "STUB: not implemented"
	return nil
}

// NewCommand initializes a new command object and takes care of tracking the indexes of its
// arguments
func NewCommand(args []string) command { _ = "STUB: not implemented"; return *new(command) }

func (*command) Close(context.Context, error) error {
	_ = "STUB: not implemented"
	// Nothing to do...
	return nil
}

// SetFlag replaces the value of the specified flag with the provided one.
// Returns an error if the flag is not present in the current arguments list.
func (cmd *command) SetFlag(flag string, val string) error { _ = "STUB: not implemented"; return nil }

// ReplaceParam will replace any parameter of the command provided it is found
// A parameter can be a flag, an option, a value, etc
func (cmd *command) ReplaceParam(param string, val string) error {
	_ = "STUB: not implemented"
	return nil
}

// RunCommandOption allows customizing a run command before execution. For example, this can be used
// to capture the output of the command instead of forwarding it to the host process' STDIO.
type RunCommandOption func(*exec.Cmd)

// RunCommand executes the underlying go tool command and forwards the program's standard fluxes
func RunCommand(ctx context.Context, cmd Command, opts ...RunCommandOption) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (*command) Type() CommandType { _ = "STUB: not implemented"; return *new(CommandType) }

func (cmd *command) Args() []string { _ = "STUB: not implemented"; return nil }

func (*command) ShowVersion() bool { _ = "STUB: not implemented"; return false }
