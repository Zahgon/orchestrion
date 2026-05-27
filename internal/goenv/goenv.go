// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package goenv

import (
	"context"
	"errors"
	"sync"
)

var (
	// ErrNoGoMod is returned when no GOMOD value could be identified.
	ErrNoGoMod = errors.New("`go env GOMOD` returned a blank string")

	// ErrNoModulePath is returned when no module path could be identified.
	ErrNoModulePath = errors.New("no module path found")
)

// Module represents basic information about a Go module.
type Module struct {
	Path string `json:"Path"`
	Dir  string `json:"Dir"`
}

var (
	muCache sync.RWMutex
	// Cache for module path lookups to avoid repeated calls to go list.
	modulePathCache = make(map[string]string)
)

// GOMOD returns the current GOMOD environment variable (from running `go env GOMOD`).
func GOMOD(dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// modulePath returns the module path of the current module using go/packages API.
// Results are cached to avoid repeated package loading calls.
func modulePath(ctx context.Context, dir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RootModulePath returns the root module path for the current working directory.
// This is a convenience function that calls ModulePath with the current directory.
func RootModulePath(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	// Getwd returns an absolute path name corresponding to the current directory.
	return "", nil
}
