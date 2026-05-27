// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package join

import (
	gocontext "context"
	"fmt"
	"regexp"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/aspect/may"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type functionCall struct {
	ImportPath string
	Name       string
}

func FunctionCall(importPath string, name string) *functionCall {
	_ = "STUB: not implemented"
	return nil
}

func (i *functionCall) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (i *functionCall) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (i *functionCall) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (i *functionCall) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

// TODO: Must actually look at whether ident.Name is the import of the relevant package path.

func (i *functionCall) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

// See: https://regex101.com/r/fjLo1l/1
var funcNamePattern = regexp.MustCompile(`\A(?:(.+)\.)?([\p{L}_][\p{L}_\p{Nd}]*)\z`)

func init() {
	unmarshalers["function-call"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var symbol string
		if err := yaml.NodeToValueContext(ctx, node, &symbol); err != nil {
			return nil, err
		}

		matches := funcNamePattern.FindStringSubmatch(symbol)
		if matches == nil {
			return nil, fmt.Errorf("invalid function name %q", symbol)
		}

		return FunctionCall(matches[1], matches[2]), nil
	}
}
