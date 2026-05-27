// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package may

import (
	"index/suffixarray"
	"sync"
)

// PackageContext is the context for a package to be matched.
type PackageContext struct {
	// ImportPath is the import path of the package in its module.
	ImportPath string

	// ImportMap is the map of import paths to their respective package archives
	ImportMap map[string]string

	// TestMain is true if the package is a test package.
	TestMain bool
}

func (ctx *PackageContext) PackageImports(path string) MatchType {
	_ = "STUB: not implemented"
	return *new(MatchType)
}

// FileContext is the context for a file to be matched.
type FileContext struct {
	// FileContent is the content of the file to be matched.
	FileContent []byte

	// PackageName is the name of the package given as seen in `package main` for example.
	PackageName string

	once  sync.Once
	index *suffixarray.Index
}

func (ctx *FileContext) FileContains(content string) MatchType {
	_ = "STUB: not implemented"
	return *new(MatchType)
}
