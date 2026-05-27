// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package injector

import (
	"context"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

// writeModifiedFile writes the modified file to disk after having restored it to Go source code,
// and returns the path to the modified file.
func (i *Injector) writeModifiedFile(ctx context.Context, decorator *decorator.Decorator, file *dst.File) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
