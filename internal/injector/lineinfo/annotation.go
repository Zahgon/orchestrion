// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package lineinfo

import (
	"go/ast"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

const generated = "<generated>"

type (
	// annotationVisitor is an ast.Visitor that adds `//line` directives to the visited nodes to
	// adjust their logical source location so it matches those of the original `*ast.File` tree.
	annotationVisitor struct {
		// dev is the decorator that transformed the original *ast.File into the visited *dst.File
		dec *decorator.Decorator
		// res is a restorer that was used to restore the visited *dst.File into an *ast.File, and hence
		// provides source location information for the restored AST.
		res *decorator.FileRestorer

		lineInfo
		stack []dst.Node
	}
	lineInfo struct {
		curFile string // The current un-adjusted file name
		adjFile string // The current adjusted file name
		curLine int    // The current un-adjusted line number
		adjLine int    // The current adjusted line number
	}
)

var _ ast.Visitor = (*annotationVisitor)(nil)

func (v *annotationVisitor) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

// If this is a [*dst.GenDecl] with a single item that's rendered without parentheses, we
// hoist decorations from the single item to the [*dst.GenDecl] itself, as it'll render
// better.

// Finished visiting a node, we don't have anything particular to do...

// Nodes such as [ast.FuncType] are not mapped directly by dst... They anyway do not represent
// lines that can show on stack frames, so it's not all that important...

// Emit a `//line <file>:1:1` directive at start of file to act as a base for all subsequent declarations.

// This is a virtual [*dst.Ident] node that was created by import management. It does not map
// back to a node in the original AST, and it's part of a [*dst.SelectorExpr] that we'll be
// able to properly map; so we can safely ignore it now.

// This is a synthetic node...

// Update the adjusted position to the current observed one...

// Current & adjusted positions match, and we've not changed adjusted files -- nothing to do!

// We're already correctly adjusted, so we don't need to add another directive...

func (l *lineInfo) directive(fileStart bool) string { _ = "STUB: not implemented"; return "" }

// We don't emit 0 line numbers

// We emit :1:1 for line 1, so we match the output of `go tool cover` for this particular case.

// Otherwise, we only emit line number (no column).
