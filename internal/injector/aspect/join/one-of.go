// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package join

import (
	gocontext "context"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/aspect/may"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type oneOf []Point

func OneOf(candidates ...Point) oneOf { _ = "STUB: not implemented"; return *new(oneOf) }

func (o oneOf) ImpliesImported() []string {
	_ = "STUB: not implemented"
	// We can only assume a package is imported if all candidates imply it.
	return nil
}

func (o oneOf) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (o oneOf) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (o oneOf) Matches(ctx context.AspectContext) bool { _ = "STUB: not implemented"; return false }

func (o oneOf) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["one-of"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var nodes []ast.Node
		if err := yaml.NodeToValueContext(ctx, node, &nodes); err != nil {
			return nil, err
		}

		if len(nodes) == 1 {
			pt, err := FromYAML(ctx, nodes[0])
			return pt, err
		}

		candidates := make([]Point, len(nodes))
		for i, n := range nodes {
			var err error
			if candidates[i], err = FromYAML(ctx, n); err != nil {
				return nil, err
			}
		}
		return OneOf(candidates...), nil
	}
}
