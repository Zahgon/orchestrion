// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package buildid

import (
	"context"
	"encoding/json"
	"sync"

	"golang.org/x/tools/go/packages"
)

type (
	VersionSuffixRequest  struct{}
	VersionSuffixResponse string
)

func (VersionSuffixRequest) Subject() string                  { _ = "STUB: not implemented"; return "" }
func (VersionSuffixRequest) ResponseIs(VersionSuffixResponse) { _ = "STUB: not implemented"; return }
func (VersionSuffixRequest) ForeachSpanTag(func(string, any)) { _ = "STUB: not implemented"; return }

func (s *service) versionSuffix(ctx context.Context, _ VersionSuffixRequest) (VersionSuffixResponse, error) {
	_ = "STUB: not implemented"
	return *new(VersionSuffixResponse), nil
}

// Explicitly disable toolexec to avoid infinite recursion

type moduleInfo struct {
	*packages.Module
	Files map[string]struct{}
}

// shouldHashContent determines whether this module's files need to be hashed in
// order to properly invalidate build caches when the module changes. This is
// necessary when the module does not have a version, or is replaced by a module
// with no version. Both cases imply the module has been replaced by some local
// directory.
// Interestingly, the [packages.Module.Replace] field is nil for replaced
// modules that belong to the current go.work workspace. This is why we also
// need to verify absence of a [packages.Module.Version] value.
func (m *moduleInfo) shouldHashContent() bool { _ = "STUB: not implemented"; return false }

func collectModules(pkg *packages.Package, modules map[string]*moduleInfo, knownIDs map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

var _ json.Marshaler = (*moduleInfo)(nil)

func (m *moduleInfo) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Ensure a consistent ordering on file names...

var (
	tagSuffix     string
	tagSuffixOnce sync.Once
)

func getTagSuffix(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// If the version is "(devel)", the command was built from a development
// tree. It is typically empty when running test suites (via `test_main`).

// The build has a version... but was it from a clean tree? If not, it still
// is a dev build!

// At this stage we don't think this is a dev build, so we don't need a
// tag suffix.

// We're in a dev build, so we'll add a checksum of this executable as the tag
// suffix, so that development iteration isn't frustrated by needing to clear
// the GOCACHE over and over again.
