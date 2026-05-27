// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package typed

import (
	"go/ast"
	"go/types"

	"github.com/dave/dst"
)

type (
	// ReferenceKind denotes the style of a reference, which influences compilation and linking requirements.
	ReferenceKind bool

	// ReferenceMap associates import paths to ReferenceKind values.
	ReferenceMap struct {
		refs    map[string]ReferenceKind
		aliases map[string]string
		nodeMap map[dst.Node]ast.Node
		scopes  map[ast.Node]*types.Scope
	}
)

const (
	// ImportStatement references must be made available to the compiler via the provided `importcfg`.
	ImportStatement ReferenceKind = true
	// RelocationTarget references must be made available to the linker, and must be referenced (directly or not) by the main package.
	RelocationTarget ReferenceKind = false
)

func NewReferenceMap(nodeMap map[dst.Node]ast.Node, scopes map[ast.Node]*types.Scope) ReferenceMap {
	_ = "STUB: not implemented"
	return *new(ReferenceMap)
}

// AddImport takes a package import path and the name in file and the result of a recursive parent lookup.
// It first determines if the import is already present
// and if it has not been shadowed by a local declaration. If both conditions are met, the import is added to the
// reference map and the function returns true. Otherwise, it returns false.
func (r *ReferenceMap) AddImport(file *dst.File, nodes []dst.Node, path string, localName string) bool {
	_ = "STUB: not implemented"
	return false
}

// If the import is already present, has a meaningful alias or no alias,
// and is accessible from the current scope, we don't need to do anything.

// Register in this ReferenceMap

// We don't register blank aliases, as this is the default behavior anyway...

// isImportInScope checks if the provided name is an import in the scope of the provided node
func (r *ReferenceMap) isImportInScope(nodes []dst.Node, path string, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// Somehow scopes are not attached to FuncDecl nodes, so we need to look at the type ¯\_(シ)_/¯

// hasImport checks if the provided file already imports the provided path and its local name.
func hasImport(file *dst.File, path string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// AddLink registers the provided path as a relocation target resolution source. If this path is
// already registered as an import, this method does nothing and returns false.
func (r *ReferenceMap) AddLink(file *dst.File, path string) bool {
	_ = "STUB: not implemented"
	return false
}

// AddSyntheticImports adds the registered imports to the provided *dst.File. This is not safe to
// call during an AST traversal by dstutil.Apply, as this may offset the declaration list by 1 in
// case a new import declaration needs to be added, which would result in re-traversing current
// declaration when the cursor moves forward. Instead, it is advise to call this method after
// dstutil.Apply has returned.
func (r *ReferenceMap) AddSyntheticImports(file *dst.File) bool {
	_ = "STUB: not implemented"
	return false
}

// Sort the import specs to ensure deterministic output.

// Find the last import declaration in the file...

// ...or create a new one if none is found...

// Add the necessary imports

func (r *ReferenceMap) Merge(other ReferenceMap) { _ = "STUB: not implemented"; return }

func (r *ReferenceMap) Map() map[string]ReferenceKind { _ = "STUB: not implemented"; return nil }

func (r *ReferenceMap) Count() int { _ = "STUB: not implemented"; return 0 }

func (r *ReferenceMap) add(path string, kind ReferenceKind) bool {
	_ = "STUB: not implemented"
	return false
}

// If it was already in as an ImportStatement, we don't do anything, since that is the strongest
// kind of reference (imported implies relocatable, the reverse is not true).

func (k ReferenceKind) String() string { _ = "STUB: not implemented"; return "" }
