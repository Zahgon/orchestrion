// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package join

import (
	gocontext "context"
	"fmt"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/aspect/may"
	"github.com/DataDog/orchestrion/internal/injector/typed"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/dave/dst"
	"github.com/goccy/go-yaml/ast"
)

type structDefinition struct {
	TypeName typed.TypeName
}

// StructDefinition matches the definition of a particular struct given its fully qualified name.
func StructDefinition(typeName typed.TypeName) *structDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (s *structDefinition) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (s *structDefinition) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (*structDefinition) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (s *structDefinition) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false

	// We can't ever match a pointer definition
}

func (s *structDefinition) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type (
	StructLiteralMatch int
	structLiteral      struct {
		TypeName typed.TypeName
		Field    string
		Match    StructLiteralMatch
	}
)

const (
	// StructLiteralMatchAny matches struct literals regardless of whether they are pointer or value.
	// [StructLiteral] join points specified with this match type may match [*dst.CompositeLit] or
	// [*dst.UnaryExpr] nodes.
	StructLiteralMatchAny StructLiteralMatch = iota
	// StructLiteralMatchValueOnly matches struct literals that are not pointers. [StructLiteral] join
	// points specified with this match type only ever match [*dst.CompositeLit] nodes.
	StructLiteralMatchValueOnly
	// StructLiteralMatchPointerOnly matches struct literals that are pointers. [StructLiteral] join
	// points specified with this match type only ever match [*dst.UnaryExpr] nodes.
	StructLiteralMatchPointerOnly
)

// StructLiteralField matches a specific field in struct literals of the designated type.
func StructLiteralField(typeName typed.TypeName, field string) *structLiteral {
	_ = "STUB: not implemented"
	return nil
}

// StructLiteral matches struct literal expressions of the designated type, filtered by the
// specified match type.
func StructLiteral(typeName typed.TypeName, match StructLiteralMatch) *structLiteral {
	_ = "STUB: not implemented"
	return nil
}

func (s *structLiteral) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (s *structLiteral) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (*structLiteral) FileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (s *structLiteral) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

// match only if the current node is equal to & and the underlying node matches
// the struct literal we are looking for

// do not match if the parent is equal to &

func (s *structLiteral) matchesLiteral(node dst.Node) bool { _ = "STUB: not implemented"; return false }

func (s *structLiteral) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["struct-definition"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var spec string
		if err := yaml.NodeToValueContext(ctx, node, &spec); err != nil {
			return nil, err
		}

		tn, err := typed.NewTypeName(spec)
		if err != nil {
			return nil, err
		}
		if tn.Pointer {
			return nil, fmt.Errorf("struct-definition type must not be a pointer (got %q)", spec)
		}

		return StructDefinition(tn), nil
	}
	unmarshalers["struct-literal"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var spec struct {
			Type  string
			Field string
			Match StructLiteralMatch
		}
		if err := yaml.NodeToValueContext(ctx, node, &spec); err != nil {
			return nil, err
		}

		tn, err := typed.NewTypeName(spec.Type)
		if err != nil {
			return nil, err
		}

		if spec.Field != "" {
			if spec.Match != StructLiteralMatchAny {
				return nil, fmt.Errorf("struct-literal.field is not allowed with struct-literal.match: %s", spec.Match)
			}
			return StructLiteralField(tn, spec.Field), nil
		}

		return StructLiteral(tn, spec.Match), nil
	}
}

var _ yaml.NodeUnmarshalerContext = (*StructLiteralMatch)(nil)

func (s *StructLiteralMatch) UnmarshalYAML(ctx gocontext.Context, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (s StructLiteralMatch) String() string { _ = "STUB: not implemented"; return "" }

func (s StructLiteralMatch) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}
