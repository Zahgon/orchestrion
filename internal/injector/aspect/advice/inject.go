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

type injectDeclarations struct {
	Template *code.Template
	Links    []string
}

// InjectDeclarations merges all declarations in the provided source file into the current file. The package name of both
// original & injected files must match.
func InjectDeclarations(template *code.Template, links []string) injectDeclarations {
	_ = "STUB: not implemented"
	return *new(injectDeclarations)
}

func (a injectDeclarations) Apply(ctx context.AdviceContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Add the declarations to the file

// Register any link-time dependencies that were declared...

// For go:linkname

func (a injectDeclarations) Hash(h *fingerprint.Hasher) error {
	_ = "STUB: not implemented"
	return nil
}

func (a injectDeclarations) AddedImports() []string { _ = "STUB: not implemented"; return nil }

func init() {
	unmarshalers["inject-declarations"] = func(ctx gocontext.Context, node ast.Node) (Advice, error) {
		var config struct {
			Template *code.Template `yaml:",inline"`
			Links    []string
		}
		if err := yaml.NodeToValueContext(ctx, node, &config); err != nil {
			return nil, err
		}

		return InjectDeclarations(config.Template, config.Links), nil
	}
}
