// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package join

import (
	gocontext "context"
	"errors"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/aspect/may"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/goccy/go-yaml/ast"
)

type importPath string

func ImportPath(name string) importPath { _ = "STUB: not implemented"; return *new(importPath) }

func (p importPath) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

// Technically the current package in this instance

func (p importPath) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (importPath) FileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (p importPath) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (p importPath) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type packageName string

func PackageName(name string) packageName { _ = "STUB: not implemented"; return *new(packageName) }

func (packageName) ImpliesImported() []string {
	_ = "STUB: not implemented"
	// Can't assume anything here...
	return nil
}

func (packageName) PackageMayMatch(_ *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (p packageName) FileMayMatch(ctx *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (p packageName) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (p packageName) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

type packageFilter struct {
	root    bool   // true if targeting the root module only, false for global matching
	pattern string // glob pattern with ** support for import path matching
}

// PackageFilter creates a package filter join point that matches import paths using glob patterns.
//
// If root is true, only matches packages within the current Go module and applies the pattern
// to relative paths within the module. If root is false, matches packages from any module
// using the full import path.
//
// Supports standard glob patterns plus ** (globstar) for recursive matching:
//   - * matches any sequence within a path segment
//   - ** matches any sequence across multiple path segments
//   - ? matches any single character except path separator
//   - [class] matches any character in the character class
//
// Examples:
//
//	PackageFilter(true, "internal/*")         - matches internal packages in root module only
//	PackageFilter(false, "**/internal/*")    - matches internal packages at any depth
//	PackageFilter(false, "github.com/myorg/**") - matches any package under myorg
func PackageFilter(root bool, pattern string) packageFilter {
	_ = "STUB: not implemented"
	return *new(packageFilter)
}

// globMatch extends path.Match to support ** (globstar) patterns.
func globMatch(pattern string, importPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// matchWithGlobstar handles patterns containing ** using segment-by-segment matching.
func matchWithGlobstar(pattern string, importPath string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// expandGlobstarSegments processes pattern segments to handle ** within segments.
// For example, "service**" becomes ["service*", "**"]
func expandGlobstarSegments(patternSegments []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// expandSingleSegment handles the expansion logic for a single pattern segment
func expandSingleSegment(segment string) []string { _ = "STUB: not implemented"; return nil }

// Pure globstar segment, keep as is.

// Mixed segment like "service**", split it.

// Mixed segment like "**service", split it.

// Handle ** in the middle of a segment by splitting around it.

// Regular segment, keep as is.

// expandMiddleGlobstar handles segments with ** in the middle
func expandMiddleGlobstar(segment string) []string { _ = "STUB: not implemented"; return nil }

func matchSegments(patternSegments []string, pathSegments []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (_ packageFilter) ImpliesImported() []string { _ = "STUB: not implemented"; return nil }

func (pf packageFilter) PackageMayMatch(ctx *may.PackageContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (_ packageFilter) FileMayMatch(_ *may.FileContext) may.MatchType {
	_ = "STUB: not implemented"
	return *new(may.MatchType)
}

func (pf packageFilter) Matches(ctx context.AspectContext) bool {
	_ = "STUB: not implemented"
	return false
}

func (pf packageFilter) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func (pf packageFilter) matchesPattern(importPath string) bool {
	_ = "STUB: not implemented"
	return false
}

// For root-only filters without pattern, match all packages in root module.

// isInRootModule checks if the given import path belongs to the root module.
func isInRootModule(importPath string) bool { _ = "STUB: not implemented"; return false }

// If we can't determine, assume it doesn't match

// getRelativePathInModule returns the relative path of an import path within its module.
func getRelativePathInModule(importPath string, rootModulePath string) string {
	_ = "STUB: not implemented"
	return ""
}

func init() {
	unmarshalers["import-path"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var name string
		if err := yaml.NodeToValueContext(ctx, node, &name); err != nil {
			return nil, err
		}
		return ImportPath(name), nil
	}

	unmarshalers["package-name"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var name string
		if err := yaml.NodeToValueContext(ctx, node, &name); err != nil {
			return nil, err
		}
		return PackageName(name), nil
	}

	unmarshalers["package-filter"] = func(ctx gocontext.Context, node ast.Node) (Point, error) {
		var pattern string
		if err := yaml.NodeToValueContext(ctx, node, &pattern); err == nil {
			return PackageFilter(false, pattern), nil
		}
		var config struct {
			Root    bool   `yaml:"root"`
			Pattern string `yaml:"pattern"`
		}
		if err := yaml.NodeToValueContext(ctx, node, &config); err != nil {
			return nil, err
		}

		if config.Pattern == "" && !config.Root {
			return nil, errors.New("package-filter requires a 'pattern' field")
		}

		return PackageFilter(config.Root, config.Pattern), nil
	}
}
