// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package code

import (
	"github.com/dave/dst"

	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
)

// dot provides the `.` value to code templates, and is used to access various bits of
// information from the template's rendering context.
type (
	placeholders struct {
		singletons map[dst.Node]string
		byName     map[string]dst.Node
	}

	dot struct {
		context      context.AdviceContext // The node in context of which the template is rendered
		placeholders                       // Placeholders used by the template
	}
)

func (d *dot) String() string { _ = "STUB: not implemented"; return "" }

// forNode obtains the placeholder syntax to use for referencing the given node. If singleton is
// true, this returns the same placeholder for each invocation with the same node argument.
// Otherwise, this returns a new placeholder for each invocation, guaranteeing that different AST
// nodes are produced (it's an error to have the same AST node multiple times in the output AST).
func (p *placeholders) forNode(node dst.Node, singleton bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Will be filled in later once we have determined the name

// replaceAllIn replaces all placeholders found in the given AST with the actual dst.Expr value.
func (p *placeholders) replaceAllIn(ast dst.Node) dst.Node {
	_ = "STUB: not implemented"
	return *new(dst.Node)
}
