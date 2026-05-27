// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package typed

import (
	"regexp"

	"github.com/dave/dst"

	"github.com/DataDog/orchestrion/internal/fingerprint"
)

// Common built-in type definitions for convenience.
// These pre-defined TypeName instances help avoid repeated string literals
// and potential typos when referring to common Go built-in types.
var (
	// Basic types currently used in the codebase
	Any    = MustTypeName("any")
	Bool   = MustTypeName("bool")
	String = MustTypeName("string")
	// Uncomment these when we used.
	// Byte   = MustTypeName("byte")
	// Int    = MustTypeName("int")
	// Error  = MustTypeName("error")
)

// TypeName represents a parsed Go type name, potentially including a package path and pointer indicator.
type TypeName struct {
	// ImportPath is the import Path that provides the type, or an empty string if the
	// type is local or built-in (like "error" or "any").
	ImportPath string
	// Name is the leaf (un-qualified) name of the type.
	Name string
	// Pointer determines whether the specified type is a pointer or not.
	Pointer bool
}

// FIXME: this does not support all the type syntax, like: "chan Event"
// It primarily handles identifiers, qualified identifiers, and pointers to those.
var typeNameRe = regexp.MustCompile(`\A(\*)?\s*(?:([A-Za-z0-9_.-]+(?:/[A-Za-z0-9_.-]+)*)\.)?([A-Za-z_][A-Za-z0-9_]*)\z`)

// NewTypeName parses a string representation of a type name into a TypeName struct.
// It returns an error if the syntax is invalid according to its limited regular expression.
func NewTypeName(n string) (tn TypeName, err error) {
	_ = "STUB: not implemented"
	return *new(TypeName), nil
}

// MustTypeName is the same as NewTypeName, except it panics in case of an error.
func MustTypeName(n string) (tn TypeName) { _ = "STUB: not implemented"; return *new(TypeName) }

// Matches determines whether the provided AST expression node represents the same type
// as this TypeName. This performs a structural comparison based on the limited types
// supported by the parsing regex (identifiers, selectors, pointers, empty interface).
func (n TypeName) Matches(node dst.Expr) bool { _ = "STUB: not implemented"; return false }

// Handle generic types with single type parameter (e.g., MyType[T])

// Handle generic types with multiple type parameters (e.g., MyType[T, U])

// We only match the empty interface (as "any")

// MatchesDefinition determines whether the provided node matches the definition
// of this TypeName. The `importPath` argument determines the context in which
// the assertion is made.
func (n TypeName) MatchesDefinition(node dst.Expr, importPath string) bool {
	_ = "STUB: not implemented"
	return false
}

// AsNode converts the TypeName back into a dst.Expr AST node.
// Useful for generating code that refers to this type.
func (n *TypeName) AsNode() dst.Expr { _ = "STUB: not implemented"; return *new(dst.Expr) }

// Hash contributes the TypeName's properties to a fingerprint hasher.
func (n TypeName) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

// FindMatchingTypeName parses a type name string and searches a field list for the first field whose type matches.
// It returns the index of the matching field and whether a match was found.
// The index accounts for fields with multiple names.
func FindMatchingTypeName(fields *dst.FieldList, typeNameStr string) (index int, found bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// If the type name string is invalid, we can't match it.

// Found a match.

// Increment index by the number of names in the field (or 1 if unnamed).

// No match found
