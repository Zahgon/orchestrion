// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package lineinfo

import (
	"github.com/dave/dst"
)

// canonicalizationVisitor is a dst.Visitor that adds dst.NewLine before select AST nodes
// (*dst.GenDecl and select dst.Stmt nodes), to get the AST closer to what the canonical go format
// is. This allows us to have more accurate line information from the transformed AST and removes
// the risk of false equivalence when we assess the need for line directives.
type (
	canonicalizationVisitor struct {
		stack []*stackEntry
	}
	stackEntry struct {
		node      dst.Node
		lastChild dst.Node
	}
)

var _ dst.Visitor = (*canonicalizationVisitor)(nil)

func (v *canonicalizationVisitor) Visit(node dst.Node) dst.Visitor {
	_ = "STUB: not implemented"
	return *new(dst.Visitor)
}

// Upon returning, set the parent's last visited child to the current one...

// dave/dst will double-count new lines between imports if an ImportSpec with "After" spacing set
// to `dst.EmptyLine` is immediately followed by another with "Before" spacing of `dst.EmptyLine`.
// The two empty lines are  satisfied by the same; so we can safely turn either one into a
// `dst.NewLine` instead, so we get accurate line numbering data from the restorer.

// Don't space up the statements if they're not within a *dst.Block, or if the block contains exactly 1 statement.
