// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package code

import (
	"regexp"
)

// DirectiveArgument represents arguments provided to directives (`//<directive> <args...>`), where
// arguments are parsed as space-separated key:value pairs.
type DirectiveArgument struct {
	Key   string // They key of the argument
	Value string // The value of the argument
}

var spaces = regexp.MustCompile(`\s+`)

// DirectiveArgs returns arguments provided to the named directive. A directive is a single-line
// comment with the directive immediately following the leading `//`, without any spacing in
// between; followed by optional arguments formatted as `key:value`, separated by spaces.
func (d *dot) DirectiveArgs(directive string) (args []DirectiveArgument) {
	_ = "STUB: not implemented"
	return nil
}

// This is not the directive we're looking for -- its name only starts the same.
