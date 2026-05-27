// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package goproxy

import (
	"context"
	"os/exec"
)

type config struct {
	toolexec string
}

type Option func(*config)

// WithToolexec forces a call to Run() to build with the -toolexec option when
// wrapping a build command
func WithToolexec(bin string, args ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// This is expected to never happen (short of running OOM, maybe?)

// We are quoting all arguments to hopefully evade shell interpretation.

// This is expected to never happen (short of running OOM, maybe?)

// Dispose of the buffer's content so its memory can be recclaimed.

// BuildCmd returns a new exec.BuildCmd that will run the given goArgs, with the given opts applied.
func BuildCmd(ctx context.Context, goArgs []string, opts ...Option) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pre-allocate space for extra arguments...

// "go build" arguments are shared by build, clean, get, install, list, run, and test.

// Add two slots to the argV array

// Move all values after the cmdIdx 2 slots forward

// Fill in the two slots for toolexec.

// We'll need a job server to support toolexec operations

// Set the process' goflags, since we know them already...

// Run takes a go command ("build", "install", etc...) with its arguments, and
// applies changes specified through opts to the command before running it in a
// different process.
func Run(ctx context.Context, goArgs []string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}
