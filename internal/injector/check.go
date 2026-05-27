// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package injector

import (
	"context"
	"go/token"
	"go/types"

	"github.com/DataDog/orchestrion/internal/injector/parse"
)

// typeCheck runs the Go type checker on the provided files, and returns the
// Uses type information map that is built in the process.
func (i *Injector) typeCheck(ctx context.Context, fset *token.FileSet, files []parse.File) (_ types.Info, err error) {
	_ = "STUB: not implemented"
	return *new(types.Info), nil
}

// This is a workaround for the fact that the Go type checker does not return a specific unexported error type
// TODO: Ask better error typing from the Go team for the go/types package

// Not returning a type-checking error here, as this error we want to surface directly to the user ourselves.

type typeCheckingError struct {
	cause error
}

var _ error = typeCheckingError{}

func (e typeCheckingError) Error() string { _ = "STUB: not implemented"; return "" }

func (typeCheckingError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e typeCheckingError) Unwrap() error { _ = "STUB: not implemented"; return nil }
