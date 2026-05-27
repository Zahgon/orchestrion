// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package context

import (
	gocontext "context"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

// GoLangVersion represents a go language level. It's a string of the form "go1.18".
type GoLangVersion struct {
	label string
}

func ParseGoLangVersion(lang string) (GoLangVersion, error) {
	_ = "STUB: not implemented"
	return *new(GoLangVersion), nil
}

func MustParseGoLangVersion(lang string) GoLangVersion {
	_ = "STUB: not implemented"
	return *new(GoLangVersion)
}

func (g GoLangVersion) String() string {
	_ = "STUB: not implemented"

	// IsAny returns true if the GoLang version selection is blank, meaning no particular constraint is
	// imposed on language level.
	return ""
}

func (g GoLangVersion) IsAny() bool { _ = "STUB: not implemented"; return false }

func (g *GoLangVersion) SetAtLeast(other GoLangVersion) { _ = "STUB: not implemented"; return }

func Compare(left GoLangVersion, right GoLangVersion) int { _ = "STUB: not implemented"; return 0 }

var _ fingerprint.Hashable = (*GoLangVersion)(nil)

func (g GoLangVersion) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

var _ yaml.NodeUnmarshalerContext = (*GoLangVersion)(nil)

func (g *GoLangVersion) UnmarshalYAML(ctx gocontext.Context, node ast.Node) error {
	_ = "STUB: not implemented"
	return nil
}
