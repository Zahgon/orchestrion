// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"embed" // For go:embed
	"html/template"
	"strings"

	"github.com/DataDog/orchestrion/internal/injector/config"
)

var (
	//go:embed "*.tmpl"
	templateFS embed.FS
	templates  *template.Template
)

type Generator struct {
	// Dir is the directory in which to write the generated files.
	Dir string
	// ConfigSource is the directory from which to load configuration.
	ConfigSource string
	// Validate determines whether the config parser runs in validation mode.
	Validate bool

	// CommonPrefix is used to trim common prefixes from package paths when
	// generating doc titles.
	CommonPrefix string
	// TrimPrefix is used to trim common prefixes from package paths when
	// generating doc titles. This is done after [Generator.CommonPrefix].
	TrimPrefix string
	// TrimSuffix is used to remove common suffixes from package paths when
	// generating doc titles.
	TrimSuffix string

	generatedFiles map[string]struct{}
}

func (g *Generator) Generate() (err error) { _ = "STUB: not implemented"; return nil }

func (g *Generator) updateToolsFile(cfg config.Config) error { _ = "STUB: not implemented"; return nil }

func (g *Generator) renderPackage(pkgPath string, files []config.File) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanupDir removes files from [Generator.Dir] that are found to no longer be
// part of the generation set, so that only needed files are left.
func (g *Generator) cleanupDir() error { _ = "STUB: not implemented"; return nil }

// Always keep the root `_index.md` file.

func init() {
	funcs := template.FuncMap{
		"packageName": packageName,
		"render":      render,
		"safe":        func(s string) template.HTML { return template.HTML(s) },
		"tabIndent":   tabIndent,
		"trim":        func(s template.HTML) template.HTML { return template.HTML(strings.TrimSpace(string(s))) },
	}

	templates = template.Must(template.New("").
		Funcs(funcs).
		ParseFS(templateFS, "*.tmpl"))
}
