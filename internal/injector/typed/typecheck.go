// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package typed

import (
	"go/types"

	"github.com/dave/dst"
)

// TypeResolver defines the capability to resolve a dst expression to its go/types type.
type TypeResolver interface {
	ResolveType(dst.Expr) types.Type
}

// ExprImplements checks if the type of a dst.Expr, resolved using the provider,
// implements the given interface.
func ExprImplements(resolver TypeResolver, expr dst.Expr, iface *types.Interface) bool {
	_ = "STUB: not implemented"
	return false
}

// typeImplements checks if a type implements an interface.
func typeImplements(t types.Type, iface *types.Interface) bool {
	_ = "STUB: not implemented"
	return false
}

// Direct implementation check.

// Fallback: check by method name only. types.Implements can fail when
// ResolveInterfaceTypeByName (importer.Default) and the type checker
// (importer.ForCompiler) produce different *types.Package objects for the
// same path, causing types.Identical to reject named types in method
// signatures (e.g., time.Time in context.Context.Deadline).
// types.LookupFieldOrMethod compares by package path, not pointer, so it
// correctly finds matching methods across importers.

// ResolveInterfaceTypeByName takes an interface name as a string and resolves it to an interface type.
func ResolveInterfaceTypeByName(name string) (*types.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle built-in types or unqualified names.

// Not found in universe scope.

// Found in universe, now validate it's an interface type name.

// Handle package-qualified types (e.g., "io.Writer").

// Specific error for import failure.

// Not found within the imported package's scope.

// Found in package scope, now validate it's an interface type name.

// SplitPackageAndName splits a fully qualified type name like "io.Reader" or "example.com/pkg.Type"
// into its package path and local name.
// Returns ("", "error") for built-in "error".
// Returns ("", "MyType") for unqualified "MyType".
// For generic types like "iter.Seq[T]" or "iter.Seq[io.Reader]", the type parameters
// are included in the local name, so it returns ("iter", "Seq[T]") and ("iter", "Seq[io.Reader]").
func SplitPackageAndName(fullName string) (pkgPath string, localName string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Assume built-in type (like "error") or unqualified local type.

// Find the position of the first '[' which indicates generic type parameters

// Determine the substring to search for the last dot

// Only search for dots before the generic type parameters

// Find the last dot in the search string

// No dot found (shouldn't happen given the initial check, but be safe)

// validateTypeNameIsInterface checks if a successfully looked-up types.Object represents
// a type name that resolves to an interface. It assumes obj is not nil.
func validateTypeNameIsInterface(obj types.Object, fullName string, pkgPath string, typeName string) (*types.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Provide context whether it was expected to be built-in or package-qualified.

// Use the original full name in the error message for clarity.

// Since types.IsInterface passed, we can safely cast typ.Underlying() to *types.Interface
