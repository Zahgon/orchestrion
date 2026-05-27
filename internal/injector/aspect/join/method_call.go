// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package join

import (
	gocontext "context"
	"errors"
	"fmt"
	"go/types"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/aspect/may"
	"github.com/DataDog/orchestrion/internal/injector/typed"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type (
	MethodCallMatch int

	methodCall struct {
		Receiver typed.TypeName
		Name     string
		Match    MethodCallMatch
	}
)

const (
	// MethodCallMatchAny matches calls regardless of whether the receiver is a pointer or value. This is the default.
	MethodCallMatchAny MethodCallMatch = iota
	// MethodCallMatchPointerOnly matches only calls where the receiver is a pointer type.
	MethodCallMatchPointerOnly
	// MethodCallMatchValueOnly matches only calls where the receiver is a value type.
	MethodCallMatchValueOnly
)

func MethodCall(receiver typed.TypeName, name string, match MethodCallMatch) *methodCall {
	_ = "STUB: not implemented"
	return nil
}

func (m *methodCall) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (m *methodCall) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (m *methodCall) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (m *methodCall) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *methodCall) matchesType(t types.Type) bool { _ = "STUB: not implemented"; return false }

// MethodCallMatchAny

func (m *methodCall) matchesNamed(t types.Type) bool { _ = "STUB: not implemented"; return false }

func (m *methodCall) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["method-call"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var spec struct {
			Receiver string          `yaml:"receiver"`
			Name     string          `yaml:"name"`
			Match    MethodCallMatch `yaml:"match"`
		}
		if err := yaml.NodeToValueContext(ctx, node, &spec); err != nil {
			return nil, err
		}

		if spec.Receiver == "" {
			return nil, errors.New("method-call: missing required field 'receiver'")
		}
		if spec.Name == "" {
			return nil, errors.New("method-call: missing required field 'name'")
		}

		tn, err := typed.NewTypeName(spec.Receiver)
		if err != nil {
			return nil, fmt.Errorf("method-call: invalid receiver type %q: %w", spec.Receiver, err)
		}
		if tn.Pointer {
			return nil, fmt.Errorf("method-call: receiver type must not include a pointer sigil (use match: pointer-only instead): %q", spec.Receiver)
		}

		return MethodCall(tn, spec.Name, spec.Match), nil
	}
}

var _ yaml.NodeUnmarshalerContext = (*MethodCallMatch)(nil)

func (m *MethodCallMatch) UnmarshalYAML(ctx gocontext.Context, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (m MethodCallMatch) String() string { _ = "STUB: not implemented"; return "" }

func (m MethodCallMatch) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }
