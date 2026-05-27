// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quoted provides string manipulation utilities. This is elements
// copied from the go command's internals:
// https://github.com/golang/go/blob/go1.23rc2/src/cmd/go/internal/base/goflags.go
package quoted

func isSpaceByte(c byte) bool { _ = "STUB: not implemented"; return false }

// Split splits s into a list of fields,
// allowing single or double quotes around elements.
// There is no unescaping or other processing within
// quoted fields.
//
// Keep in sync with cmd/dist/quoted.go
func Split(s string) ([]string, error) {
	_ = "STUB: not implemented"
	// Split fields allowing ” or "" around elements.
	// Quotes further inside the string do not count.
	return nil, nil
}

// Accepted quoted string. No unescaping inside.
