// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package injector

var lineDirectivePrefix = []byte("//line ")

// postProcess modifies the provided source text to remove leading white space ahead of //line directives, as they are
// otherwise ignored by the compiler (directives must begin at the first column of a line). It modifies the input slice
// to avoid re-allocating (since the output is guaranteed to be the same size or smaller than the input).
func postProcess(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// This line has a line directive with leading white space. We need to remove that whilte space, otherwise the
// directive will be ignored by the compiler. To do so we move the source data left by the padding amount, and
// then truncate the source slice to its new length.

// Advance to the next line

// This was the last line: we can break out.
