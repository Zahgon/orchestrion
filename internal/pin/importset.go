// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package pin

import (
	"github.com/dave/dst"
)

// importSet is a set of imported packages from a given Go source file. It can be
// used to inspect whether a particular import path is imported or not, and to
// add new imports to the file.
type importSet struct {
	file     *dst.File
	imports  *dst.GenDecl
	imported map[string]*dst.ImportSpec
}

// importSetFrom creates a new [importSet] from the provided [*dst.File].
func importSetFrom(file *dst.File) *importSet { _ = "STUB: not implemented"; return nil }

// This should never happen as go/parser already verified this.

// This should never happen as go/parser already verified this.

// Add registers a new import into the receiver, and returns the resulting
// import spec AST node. If the import was already present, the function returns
// the currently registered import spec and false.
func (s *importSet) Add(path string) (*dst.ImportSpec, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Except returns the list of all import paths present in the receiver, except
// those present in the `omit` list.
func (s *importSet) Except(omit ...string) []string { _ = "STUB: not implemented"; return nil }

// Find returns the import spec for the provided import path if it's present in
// the receiver. If the import path is not imported, `nil` is returned instead.
func (s *importSet) Find(path string) *dst.ImportSpec { _ = "STUB: not implemented"; return nil }

// Remove removes the provided import path from the receiver, and returns true
// if any change to the AST was made.
func (s *importSet) Remove(toRemove string) bool {
	_ = "STUB: not implemented"

	// Remove actual import declarations from the AST.
	return false
}

// Remove imports from the file-level registry.

// Finally, remove from the quick-access set.

// firstImportIn returns the first import declaration found in the provided
// [*dst.File]. If no import declaration is present, returns `nil`.
func firstImportIn(file *dst.File) *dst.GenDecl { _ = "STUB: not implemented"; return nil }

// newImportDeclIn introduces a new import declaration at the top of the
// provided [*dst.File], and returns the newly created declaration.
func newImportDeclIn(file *dst.File) *dst.GenDecl { _ = "STUB: not implemented"; return nil }

// If there's no imports array, pre-allocate one with the default capacity.
