// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package code

import (
	"errors"

	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/dave/dst"
)

type (
	function interface {
		// Receiver returns the name of the receiver of this method. Fails if the current function is
		// not a method.
		Receiver() (string, error)
		// Name returns the name of this function, or an empty string if it is a function literal.
		Name() (string, error)

		// Argument returns the name of the argument at the given index in this function's type,
		// returningan error if the index is out of bounds.
		Argument(int) (string, error)
		// ArgumentOfType returns the name of the first argument in this function that has the provided
		// type, or an empty string if none is found.
		ArgumentOfType(string) (string, error)
		// ArgumentThatImplements returns the name of the first argument in this function that implements
		// the provided interface type, or an empty string if none is found.
		ArgumentThatImplements(string) (string, error)

		// Result returns the name of the return value at the given index in this function's type,
		// returning an error if the index is out of bounds.
		Result(int) (string, error)
		// ResultOfType returns the name of the first return value in this function that has the
		// provided type, or a empty string if none is found.
		ResultOfType(string) (string, error)
		// ResultThatImplements returns the name of the first return value in this function that implements
		// the provided interface type, or an empty string if none is found.
		ResultThatImplements(string) (string, error)
		// LastResultThatImplements returns the name of the last return value
		// (may not be the last item in the function's return list) in this function that implements
		// the provided interface type, or an empty string if none is found.
		LastResultThatImplements(string) (string, error)
		// FinalResultImplements returns whether the final (very last item that is returned) result implements the provided interface type.
		FinalResultImplements(string) (bool, error)
	}

	declaredFunc struct {
		signature
		Decl *dst.FuncDecl
	}

	literalFunc struct {
		signature
		Lit *dst.FuncLit
	}

	noFunc struct{}
)

var (
	errNoFunction = errors.New("no function is present in this node chain")
	errNotMethod  = errors.New("the function in this context is not a method")
)

func (d *dot) Function() function { _ = "STUB: not implemented"; return *new(function) }

func (f *declaredFunc) Receiver() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *declaredFunc) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (*literalFunc) Receiver() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (*literalFunc) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (noFunc) Receiver() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (noFunc) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (noFunc) Argument(int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (noFunc) ArgumentOfType(string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (noFunc) ArgumentThatImplements(string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (noFunc) Result(int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (noFunc) ResultOfType(string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (noFunc) ResultThatImplements(string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (noFunc) LastResultThatImplements(string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (noFunc) FinalResultImplements(string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type signature struct {
	context context.AdviceContext
	*dst.FuncType
}

func (s signature) Argument(index int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (s signature) ArgumentOfType(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s signature) ArgumentThatImplements(interfaceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s signature) Result(index int) (name string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s signature) ResultOfType(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s signature) ResultThatImplements(interfaceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s signature) LastResultThatImplements(name string) (string, error) {
	_ = "STUB: not implemented"
	// Return blank if there are no results.
	return "", nil
}

// Optimization: First, check for an exact match using TypeName parsing, finding the last one.

// Update last found index

// Increment index by the number of names in the field (or 1 if unnamed).

// If we found a match via TypeName, return it.

// If parsing failed or no match, fall through to type resolution.

// Resolve the interface type.

// Propagate error if interface resolution fails

// Fallback: Check using ExprImplements, iterating backward.
// Need field indices map again for this path.

// Found a match, return the corresponding field.

// Not found

func fieldAt(fields *dst.FieldList, index int, use string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Give a name to all items (if there are unnamed items, all items are unnamed).

// Give it a referenceable name if necessary.

// If the items were not anonymous, we can return immediately!

// Use <= to catch index being exactly the number of items

// If anonymous, we should have assigned the synthetic name earlier.
// If named, it should have been returned immediately.
// If we reach here and it was anonymous, we return the generated name.

// This path should ideally not be reached for named parameters if logic is correct.

func fieldOfType(fields *dst.FieldList, typeName string, use string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// No fields, no match!

// If the field is not named it's as if there is one.

// Not found!

// FinalResultImplements returns whether the final result implements the provided interface type.
func (s signature) FinalResultImplements(interfaceName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Optimization: First, check for an exact match using TypeName parsing.
// Note: Not using FindMatchingTypeName as we only need to check the last field.

// If parsing failed or no match, fall through to type resolution.

// Check if the last field type implements the interface.

// findImplementingField is a helper to find the first field in a list that matches
// an interface, either by exact type name or by implementation.
func findImplementingField(ctx context.AdviceContext, fields *dst.FieldList, interfaceName string, fieldKind string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// No fields, no match.

// 1. Check for exact type name match first.

// 2. If no exact name match, check for interface implementation.

// Invalid interface name, cannot proceed with implementation check.

// Iterate through fields to check for implementation.

// Increment index based on field names.

// 3. No match found.
