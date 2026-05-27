// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package advice

import (
	gocontext "context"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/advice/code"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type assignValue struct {
	Template *code.Template
}

func AssignValue(template *code.Template) *assignValue { _ = "STUB: not implemented"; return nil }

func (a *assignValue) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a *assignValue) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func (a *assignValue) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["assign-value"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		var template *code.Template
		if err := yaml.NodeToValueContext(ctx, node, &template); err != nil {
			return nil, err
		}
		return AssignValue(template), nil
	}
}
