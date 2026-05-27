// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package config

import (
	_ "embed" // For go:embed
	"sync"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateObject checks the provided object for conformance to the embedded
// JSON schema. Returns an error if the object does not conform to the schema.
func ValidateObject(obj map[string]any) error { _ = "STUB: not implemented"; return nil }

var (
	//go:embed "schema.json"
	schemaBytes []byte
	schema      *gojsonschema.Schema
	schemaOnce  sync.Once
)

func getSchema() *gojsonschema.Schema { _ = "STUB: not implemented"; return nil }

func compileSchema() { _ = "STUB: not implemented"; return }
