// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Package config contains APIs used to work with injector configuration files,
// which are formed by [FilenameOrchestrionToolGo] and [FilenameOrchestrionYML] files.
package config

import (
	"context"

	"github.com/DataDog/orchestrion/internal/injector/aspect"
	"golang.org/x/tools/go/packages"
)

// Config represents an injector's configuration. It can be obtained using
// [Loader.Load].
type Config interface {
	// Aspects returns all aspects defined in this configuration in a single list.
	Aspects() []*aspect.Aspect

	visit(Visitor, string) error
}

type PackageLoader = func(context.Context, string, ...string) ([]*packages.Package, error)

// HasConfig determines whether the specified package contains injector
// configuration, and optionally validates it. If the [PackageLoader] is nil,
// a default implementation is used.
func HasConfig(ctx context.Context, pkgLoader PackageLoader, pkg *packages.Package, validate bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// It contains no .go file, so it can't contain configuration.

// Loader is a facility to load configuration from available sources.
type Loader struct {
	pkgLoader PackageLoader
	loaded    map[string]struct{}
	dir       string
	validate  bool
}

func defaultPackageLoader(ctx context.Context, dir string, patterns ...string) ([]*packages.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLoader creates a new [Loader] in the specified directory.
//
//	If the [PackageLoader] is nil, a default implementation is used.
//
// The directory is used to resolve relative paths and must be a valid Go
// package directory, meaning it must contain at least one `.go` file. If
// [Loader.validate] is true, the YAML documents will be validated against the
// JSON schema.
func NewLoader(pkgLoader PackageLoader, dir string, validate bool) *Loader {
	_ = "STUB: not implemented"
	return nil
}

// Load proceeds to load the configuration from this loader's directory.
func (l *Loader) Load(ctx context.Context) (_ Config, err error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

// This is not supposed to happen if `err == nil`.

// markLoaded marks the specified file as loaded. Return true if the file was
// not already marked previously.
func (l *Loader) markLoaded(filename string) bool { _ = "STUB: not implemented"; return false }

func (l *Loader) packages(ctx context.Context, patterns ...string) ([]*packages.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
