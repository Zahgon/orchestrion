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

	_ "embed" // For go:embed
)

type configuration map[string]string

func Configuration(requirements map[string]string) configuration {
	_ = "STUB: not implemented"
	return *new(configuration)
}

func (configuration) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (configuration) PackageMayMatch(_ *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (configuration) FileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (jp configuration) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (jp configuration) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["configuration"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var c configuration
		return c, yaml.NodeToValueContext(ctx, node, &c)
	}
}
