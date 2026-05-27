// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package aspect

import (
	"context"
)

// resolvePackageFiles attempts to retrieve the archive for the designated import path. It attempts
// to locate the archive for `importPath` and its dependencies using `go list`. If that fails, it
// will try to resolve it using `go get`.
func resolvePackageFiles(ctx context.Context, importPath string, workDir string) (_ map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Nest the future GOTMPDIR under this $WORK directory, so that builds with `-work` are nested,
// and the root work tree contains all child work trees involved in resolutions.

// Check for missing archives...
