// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package parse

import (
	"context"
	"go/ast"
	"go/token"
	"sync/atomic"

	"github.com/DataDog/orchestrion/internal/injector/aspect"
	"golang.org/x/sync/errgroup"
)

// maxBytesEagerness is the maximum number of bytes the files of a certain package can have before
// stop we decide to stop trying to run join.Point.FileMayMatch on each file.
// Since 99% of package have no aspects that ACTUALLY match on them, we can save a lot of time by
// applying the join.Point.FileMayMatch heuristic. But if the package has a lot of files, we may
// end up parsing all files anyway so we can just skip this heuristic if the package is too big.
const maxBytesEagerness = 1 << 19 // 512 KiB

type rawFile struct {
	name       string
	mappedName string
	content    []byte
}

// File represents a parsed file with its name, its AST and with the aspects that may match on it.
type File struct {
	// Name is the name of the file.
	Name string
	// AstFile is the parsed AST of the file, cannot be nil
	AstFile *ast.File
	// Aspects is the list of aspects that may match on this file.
	Aspects []*aspect.Aspect
}

type Parser struct {
	fset *token.FileSet // thread-safe data structure

	// rawFiles is an intermediary data structure to store the raw content of the files before parsing them.
	rawFiles []rawFile

	// filesBytesCount is the sum of the bytes of all files parsed so far.
	filesBytesCount atomic.Uint64

	// mustParseAll is a flag that is set to true if at least one file has been parsed.
	// at this point all files must be parsed. It also signals that an aspect matched on a file.
	mustParseAll atomic.Bool

	// parsedFiles is what is returned by ParseFiles.
	parsedFiles []File

	wg errgroup.Group
}

// NewParser creates a new parser with the given [token.FileSet] and the number of files to parse.
func NewParser(fset *token.FileSet, nbFiles int) *Parser { _ = "STUB: not implemented"; return nil }

// ParseFiles return either zero files if no aspect matched on any file of the package,
// or all files parsed with their respective aspects that can match on them.
func (p *Parser) ParseFiles(ctx context.Context, files []string, aspects []*aspect.Aspect) ([]File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// While the current package still has a chance to not require all files to be parsed, we can try to filter out
// aspects that cannot match on this file before parsing it.

// No aspects can match on this file, no need to fill up the File.AstFile field.

// No aspects can match on this package, return nothing

// If we arrived here, this means we need to parse all files anyway because the type-checking pass will need them.

// hasApplicableAspects returns true if the parser should parse all files because at least one file requires it.
func (p *Parser) hasApplicableAspects() bool { _ = "STUB: not implemented"; return false }

func (p *Parser) parseFile(ctx context.Context, rawFile rawFile, aspects []*aspect.Aspect) (File, error) {
	_ = "STUB: not implemented"
	return *new(File), nil
}

func (p *Parser) parseMissingFiles(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip files that have already been parsed.

// fileFilterAspects filters out aspects for a specific file and returns a copy of them
func (p *Parser) fileFilterAspects(aspects []*aspect.Aspect, file rawFile) ([]*aspect.Aspect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readFile(filename string) (rawFile, error) {
	_ = "STUB: not implemented"
	return *new(rawFile), nil
}

// If the file begins with a "//line <path>:1:1" directive, we consume it and
// then pretend the "<path>" was our filename all along. This simplifies
// handling of line offsets further down the line and removes some duplicated
// effort to do it early.
