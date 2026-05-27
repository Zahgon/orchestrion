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

type testMain bool

func (t testMain) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (testMain) FileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *

	// TestMain matches only nodes in ASTs in files that either are (if true), or
	// are not (if false) part of a synthetic test main package.
	new(may.MatchType)
}

func TestMain(v bool) testMain { _ = "STUB: not implemented"; return *new(testMain) }

func (t testMain) Matches(ctx context.AspectContext) bool { _ = "STUB: not implemented"; return false }

func (testMain) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (t testMain) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["test-main"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var val bool
		if err := yaml.NodeToValueContext(ctx, node, &val); err != nil {
			return nil, err
		}
		return TestMain(val), nil
	}
}
