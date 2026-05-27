// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"github.com/santhosh-tekuri/jsonschema/v6"
)

func documentSchema(dir string) error { _ = "STUB: not implemented"; return nil }

func documentSchemaInstance(schema *jsonschema.Schema, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func loadSchema() (*jsonschema.Schema, *jsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func normalizeSchema(original any) any { _ = "STUB: not implemented"; return *new(any) }

// Copy `markdownDescription` to `description` because the jsonschema library doesn't support
// the `markdownDescription` attribute, despite it is common in IDE validators... And since we
// render as markdown anyway...

// Normalize all the nested sub-schemas.

// Nothing to do...
