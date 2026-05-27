// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Package advice provides implementations of the injector.Action interface for
// common AST changes.
package advice

import (
	gocontext "context"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type addBlankImport string

func AddBlankImport(path string) addBlankImport {
	_ = "STUB: not implemented"
	return *new(addBlankImport)
}

func (a addBlankImport) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a addBlankImport) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func (a addBlankImport) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["add-blank-import"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		var path string
		if err := yaml.NodeToValueContext(ctx, node, &path); err != nil {
			return nil, err
		}
		return AddBlankImport(path), nil
	}
}
