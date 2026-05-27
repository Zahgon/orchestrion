// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

// Package injector provides a facility to inject code into go programs, either
// in source (intended to be checked in by the user) or at compilation time
// (via `-toolexec`).
package injector

import (
	gocontext "context"
	"go/importer"
	"go/types"

	"github.com/DataDog/orchestrion/internal/injector/aspect"
	"github.com/DataDog/orchestrion/internal/injector/aspect/context"
	"github.com/DataDog/orchestrion/internal/injector/typed"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver"
)

type (
	// Injector injects go code into a specific Go package.
	Injector struct {
		// ImportPath is the import path of the package that will be injected.
		ImportPath string
		// Name is the name of the package that will be injected. If blank, it will be determined from parsing source files.
		Name string
		// GoVersion is the go runtime version required by this package. If blank, no go runtime compatibility will be
		// asserted.
		GoVersion string
		// TestMain must be set to true when injecting into the generated test main package.
		TestMain bool

		// ImportMap is a map of import paths to their respective .a archive file. Without transitive dependencies
		ImportMap map[string]string

		// ModifiedFile is called to determine the output file name for a modified file. If nil, the input file is modified
		// in-place.
		ModifiedFile func(string) string
		// Lookup is a function that resolves and imported package's archive file.
		Lookup importer.Lookup
		// RootConfig is the root configuration value to use.
		RootConfig map[string]string

		// restorerResolver is used to restore modified files. It's created on-demand then re-used.
		restorerResolver resolver.RestorerResolver
	}

	// InjectedFile contains information about a modified file. It can be used to update compilation instructions.
	InjectedFile struct {
		// References holds new references created while injecting the package, if any.
		References typed.ReferenceMap
		// Filename is the name of the file that needs to be compiled in place of the original one. It may be identical to
		// the input file if the Injector.ModifiedFile function is nil or returns identity.
		Filename string
	}

	parameters struct {
		Decorator *decorator.Decorator
		File      *dst.File
		TypeInfo  types.Info
		Aspects   []*aspect.Aspect
	}

	result struct {
		InjectedFile
		Modified bool
		GoLang   context.GoLangVersion
	}
)

// InjectFiles performs injections on the specified files. All provided file paths must belong to the import path set on
// the receiving Injector. The method returns a map that associates the original source file path to the modified file
// information. It does not contain entries for unmodified files.
func (i *Injector) InjectFiles(ctx gocontext.Context, files []string, aspects []*aspect.Aspect) (_ map[string]InjectedFile, _ context.GoLangVersion, err error) {
	_ = "STUB: not implemented"
	return nil, *new(context.GoLangVersion), nil
}

// We don't want to fail here on type-checking errors... Instead do nothing and let the standard
// go compiler/toolchain surface the error to the user in a canonical way.

func (i *Injector) validate() error { _ = "STUB: not implemented"; return nil }

// Initialize the restorerResolver field, too...

// injectFile injects code in the specified file. This method can be called concurrently by multiple goroutines,
// as is guarded by a sync.Mutex.
func (i *Injector) injectFile(ctx gocontext.Context, decorator *decorator.Decorator, file *dst.File, typeInfo types.Info, aspects []*aspect.Aspect) (result, error) {
	_ = "STUB: not implemented"
	return *new(result), nil
}

func (i *Injector) applyAspects(ctx gocontext.Context, params parameters) (result, error) {
	_ = "STUB: not implemented"
	return *new(result), nil
}

// Pop the ancestry stack now that we're done with this node.

// We only inject synthetic imports here because it may offset declarations by one position in
// case a new import declaration is necessary, which causes dstutil.Apply to re-traverse the
// current declaration.

// injectNode assesses all configured aspects against the current node, and performs any AST
// transformations. It returns whether the AST was indeed modified. In case of an error, the
// injector aborts immediately and returns the error.
func injectNode(ctx context.AdviceContext, aspects []*aspect.Aspect) (mod bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
