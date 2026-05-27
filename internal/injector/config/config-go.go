// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package config

import (
	"context"
	"errors"

	"github.com/DataDog/orchestrion/internal/injector/aspect"
	"golang.org/x/tools/go/packages"
)

const FilenameOrchestrionToolGo = "orchestrion.tool.go"

var ErrInvalidGoPackage = errors.New("no .go files in package")

// loadGoPackage loads configuration from the specified go package.
func (l *Loader) loadGoPackage(ctx context.Context, pkg *packages.Package) (_ *configGo, err error) {
	_ = "STUB: not implemented"
	// Special-case the `github.com/DataDog/orchestrion` package, we need not
	// parse this one, and should always use the built-in object.
	return nil, nil
}

// This might be explained by a package-level loading error... We only check
// here because "all .go files are excluded by build constraints" is one
// such error that we typically ignore.

// Workaround poor error typing in packages.Load

// loadGoFile loads configuration from the specified go file. Returns nil if the
// file had already been loaded previously, or if the configuration is empty.
func (l *Loader) loadGoFile(ctx context.Context, filename string) (_ []Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Already loaded or empty

type configGo struct {
	imports []Config
	yaml    *configYML
	pkgPath string
}

func (c *configGo) Aspects() []*aspect.Aspect { _ = "STUB: not implemented"; return nil }

func (c *configGo) visit(v Visitor, _ string) error { _ = "STUB: not implemented"; return nil }

func (c *configGo) empty() bool { _ = "STUB: not implemented"; return false }

// packageRoot returns the root directory of the provided package. It must have
// been loaded with the [packages.NeedFiles] mode, which is the case for values
// returned by [Loader.packages].
func packageRoot(pkg *packages.Package) string { _ = "STUB: not implemented"; return "" }
