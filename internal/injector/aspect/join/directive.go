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

type directive string

// Directive matches nodes that are prefaced by a special pragma comment, which
// is a single-line style comment without any blanks between the leading // and
// the directive name. Directives apply to the node they are directly attached
// to, but also to certain nested nodes:
//   - For assignments, it applies to the RHS only; unless it's a declaration
//     assignment (the := token), in which case it also applies to the LHS,
//   - For call expressions, it applies only to the function part (not the
//     arguments)n
//   - For channel send operations, it only applies to the value being sent,
//   - For defer, go, and return statements, it applies to the value side.
func Directive(name string) directive { _ = "STUB: not implemented"; return *new(directive) }

func (directive) PackageMayMatch(_ *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (d directive) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (d directive) Matches(ctx context.AspectContext) bool { _ = "STUB: not implemented"; return false }

func (d directive) matchesChain(chain *context.NodeChain) bool {
	_ = "STUB: not implemented"
	return false
}

// If this is a spec (variable, const, import), so we need to also check the parent for directives!

// Also check whether the parent carries the directive if it's one of the node types that would
// typically carry directives that applies to its nested node.

// For assignments, the directive only applies downwards to the RHS, unless it's a declaration,
// then it also applies to any declared identifier.

// For call expressions, the directive only applies to the called function, not its type
// signature or arguments list.

// For channel send statements, the directive only applies to the value being sent, not to the
// receiving channel.

// Defer statements, go statements, and return statements all forward the directive to the
// value(s); and expression statements are just wrappers of expressions, so naturally directives
// that apply to the statement also apply to the expression.

func (d directive) matches(dec string) bool {
	_ = "STUB: not implemented"
	// Trim leading white space
	return false
}

// Check we have a single-line comment

// Check the // is followed immediately by the directive name

// If there is something after the directive name, it must be white space

func (directive) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (d directive) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["directive"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var name string
		if err := yaml.NodeToValueContext(ctx, node, &name); err != nil {
			return nil, err
		}
		return Directive(name), nil
	}
}
