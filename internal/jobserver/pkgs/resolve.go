// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package pkgs

import (
	"context"

	"golang.org/x/tools/go/packages"
)

const (
	envVarParentID = "ORCHESTRION_PKG.RESOLVE_PARENT_ID"
	envVarGotmpdir = "GOTMPDIR"
)

var envIgnoreList = map[string]func(*ResolveRequest, string){
	// We don't use this, instead rely on the [ResolveRequest.Dir] field.
	"PWD": nil,
	// We override `GOTMPDIR` with the [ResolveRequest.TempDir] field.
	envVarGotmpdir: func(r *ResolveRequest, dir string) {
		if r.TempDir != "" {
			return
		}
		r.TempDir = dir
	},
	// Known to change between invocations & irrelevant to the resolution, but can be used to detect cycles.
	"TOOLEXEC_IMPORTPATH": func(r *ResolveRequest, path string) { r.toolexecImportpath = path },
	envVarParentID:        func(r *ResolveRequest, id string) { r.resolveParentID = id },
}

type (
	ResolveRequest struct {
		Dir     string   `json:"dir"`              // The directory to resolve from (usually where `go.mod` is)
		Env     []string `json:"env"`              // Environment variables to use during resolution
		Pattern string   `json:"pattern"`          // Package pattern to resolve
		TempDir string   `json:"tmpdir,omitempty"` // A temporary directory to use for Go build artifacts

		// Fields set by canonicalization
		resolveParentID    string // The value of the [envVarParentID] environment variable
		toolexecImportpath string // The value of the TOOLEXEC_IMPORTPATH environment variable
		canonical          bool   // Whether this request was canonicalized yet
	}
	// ResolveResponse is a map of package import path to their respective export file, if one is
	// found. Users should handle possibly missing export files as is relevant to their use-case.
	ResolveResponse map[string]string
)

func NewResolveRequest(dir string, pattern string) ResolveRequest {
	_ = "STUB: not implemented"
	return *new(ResolveRequest)
}

func (ResolveRequest) Subject() string            { _ = "STUB: not implemented"; return "" }
func (ResolveRequest) ResponseIs(ResolveResponse) { _ = "STUB: not implemented"; return }
func (r ResolveRequest) ForeachSpanTag(set func(key string, value any)) {
	_ = "STUB: not implemented"
	return
}

func (r *ResolveRequest) canonicalizeEnviron() { _ = "STUB: not implemented"; return }

func (s *service) resolve(ctx context.Context, req *ResolveRequest) (ResolveResponse, error) {
	_ = "STUB: not implemented"
	return *

	// Make sure all children jobs connect to THIS jobserver; this is more efficient than checking for
	// the local file system beacon.
	new(ResolveResponse), nil
}

// Make sure the directory exists (go blindly assumes that...)

// Re-building everything here would be VERY expensive, as we'd re-build a lot of stuff multiple times
// We'll override `-toolexec` later with `orchestrion toolexec`, no need to pass multiple times...

// We need the export file (the whole point of the resolution)

// We want to also resolve transitive dependencies, so we need Deps & Imports. We also
// need CompiledGoFiles in order to see imports possibly added by the toolchain (cgo,
// cover, etc...)

// Finally, we need the resolved package import path

func (r *ResolveRequest) canonicalize() { _ = "STUB: not implemented"; return }

func (r *ResolveRequest) hash() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r ResolveResponse) mergeFrom(pkg *packages.Package) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore the "unsafe" package (no archive file, ever), packages with an empty import path
// (standard library), and those already present in the map (already processed previously).
