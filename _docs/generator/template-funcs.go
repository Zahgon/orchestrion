// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"html/template"
	"path/filepath"
	"regexp"
	"runtime"
)

var (
	_, thisFile, _, _ = runtime.Caller(0)
	docsDir           = filepath.Join(thisFile, "..", "..")

	packageNames = make(map[string]string)
)

func packageName(pkgPath string) (name string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func render(val any) (template.HTML, error) {
	_ = "STUB: not implemented"
	return *new(template.HTML), nil
}

var indentRe = regexp.MustCompile(`(?m)^(  )+`)

func tabIndent(s string) string { _ = "STUB: not implemented"; return "" }

func camelToKebab(text string) string { _ = "STUB: not implemented"; return "" }
