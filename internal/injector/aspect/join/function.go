// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package join

import (
	gocontext "context"
	"go/types"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/aspect/may"
	"github.com/DataDog/orchestrion/internal/injector/typed"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/dave/dst"
	"github.com/goccy/go-yaml/ast"
)

type (
	// typeResolver defines the capability to resolve a dst expression to its go/types type.
	typeResolver interface {
		ResolveType(dst.Expr) types.Type
	}

	functionInformation struct {
		Receiver   dst.Expr      // The receiver if this is a method declaration
		Type       *dst.FuncType // The function's type signature
		ImportPath string        // The import path of the package containing the function
		Name       string        // The name of the function (blank for function literal expressions)

		typeResolver typeResolver // The type resolver to use for type checking
	}

	FunctionOption interface {
		fingerprint.Hashable

		impliesImported() []string

		packageMayMatch(ctx *may.PackageContext) may.MatchType
		fileMayMatch(ctx *may.FileContext) may.MatchType

		evaluate(functionInformation) bool
	}

	functionDeclaration struct {
		Options []FunctionOption
	}
)

// Function matches function declaration nodes based on properties of
// their signature.
func Function(opts ...FunctionOption) *functionDeclaration { _ = "STUB: not implemented"; return nil }

func (s *functionDeclaration) ImpliesImported() (list []string) {
	_ = "STUB: not implemented"
	return nil
}

func (s *functionDeclaration) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (s *functionDeclaration) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (s *functionDeclaration) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *functionDeclaration) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}

type functionName string

func Name(name string) FunctionOption { _ = "STUB: not implemented"; return *new(FunctionOption) }

func (functionName) impliesImported() []string { _ = "STUB: not implemented"; return nil }

func (functionName) packageMayMatch(_ *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (fo functionName) fileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (fo functionName) evaluate(info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

func (fo functionName) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type signature struct {
	Arguments []typed.TypeName
	Results   []typed.TypeName
}

// Signature matches function declarations based on their arguments and return
// value types.
func Signature(args []typed.TypeName, ret []typed.TypeName) FunctionOption {
	_ = "STUB: not implemented"
	return *new(FunctionOption)
}

func (fo *signature) packageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (*signature) fileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (fo *signature) impliesImported() (list []string) { _ = "STUB: not implemented"; return nil }

func (fo *signature) evaluate(info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

func (fo *signature) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type signatureContains struct {
	signature
}

// SignatureContains matches function declarations based on their arguments and
// return value types in any order and does not require all arguments or return values to be present.
func SignatureContains(args []typed.TypeName, ret []typed.TypeName) FunctionOption {
	_ = "STUB: not implemented"
	return *new(FunctionOption)
}

func (fo *signatureContains) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}

func (fo *signatureContains) evaluate(info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

// containsAnyType checks if any of the expected types match any of the actual types in the field list.
// Returns false if either slice is empty or nil.
func containsAnyType(expectedTypes []typed.TypeName, fieldList *dst.FieldList) bool {
	_ = "STUB: not implemented"
	// Quick return if either side is empty.
	return false
}

// Check if any expected type matches any actual type.

type receiver struct {
	TypeName typed.TypeName
}

func Receiver(typeName typed.TypeName) FunctionOption {
	_ = "STUB: not implemented"
	return *new(FunctionOption)
}

func (fo *receiver) packageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (fo *receiver) fileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (fo *receiver) evaluate(info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

func (fo *receiver) impliesImported() []string { _ = "STUB: not implemented"; return nil }

func (fo *receiver) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type functionBody struct {
	Function Point
}

// FunctionBody returns the *dst.BlockStmt of the matched *dst.FuncDecl body.
func FunctionBody(up Point) *functionBody { _ = "STUB: not implemented"; return nil }

func (s *functionBody) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (s *functionBody) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (s *functionBody) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (s *functionBody) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *functionBody) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

// resultImplements matches functions where at least one return value's type
// implements the specified interface.
type resultImplements struct {
	InterfaceName string
}

// ResultImplements creates a FunctionOption that matches functions where at least one
// return value implements the named interface.
func ResultImplements(interfaceName string) FunctionOption {
	_ = "STUB: not implemented"
	return *new(FunctionOption)
}

func (*resultImplements) impliesImported() []string {
	_ = "STUB: not implemented"
	// A type can implement an interface without importing the interface's package
	// due to Go's structural typing system.
	return nil
}

func (_ *resultImplements) packageMayMatch(_ *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	// Cannot reliably determine possibility of match based on package imports
	// due to structural typing. A type can implement an interface without
	// importing the interface's package.
	return *new(may.MatchType)
}

func (_ *resultImplements) fileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	// Cannot reliably determine possibility of match based on file contents
	// due to structural typing and type aliases.
	return *new(may.MatchType)
}

// evaluateFieldListImplements checks if any field in the list matches the interfaceName,
// either by exact type name or by interface implementation.
func evaluateFieldListImplements(fields *dst.FieldList, interfaceName string, info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

// Optimization: First, check for an exact match using the helper.

// Found direct match

// If no exact match, check implementation (requires type resolver).

// Cannot check implementation without resolver.

// Invalid interface name.

// Found an implementing type.

// No match found.

func (fo *resultImplements) evaluate(info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

func (fo *resultImplements) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}

// finalResultImplements matches functions where specifically the final return value
// implements the specified interface.
type finalResultImplements struct {
	InterfaceName string
}

// FinalResultImplements creates a FunctionOption that matches functions where the final
// return value implements the named interface.
func FinalResultImplements(interfaceName string) FunctionOption {
	_ = "STUB: not implemented"
	return *new(FunctionOption)
}

func (*finalResultImplements) impliesImported() []string {
	_ = "STUB: not implemented"
	// A type can implement an interface without importing the interface's package
	// due to Go's structural typing system.
	return nil
}

func (_ *finalResultImplements) packageMayMatch(_ *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	// Cannot reliably determine possibility of match based on package imports
	// due to structural typing. A type can implement an interface without
	// importing the interface's package.
	return *new(may.MatchType)
}

func (_ *finalResultImplements) fileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	// Cannot reliably determine possibility of match based on file contents
	// due to structural typing and type aliases.
	return *new(may.MatchType)
}

func (fo *finalResultImplements) evaluate(info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

// No return values, no match.

// Optimization: First, check for an exact match using TypeName parsing.

// Found direct match

// If parsing failed or no match, fall through to type resolution.

// Ensure the type resolver is available.

// Resolve the target interface name (e.g., "io.Reader", "error") to a types.Interface.

// If the interface name is invalid or cannot be resolved, we cannot match.

// Check if the last field implements the interface.

func (fo *finalResultImplements) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}

// argumentImplements matches functions where at least one argument's type
// implements the specified interface.
type argumentImplements struct {
	InterfaceName string
}

// ArgumentImplements creates a FunctionOption that matches functions where at least one
// argument implements the named interface.
func ArgumentImplements(interfaceName string) FunctionOption {
	_ = "STUB: not implemented"
	return *new(FunctionOption)
}

func (fo *argumentImplements) impliesImported() []string { _ = "STUB: not implemented"; return nil }

func (_ *argumentImplements) packageMayMatch(_ *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	// Cannot reliably determine possibility of match based on package imports
	// due to structural typing. A type can implement an interface without
	// importing the interface's package.
	return *new(may.MatchType)
}

func (_ *argumentImplements) fileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	// Cannot reliably determine possibility of match based on file contents
	// due to structural typing and type aliases.
	return *new(may.MatchType)
}

func (fo *argumentImplements) evaluate(info functionInformation) bool {
	_ = "STUB: not implemented"
	return false
}

func (fo *argumentImplements) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	unmarshalers["function-body"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		up, err := FromYAML(ctx, node)
		if err != nil {
			return nil, err
		}
		return FunctionBody(up), nil
	}

	unmarshalers["function"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var unmarshalOpts []unmarshalFuncDeclOption
		if err := yaml.NodeToValueContext(ctx, node, &unmarshalOpts); err != nil {
			return nil, err
		}
		opts := make([]FunctionOption, len(unmarshalOpts))
		for i, opt := range unmarshalOpts {
			opts[i] = opt.FunctionOption
		}
		return Function(opts...), nil
	}
}

type unmarshalFuncDeclOption struct {
	FunctionOption
}

func (o *unmarshalFuncDeclOption) UnmarshalYAML(ctx gocontext.Context, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: Validation happens later during type resolution.

// NOTE: Validation happens later during type resolution.
