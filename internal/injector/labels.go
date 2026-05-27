// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package injector

import (
	"context"
	"sync"

	"github.com/dave/dst"
)

const (
	ddIgnore          = "//dd:ignore"
	orchestrionIgnore = "//orchestrion:ignore"
)

var warnOnce sync.Once

// isIgnored returns true if the node is prefixed by an `//orchestrion:ignore` (or the legacy `//dd:ignore`) directive.
func isIgnored(ctx context.Context, node dst.Node) bool { _ = "STUB: not implemented"; return false }
