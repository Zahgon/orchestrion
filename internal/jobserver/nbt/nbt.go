// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package nbt

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

const (
	subjectPrefix = "never-build-twice."

	startSubject  = subjectPrefix + "start"
	finishSubject = subjectPrefix + "finish"
)

type (
	service struct {
		state sync.Map
		dir   string
	}
	buildState struct {
		initOnce sync.Once
		buildID  string
		token    string          // Finalization token
		onDone   func()          // Called once the original task has completed
		done     <-chan struct{} // Blocks until the original task has completed

		isDone atomic.Bool      // Whether the original task has completed yet.
		files  map[Label]string // Additional files produced by the original task. Available once done.
		error  error            // Error from the original task. Available once done.
	}
)

func Subscribe(ctx context.Context, conn *nats.Conn) (cleanup func(context.Context) error, resErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type (
	// StartRequest informs the job server that the caller is starting a new
	// compilation task for the specified [StartRequest.ImportPath].
	StartRequest struct {
		ImportPath string `json:"importPath"`
		BuildID    string `json:"buildID"`
	}
	// StartResponse informs the caller about what should be done with the
	// compilation task. If a [*StartResponse.FinishToken] is present, the caller
	// must proceed with the task, then send a [FinishRequest] using that token.
	// If a [*StartResponse.ArchivePath] is present, the caller must skip the task
	// and instead re-use the file at the specified path.
	StartResponse struct {
		// FinishToken is the token to be forwarded to the [FinishRequest] to inform
		// the job server about the outcome of the build. It cannot be blank unless
		// [*StartResponse.ArchivePath] is not blank.
		FinishToken string `json:"token,omitempty"`
		// Files is the set of files produced by the original task, by label. These
		// files must be re-used instead of re-created. Always blank if
		// [*StartResponse.FinishToken] is not blank.
		Files map[Label]string `json:"extra,omitempty"`
	}
	// Label is a label identifying an additional object produced by a task. It
	// must be a valid part of a file name.
	Label string
)

const (
	LabelArchive Label = "_pkg_.a"
	LabelAsmhdr  Label = "go_asm.h"
)

func (StartRequest) Subject() string           { _ = "STUB: not implemented"; return "" }
func (StartRequest) ResponseIs(*StartResponse) { _ = "STUB: not implemented"; return }
func (r StartRequest) ForeachSpanTag(set func(key string, value any)) {
	_ = "STUB: not implemented"
	return
}

// cacheKey creates a composite key from importPath and buildID to support
// different build configurations (e.g., with/without PGO) of the same package.
// See: https://github.com/DataDog/orchestrion/issues/653
func cacheKey(importPath string, buildID string) string { _ = "STUB: not implemented"; return "" }

func (s *service) start(ctx context.Context, req StartRequest) (*StartResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize the build state.

// We use a cancellable context as a barrier here...

// If the build state is re-used, wait for the original to complete...

// The context has expired or the upstream context has been canceled.
// We'll return this as [context.Canceled] either way.

// Otherwise, return a finalization token, etc...

type (
	// FinishRequest informs the job server about the result of a compilation
	// task.
	FinishRequest struct {
		// ImportPath is the import path of the package that was built.
		ImportPath string `json:"importPath"`
		// BuildID is the build ID of the package that was built.
		BuildID string `json:"buildID"`
		// FinishToken is forwarded from [*StartResponse.FinishToken], and cannot be
		// blank.
		FinishToken string `json:"token"`
		// Files is a list of files produced by the compilation task, associated to
		// a user-defined label.
		Files map[Label]string `json:"extra,omitempty"`
		// Error is the error that occurred as a result of this compilation task, if
		// any.
		Error *string `json:"error,omitempty"`
	}
	FinishResponse struct {
		/* unused */
	}
)

func (FinishRequest) Subject() string            { _ = "STUB: not implemented"; return "" }
func (FinishRequest) ResponseIs(*FinishResponse) { _ = "STUB: not implemented"; return }
func (r FinishRequest) ForeachSpanTag(set func(key string, value any)) {
	_ = "STUB: not implemented"
	return
}

var errNoFilesNorError = errors.New("missing files, and no error reported")

func (s *service) finish(ctx context.Context, req FinishRequest) (*FinishResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use composite key for storage directory to support different build IDs (e.g., with/without PGO)

// ns is an arbitrary UUID used as a namespace for hashing import paths when storing artifacts in
// the temporary storage location.
var ns = uuid.MustParse("4BFB6F4B-212C-43A0-A581-A29C8B3D3BE4")
