// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package pkgs

import (
	"context"

	"golang.org/x/tools/go/packages"
)

type (
	// LoadRequest is a request to load packages relative to a specific directory. It only loads the
	// packages' names and (source) files, not their dependencies or export file. The result is cached
	// for a given Dir+Pattern pair. Each pattern is loaded individually (so that they can be cached
	// independently).
	LoadRequest struct {
		Dir      string   `json:"dir"`      // The directory to resolve from (usually where `go.mod` is)
		Patterns []string `json:"patterns"` // Package pattern to resolve
	}
	// LoadResponse is the response to a [LoadRequest]. It contains the packages that were loaded.
	LoadResponse []*packages.Package
)

// packageLoader can be used as a [config.PackageLoader] implementation.
func (s *service) packageLoader(ctx context.Context, dir string, patterns ...string) ([]*packages.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (LoadRequest) Subject() string         { _ = "STUB: not implemented"; return "" }
func (LoadRequest) ResponseIs(LoadResponse) { _ = "STUB: not implemented"; return }
func (r LoadRequest) ForeachSpanTag(set func(key string, value any)) {
	_ = "STUB: not implemented"
	return
}

func (s *service) load(ctx context.Context, req LoadRequest) (LoadResponse, error) {
	_ = "STUB: not implemented"
	return *new(LoadResponse), nil
}

// Re-building everything here would be VERY expensive, as we'd re-build a lot of stuff multiple times
// We'll override `-toolexec` later with `orchestrion toolexec`, no need to pass multiple times...

// Explicitly disable toolexec if it's in GOFLAGS
