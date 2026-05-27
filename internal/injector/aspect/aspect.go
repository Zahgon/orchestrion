// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package aspect

import (
	"context"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/advice"
	"github.com/DataDog/orchestrion/internal/injector/aspect/join"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

// Aspect binds advice.Advice to a join.Point, effectively defining a complete
// code injection.
type Aspect struct {
	// JoinPoint determines whether the injection should be performed on a given node or not.
	JoinPoint join.Point
	// Advice is the set of actions to use for performing the actual injection.
	Advice []advice.Advice
	// TracerInternal determines whether the aspect can be woven into the tracer's internal code.
	TracerInternal bool
	// ID is the identifier of the aspect within its configuration file.
	ID string
}

func (a *Aspect) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func (a *Aspect) AddedImports() (imports []string) {
	_ = "STUB: not implemented"
	// "unsafe" is always implied, because it's special-cased in the go toolchain, and is not a "normal" module.
	return nil
}

// InjectedPaths returns the list of import paths that may be injected by the
// supplied list of aspects. The output list is not sorted in any particular way
// but does not contain duplicated entries.
func InjectedPaths(list []*Aspect) []string { _ = "STUB: not implemented"; return nil }

func (a *Aspect) UnmarshalYAML(ctx context.Context, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	_ yaml.NodeUnmarshalerContext = (*Aspect)(nil)
)
