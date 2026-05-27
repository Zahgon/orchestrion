// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package advice

import (
	gocontext "context"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/typed"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type addStructField struct {
	Name     string
	TypeName typed.TypeName
}

// AddStructField adds a new synthetic field at the tail end of a struct declaration.
func AddStructField(fieldName string, fieldType typed.TypeName) *addStructField {
	_ = "STUB: not implemented"
	return nil
}

func (a *addStructField) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If the type name is qualified, we may need to import the package, too.

func (a *addStructField) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func (a *addStructField) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["add-struct-field"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		var spec struct {
			Name string
			Type string
		}

		if err := yaml.NodeToValueContext(ctx, node, &spec); err != nil {
			return nil, err
		}
		tn, err := typed.NewTypeName(spec.Type)
		if err != nil {
			return nil, err
		}

		return AddStructField(spec.Name, tn), nil
	}
}
