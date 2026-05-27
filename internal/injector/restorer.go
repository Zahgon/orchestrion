// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package injector

import (
	"go/importer"
	"go/token"
	"go/types"
	"sync"

	"github.com/dave/dst/decorator"
)

type lookupResolver struct {
	lookup importer.Lookup

	fset    *token.FileSet
	imports map[string]*types.Package

	mu sync.Mutex
}

func (i *Injector) newRestorer(filename string) *decorator.FileRestorer {
	_ = "STUB: not implemented"
	return nil
}

func (r *lookupResolver) ResolvePackage(path string) (string, error) {
	_ = "STUB: not implemented"
	// The "unsafe" package does not have an archive, so it's hard-coded here.
	return "", nil
}

// If this is present in "cache", we can return right away!
