// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package basiclit

import (
	"github.com/dave/dst"
)

// String parses the value of a *dst.BasicLit to a regular go string.
func String(lit *dst.BasicLit) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Note: We blindly assume that Go string literals have the same syntax as JSON here...
