// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package goproxy

import (
	"context"
)

// processDashC gets the command line arguments passed to a "go" command (without "go" itself), and processes the "-C"
// flag present at the beginning of the slice (if present), changing directories as requested, then returns the slice
// without it.
//
// The "-C" flags is required to be the very first argument provided to "go" commands.
func processDashC(ctx context.Context, args []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ["-C", "directory", ...]

// Probably not the flag we're looking for... ignoring that...

// ["-C=directory", ...]
