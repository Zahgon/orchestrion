// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package code

import (
	gocontext "context"
	"text/template"

	"github.com/DataDog/orchestrion/internal/fingerprint"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/version"
	"github.com/DataDog/orchestrion/internal/yaml"
	"github.com/dave/dst"
	"github.com/goccy/go-yaml/ast"
)

type Template struct {
	template *template.Template
	Imports  map[string]string
	Source   string
	Lang     context.GoLangVersion
}

var wrapper = template.Must(template.New("code.Template").Funcs(template.FuncMap{
	"Version": version.Tag,
}).Parse(
	`
{{- define "_statements_" -}}
package _
func _() {
	{{ template "code.Template" . }}
}
{{- end -}}
{{- define "_declarations_" -}}
package _
{{ template "code.Template" . }}
{{- end -}}}
	`,
))

// NewTemplate creates a new Template using the provided template string and
// imports map. The imports map associates names to import paths. The produced
// AST nodes will feature qualified *dst.Ident nodes in all places where a
// property of mapped names is selected.
func NewTemplate(text string, imports map[string]string, lang context.GoLangVersion) (*Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustTemplate is the same as NewTemplate, but panics if an error occurs.
func MustTemplate(text string, imports map[string]string, lang context.GoLangVersion) (template *Template) {
	_ = "STUB: not implemented"
	return nil
}

// CompileBlock generates new source based on this Template and wraps the
// resulting dst.Stmt nodes in a new *dst.BlockStmt. The provided
// context.Context and *dstutil.Cursor are used to supply context information to
// the template functions.
func (t *Template) CompileBlock(ctx context.AdviceContext) (*dst.BlockStmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CompileDeclarations generates new source based on this Template and extracts
// all produced declarations.
func (t *Template) CompileDeclarations(ctx context.AdviceContext) ([]dst.Decl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CompileExpression generates new source based on this Template and extracts
// the produced dst.Expr node. The provided context.Context and *dstutil.Cursor
// are used to supply context information to the template functions. The
// provided dst.Expr will be copied in places where the `{{Expr}}` template
// function is used, unless `expr` is nil.
func (t *Template) CompileExpression(ctx context.AdviceContext) (dst.Expr, error) {
	_ = "STUB: not implemented"
	return *new(dst.Expr), nil
}

// Move the decorations from the statement to the expression itself.

// compile generates new source based on this Template and returns a cloned
// version of minimally post-processed dst.Stmt nodes this produced.
func (t *Template) compile(ctx context.AdviceContext) ([]dst.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Template) compileTemplate(ctx context.AdviceContext, name string) ([]dst.Decl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IMPORTANT: process imports BEFORE replacing placeholders, so we never replace a symbol from the original AST.

// processImports replaces all [*dst.SelectorExpr] based on one of the names
// present in the [Template.Imports] map with a qualified [*dst.Ident] node, so
// that the import-enabled decorator.Restorer can emit the correct code, and
// knows not to remove the inserted import statements.
func (t *Template) processImports(ctx context.AdviceContext, node dst.Decl) dst.Decl {
	_ = "STUB: not implemented"
	return *new(dst.Decl)
}

// We apply an alias to the import to mitigate the risk of conflicting with an existing symbol in the surrounding scope.

func (t *Template) Hash(h *fingerprint.Hasher) error { _ = "STUB: not implemented"; return nil }

func (t *Template) AddedImports() []string { _ = "STUB: not implemented"; return nil }

var _ yaml.NodeUnmarshalerContext = (*Template)(nil)

func (t *Template) UnmarshalYAML(ctx gocontext.Context, node ast.Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func numberLines(text string) string { _ = "STUB: not implemented"; return "" }
