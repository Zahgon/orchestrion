// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package proxy

import (
	"context"
)

// ParseCommand parses the Go tool call and its arguments and returns it as a [Command]. The go tool
// call path should be the first element of args. A nil [Command] may be returned if the command is
// to be ignored (as the result of re-using a previous identical command's side effects).
func ParseCommand(ctx context.Context, importPath string, args []string) (Command, error) {
	_ = "STUB: not implemented"
	return *new(Command), nil
}

// There was an error, or we re-used command outputs.

// We currently don't need to inject other tool calls, so we parse them as generic unsupported commands

// MustParseCommand calls ParseCommand and exits on error
func MustParseCommand(ctx context.Context, importPath string, args []string) Command {
	_ = "STUB: not implemented"
	return *new(Command)
}

func parseCommandID(cmd string) (CommandType, error) {
	_ = "STUB: not implemented"
	return *new(CommandType), nil
}

// Take the base of the absolute path of the Go tool

// Depending on the architecture/environment, go tools may have extensions. Remove the extension - if any
