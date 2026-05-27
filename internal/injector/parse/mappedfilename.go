// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package parse

import (
	"io"
)

// ConsumeLineDirective consumes the first line from r if it's a "//line"
// directive that either does not have line/column information or has it set to
// line 1 (and column 1). If the directive is consumed, the filename it refers
// to is returned. Otherwise, the reader is rewound to its original position
// if the provided reader supports the [io.Seeker] interface.
func ConsumeLineDirective(r io.Reader) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Reached EOF

// We saw a CR and this is not an LF, so we rewind one byte and bail out.

// Remove any leading or trailing white space

// It was not at position 1, so it's not the directive we're looking for.

// It was not at position 1, so it's not the directive we're looking for.

// cutPositionSuffix removes a trailing ":<int>" from the provided buffer, if present.
func cutPositionSuffix(buf []byte) ([]byte, int, bool) {
	_ = "STUB: not implemented"
	return nil,

		// First, consume the integer at the end of the buffer.
		0, false
}

// If there's no ":" before the integer, or there was no digit at all, it was not a position...
