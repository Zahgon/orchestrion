// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package injector

import (
	"context"

	"github.com/DataDog/orchestrion/internal/injector/aspect"
	"github.com/dave/dst"
)

// packageFilterAspects filters out aspects that imply imports not present in the import map and return a copy
// of the aspect array.
func (i *Injector) packageFilterAspects(aspects []*aspect.Aspect) []*aspect.Aspect {
	_ = "STUB: not implemented"
	return nil
}

// canonicalizeImports works around the issue detailed in https://github.com/dave/dst/issues/45
// where dave/dst improperly handles multiple imports of the same package with different aliases,
// resulting in invalid output source code.
//
// To do so, it modifies the AST file so that it only includes a single import per path, using the
// first non-empty alias found.
func canonicalizeImports(ctx context.Context, file *dst.File) { _ = "STUB: not implemented"; return }

// Re-use the backing store, we'll keep <= what was there.

func importSpecsByImportPath(ctx context.Context, imports []*dst.ImportSpec) map[string][]*dst.ImportSpec {
	_ = "STUB: not implemented"
	return nil
}

func filterExtraneousImports(byPath map[string][]*dst.ImportSpec) map[*dst.ImportSpec]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// Preference order is: spec.Name == nil < spec.Name.Name == "_" < spec.Name.Name != "_"

// We found a non-empty alias, no need to look further.

func filterDecls(file *dst.File, retain map[*dst.ImportSpec]struct{}) {
	_ = "STUB: not implemented"
	return
}

// Only visit the children of `import` declarations.

// Filter out ImportSpec entries to keep only those in retain

// No need to traverse children.

// No need to visit any other kind of declaration

// Visit other node types (e.g, the *ast.File)

// Imports are before any other kind of declaration, we can abort traversal as soon as we
// find a declaration that is not an `import` declaration.

// Proceed with the rest of the nodes (there may be more imports).

// Imports are before any other kind of declaration, we can abort traversal as soon as we
// find a declaration that is not an `import` declaration.

// Proceed with the rest of the nodes (there may be imports down there).
