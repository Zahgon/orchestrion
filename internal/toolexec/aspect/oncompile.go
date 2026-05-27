// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package aspect

import (
	"context"
	"path/filepath"

	"github.com/DataDog/orchestrion/internal/injector/config"
	"github.com/DataDog/orchestrion/internal/jobserver/client"
	"github.com/DataDog/orchestrion/internal/toolexec/importcfg"
	"github.com/DataDog/orchestrion/internal/toolexec/proxy"
	"github.com/rs/zerolog"
)

// OrchestrionDirPathElement is the prefix for orchestrion source files in the build output directory.
var OrchestrionDirPathElement = filepath.Join("orchestrion", "src")

func (w Weaver) OnCompile(ctx context.Context, cmd *proxy.CompileCommand) (resErr error) {
	_ = "STUB: not implemented"
	return nil
}

// No-op

// Unreachable

// Unsafe isn't like other go packages, and it does not have an associated archive file.

// Already part of natural dependencies, nothing to do...

// We cannot attempt to resolve link-time dependencies (relocation targets), as these are
// typically used to avoid creating dependency cycles. Corollary to this, the `link.deps`
// file will not contain transitive closures for these packages, so we need to resolve these
// at link-time. If the package being built is "main", then we can ignore this, as we are at
// the top-level of a dependency tree anyway, and if we cannot resolve a dependency, then we
// will not be able to link the final binary.

// Imported packages need to be provided in the compilation's importcfg file

// Already part of natural dependencies, nothing to do...

// Creating updated version of the importcfg file, with new dependencies

func writeUpdatedImportConfig(log zerolog.Logger, reg importcfg.ImportConfig, filename string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func packageLoader(js *client.Client) config.PackageLoader {
	_ = "STUB: not implemented"
	return *new(config.PackageLoader)
}
