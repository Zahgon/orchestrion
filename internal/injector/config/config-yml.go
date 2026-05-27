// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package config

import (
	"context"

	"github.com/DataDog/orchestrion/internal/injector/aspect"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

const FilenameOrchestrionYML = "orchestrion.yml"

// loadYMLFile loads configuration from the specified directory.
func (l *Loader) loadYMLFile(ctx context.Context, dir string, name string) (_ *configYML, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Already loaded, ignoring...

// This is not supposed to happen if `err == nil`.

// Empty, nothing to do...

// Empty, nothing to do...

type (
	configYML struct {
		extends []Config
		aspects []*aspect.Aspect
		name    string
		meta    configYMLMeta
	}
	configYMLMeta struct {
		name        string
		description string
		icon        string
		caveats     string
	}
)

func (c *configYML) Aspects() []*aspect.Aspect { _ = "STUB: not implemented"; return nil }

func (c *configYML) visit(v Visitor, pkgPath string) error { _ = "STUB: not implemented"; return nil }

func (c *configYML) empty() bool { _ = "STUB: not implemented"; return false }

type ymlFile struct {
	Aspects []*aspect.Aspect
	Extends []string
	Meta    struct {
		Name        string
		Description string
		Icon        string // Optional
		Caveats     string // Optional
	}
}

func (l *Loader) parseYMLFile(ctx context.Context, filename string) (*ymlFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In validation mode, we will pre-parse the YAML into a [yaml.Node] tree,
// which can then be cheaply decoded into a value type that validation
// supports; and then re-decoded into the actual data structure.
// This dance is significantly cheaper (both in time & allocations) than doing
// a full blown [yaml.Decoder.Decode] twice (as it internally transits through
// the [yaml.Node] representation anyway).

type decodedNode struct {
	*yaml.Decoder
	ast.Node
}

func (n decodedNode) DecodeContext(ctx context.Context, out any) error {
	_ = "STUB: not implemented"
	return nil
}

// maskErrNotExist intentionally "breaks" the error chaining if the provided
// error is an [fs.ErrNotExist] so that the returned error is not
// [fs.ErrNotExist]. Otherwise, returns the original error unmodified.
func maskErrNotExist(err error) error { _ = "STUB: not implemented"; return nil }
