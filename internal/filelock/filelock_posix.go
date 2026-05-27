// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

//go:build unix

package filelock

import (
	"context"
	"os"
)

// rlock places an advisory shared lock on the specified file.
func rlock(f *os.File) error { _ = "STUB: not implemented"; return nil }

// lock places an advisory exclusive lock on the specified file.
func lock(f *os.File) error { _ = "STUB: not implemented"; return nil }

// unlock removes any advisory locks from the specified file.
func unlock(f *os.File) error { _ = "STUB: not implemented"; return nil }

// beforeLockChange is called before the lock state is changed. It is a no-op on
// POSIX platforms, as [syscall.Flock] allows for a lock to be upgraded or
// downgraded freely. It returns `false` if the currently held lock is identical
// to the target state (idempotent), and always returns a `nil` error.
func (m *Mutex) beforeLockChange(_ context.Context, to lockState) (cont bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
