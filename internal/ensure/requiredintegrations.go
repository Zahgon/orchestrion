// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package ensure

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/DataDog/orchestrion/internal/gomod"
)

func RequiredIntegrations(ctx context.Context, goMod string) ([]gomod.Edit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// V1

// V2

// We install/upgrade the `orchestrion/all/v2` module as it includes all interesting contribs in its dependency
// closure, so we don't have to manually verify all of them. The `go mod tidy` later will clean up if needed.

// versionFetcher is a function type for fetching latest versions
type versionFetcher func(ctx context.Context, modPath string) (string, error)

type versions struct {
	found   bool
	current string
	shipped string
	latest  string
}

func fetchVersions(ctx context.Context, curMod gomod.File, integration string, fetcher versionFetcher) (*versions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolveIntegrationVersion determines if the specified integration should be upgraded, and to which version.
func resolveIntegrationVersion(ver *versions) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// Only run go get if we need to upgrade or if the module is not present.

// Force upgrade, otherwise, go mod tidy will fail.

// fetchLatestVersion queries the Go module registry to get the actual latest version
// of the specified module path.
func fetchLatestVersion(ctx context.Context, modPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Build environment with GOTOOLCHAIN=local and explicit -mod=mod
// This is necessary because GOFLAGS might contain -mod=vendor which prevents
// querying the module registry for @latest versions.

// Remove any existing GOFLAGS that might contain -mod=vendor

// Set GOFLAGS with -mod=mod to ensure we can query the registry

func maxVersion(versions ...string) string { _ = "STUB: not implemented"; return "" }

func resolveDependencyVersion(modDir string, dependency string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var (
	initOnce                   sync.Once
	orchestrionShippedVersions = atomic.Pointer[map[string]string]{}
)

func fetchShippedVersions() map[string]string { _ = "STUB: not implemented"; return nil }

func loadShippedVersions(orchestrionRoot string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// The version of dd-trace-go that shipped with the current version of orchestrion.
// We use this to determine if we need to upgrade dd-trace-go when pinning.

// Default to using the same version as DatadogTracerV2 for DatadogTracerV2All.
// This serves as a fallback when the instrument directory (which is a separate submodule)
// is not accessible, such as when orchestrion is used as a module dependency from the
// Go module cache where nested submodules may not be present.

// If the instrument directory exists, override with the actual version from its go.mod
