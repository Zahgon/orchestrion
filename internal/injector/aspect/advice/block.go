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

type prependStatements struct {
	Template *code.Template

	order     int
	namespace string
}

// PrependStmts prepends statements to the matched *dst.BlockStmt. This action
// can only be used if the selector matches on a *dst.BlockStmt. The prepended
// statements are wrapped in a new block statement to prevent scope leakage.
func PrependStmts(template *code.Template) *prependStatements {
	_ = "STUB: not implemented"
	return nil
}

// PrependStmtsWithOrder creates a prepend-statements advice with explicit ordering
func PrependStmtsWithOrder(template *code.Template, namespace string, order int) *prependStatements {
	_ = "STUB: not implemented"
	return nil
}

func (a *prependStatements) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a *prependStatements) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *prependStatements) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func (a *prependStatements) Order() int { _ = "STUB: not implemented"; return 0 }

func (a *prependStatements) Namespace() string { _ = "STUB: not implemented"; return "" }

func init() {
	unmarshalers["prepend-statements"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		config := struct {
			Template  code.Template `yaml:",inline"`
			Order     int           `yaml:"order,omitempty"`
			Namespace string        `yaml:"namespace,omitempty"`
		}{
			Order:     DefaultOrder,
			Namespace: DefaultNamespace,
		}

		if err := yaml.NodeToValueContext(ctx, node, &config); err != nil {
			return nil, err
		}

		if config.Namespace == "" {
			// If someone sets it to empty, reset to default.
			config.Namespace = DefaultNamespace
		}

		return &prependStatements{
			Template:  &config.Template,
			order:     config.Order,
			namespace: config.Namespace,
		}, nil
	}
}
