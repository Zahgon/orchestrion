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

type wrapExpression struct {
	Template *code.Template
}

func WrapExpression(template *code.Template) *wrapExpression { _ = "STUB: not implemented"; return nil }

func (a *wrapExpression) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a *wrapExpression) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func (a *wrapExpression) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["wrap-expression"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		var template code.Template
		if err := yaml.NodeToValueContext(ctx, node, &template); err != nil {
			return nil, err
		}
		return WrapExpression(&template), nil
	}
}
