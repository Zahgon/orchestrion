// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package filelock

import (
	"io"
)

var _ io.ReadWriteSeeker = (*Mutex)(nil)
var _ io.ReaderAt = (*Mutex)(nil)
var _ io.WriterAt = (*Mutex)(nil)

func (m *Mutex) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Mutex) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Mutex) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Mutex) ReadAt(b []byte, off int64) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Mutex) WriteAt(b []byte, off int64) (int, error) { _ = "STUB: not implemented"; return 0, nil }
