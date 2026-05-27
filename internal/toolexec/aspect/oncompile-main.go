// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package aspect

import (
	"context"

	"github.com/DataDog/orchestrion/internal/toolexec/proxy"
)

// SyntheticPackageName is the name of the synthetic package that will be created when the compilation of the main
// package is performed. This folder contains blank imports for all link-time dependencies that are not already
// in the build tree
var SyntheticPackageName = "synthetic"

// OnCompileMain only performs changes when compiling the "main" package, adding blank imports for
// any linkdeps dependencies that are not yet satisfied by the importcfg file (this is the case for
// link-time dependencies implied by use of the go:linkname directive, which are used to avoid
// creating circular import dependencies).
// This ensures that the relevant packages' `init` (if any) are appropriately run, and that the
// linker automatically picks up these dependencies when creating the full binary.
func (w Weaver) OnCompileMain(ctx context.Context, cmd *proxy.CompileCommand) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Nothing was added, we're done!

// Add package resolutions of link-time dependencies to the importcfg file:

// Pop from the stack of things to process...

// The package may have its own link-time dependencies we need to resolve

// Already resolved, or already going to be resolved...

// Push it to the stack
// Record it as asynthetic import to add
// Record it as a link-time dependency

// We back up the original ImportCfg file only if there's not already such a file (could have been created by OnCompile)

// Generate a synthetic source file with blank imports to link-time
// dependencies, so the linker actually sees them.

// Consistent order for deterministic output
