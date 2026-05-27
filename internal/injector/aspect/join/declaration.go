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
	"github.com/DataDog/orchestrion/internal/injector/typed"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type declarationOf struct {
	ImportPath string
	Name       string
}

// DeclarationOf matches the (top-level) declaration of the specified symbol.
func DeclarationOf(importPath string, name string) *declarationOf {
	_ = "STUB: not implemented"
	return nil
}

func (i *declarationOf) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

// No parent, this is almost certainly a syntax error...

// Parent isn't a GenDecl, so this is not a top-level declaration.

func (i *declarationOf) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (i *declarationOf) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (i *declarationOf) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (i *declarationOf) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type valueDeclaration struct {
	TypeName typed.TypeName
}

func ValueDeclaration(typeName typed.TypeName) *valueDeclaration {
	_ = "STUB: not implemented"
	return nil
}

func (i *valueDeclaration) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (*valueDeclaration) FileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (i *valueDeclaration) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (i *valueDeclaration) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (i *valueDeclaration) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

// See: https://regex101.com/r/OXDfJ1/1
var symbolNamePattern = regexp.MustCompile(`\A(.+)\.([\p{L}_][\p{L}_\p{Nd}]*)\z`)

func init() {
	unmarshalers["declaration-of"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var symbol string
		if err := yaml.NodeToValueContext(ctx, node, &symbol); err != nil {
			return nil, err
		}

		matches := symbolNamePattern.FindStringSubmatch(symbol)
		if matches == nil {
			return nil, fmt.Errorf("invalid symbol name %q", symbol)
		}

		return DeclarationOf(matches[1], matches[2]), nil
	}

	unmarshalers["value-declaration"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var typeName string
		if err := yaml.NodeToValueContext(ctx, node, &typeName); err != nil {
			return nil, err
		}

		tn, err := typed.NewTypeName(typeName)
		if err != nil {
			return nil, err
		}

		return ValueDeclaration(tn), nil
	}
}
