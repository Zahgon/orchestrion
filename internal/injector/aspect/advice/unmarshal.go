// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package advice

import (
	"context"

	"github.com/goccy/go-yaml/ast"
)

type unmarshalerFn func(context.Context, ast.Node) (Advice, error)

var unmarshalers = make(map[string]unmarshalerFn)

func FromYAML(ctx context.Context, node ast.Node) (Advice, error) {
	_ = "STUB: not implemented"
	return *new(Advice), nil
}
