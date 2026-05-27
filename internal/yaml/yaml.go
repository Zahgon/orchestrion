// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package yaml

import (
	"context"
	"io"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
)

type decoderContextKey struct{}

type (
	Decoder                = yaml.Decoder
	NodeUnmarshalerContext = yaml.NodeUnmarshalerContext
)

// NewDecoderContext creates a new [yaml.Decoder] and binds it to the given
// [context.Context].
func NewDecoderContext(ctx context.Context, rd io.Reader, opts ...yaml.DecodeOption) (context.Context, *yaml.Decoder) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// UnmarshalYAML unmarshals from the given reader.
func UnmarshalContext(ctx context.Context, rd io.Reader, val any) error {
	_ = "STUB: not implemented"
	return nil
}

// NodeToValueContext unmarshals from the given node.
func NodeToValueContext(ctx context.Context, node ast.Node, val any) error {
	_ = "STUB: not implemented"
	return nil
}
