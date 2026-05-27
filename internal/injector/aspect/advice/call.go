// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package advice

import (
	gocontext "context"
	"regexp"
	"strings"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/advice/code"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/typed"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type appendArgs struct {
	TypeName  typed.TypeName
	Templates []*code.Template
}

// AppendArgs appends arguments of a given type to the end of a function call. All arguments must be
// of the same type, as they may be appended at the tail end of a variadic call.
func AppendArgs(typeName typed.TypeName, templates ...*code.Template) *appendArgs {
	_ = "STUB: not implemented"
	return nil
}

func (a *appendArgs) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// The function call has an ellipsis, so we need to append our new arguments to the last argument,
// which is a slice. To do so, we need to provision a new slice of the right type and size, append
// all the relevant data in there, and then replace the last argument with the new slice.

func (a *appendArgs) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func (a *appendArgs) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type redirectCall struct {
	ImportPath string
	Name       string
}

// ReplaceFunction replaces the called function with the provided drop-in replacement. The signature
// must be compatible with the original function (it may accept a new variadic argument).
func ReplaceFunction(path string, name string) *redirectCall { _ = "STUB: not implemented"; return nil }

func (r *redirectCall) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Just in case

func (r *redirectCall) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func (r *redirectCall) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["append-args"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		var args struct {
			TypeName string           `yaml:"type"`
			Values   []*code.Template `yaml:"values"`
		}

		if err := yaml.NodeToValueContext(ctx, node, &args); err != nil {
			return nil, err
		}

		tn, err := typed.NewTypeName(args.TypeName)
		if err != nil {
			return nil, err
		}

		return AppendArgs(tn, args.Values...), nil
	}
	unmarshalers["replace-function"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		var (
			fqn  string
			path string
			name string
		)
		if err := yaml.NodeToValueContext(ctx, node, &fqn); err != nil {
			return nil, err
		}

		if idx := strings.LastIndex(fqn, "."); idx >= 0 {
			path = fqn[:idx]
			name = fqn[idx+1:]
		} else {
			name = fqn // Built-in function, function from the same package
		}

		return ReplaceFunction(path, name), nil
	}
}

var (
	importPathRe        = regexp.MustCompile(`^(?:.+/)?([^/]+?)(?:\.v\d+|/v\d+)?$`)
	notValidIdentCharRe = regexp.MustCompile(`[^a-zA-Z0-9_]`)
)

// inferPkgName extracts the last part of an import path and sanitizes it to be a valid Go
// identifier continuation (meaning it assumes it would be appended to a valid Go identifier). This
// is done for cosmetic purposes and does not require being accurate.
func inferPkgName(importPath string) string { _ = "STUB: not implemented"; return "" }
