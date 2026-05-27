// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package proxy

import (
	"context"
)

//go:generate go run github.com/DataDog/orchestrion/internal/toolexec/proxy/generator -command=link

type linkFlagSet struct {
	BuildMode   string `ddflag:"-buildmode"`
	ImportCfg   string `ddflag:"-importcfg"`
	Output      string `ddflag:"-o"`
	ShowVersion bool   `ddflag:"-V"`
}

// LinkCommand represents a go tool `link` invocation
type LinkCommand struct {
	command
	Flags linkFlagSet
	// WorkDir is the $WORK directory managed by the go toolchain.
	WorkDir string
}

func (*LinkCommand) Type() CommandType { _ = "STUB: not implemented"; return *new(CommandType) }

func (cmd *LinkCommand) ShowVersion() bool { _ = "STUB: not implemented"; return false }

func (cmd *LinkCommand) Stage() string { _ = "STUB: not implemented"; return "" }

func parseLinkCommand(_ context.Context, args []string) (Command, error) {
	_ = "STUB: not implemented"
	return *new(Command), nil
}

// The WorkDir is the parent of the stage dir, and the ImportCfg file is directly in the stage dir.
