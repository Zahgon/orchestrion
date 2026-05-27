// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package importcfg

import (
	"io"
)

// Lookup opens the archive file for the provided import path. This allows the
// ImportConfig object to serve as a package information resolver's Lookup
// function.
func (r *ImportConfig) Lookup(path string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
