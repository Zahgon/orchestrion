// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package singleton

import (
	"context"

	"github.com/goccy/go-yaml/ast"
)

func Unmarshal(ctx context.Context, node ast.Node) (key string, value ast.Node, err error) {
	_ = "STUB: not implemented"
	return "", *new(ast.Node), nil
}
