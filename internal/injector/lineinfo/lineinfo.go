// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package lineinfo

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

// AnnotateMovedNodes adds `//line` directives to the provided `*dst.File` to adjust source location
// information of each AST node that exists in the original source file to its location there, and
// marks other nodes as originating from `<generated>`.
func AnnotateMovedNodes(
	// The decorator that produced the *dst.File
	decorator *decorator.Decorator,
	// The *dst.File to annotate
	file *dst.File,
	// A function that creates a new *decorator.FileRestorer for the given filename
	newRestorer func(string) *decorator.FileRestorer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Pre-process the AST to make it closer to the canonical go format, which will allow us to have
// more accurate "after-printing" line information.

// Restore to an *ast.File so we can obtain the new line information data.

// Visit the AST to add `//line` directives where the updated line information no longer matches
// the original source file's.
