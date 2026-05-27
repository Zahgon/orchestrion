// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

//go:build windows

package filelock

import (
	"context"
	"os"
)

const (
	reserved = 0
	allBytes = ^uint32(0)
)

// rlock places an advisory shared lock on the specified file.
func rlock(f *os.File) error { _ = "STUB: not implemented"; return nil }

// lock places an advisory exclusive lock on the specified file.
func lock(f *os.File) error { _ = "STUB: not implemented"; return nil }

func winLock(f *os.File, lockType uint32) error { _ = "STUB: not implemented"; return nil }

// unlock removes any advisory locks from the specified file.
func unlock(f *os.File) error { _ = "STUB: not implemented"; return nil }

// beforeLockChange is called before the lock state is changed. On Windows, one
// must release the lock before changing it, as attempting to lock an
// segment of a file on which a lock is already held (including by the current
// process) will block indefinitely. It returns `false` if the desired lock is
// the currently held lock (idempotent success).
func (m *Mutex) beforeLockChange(ctx context.Context, to lockState) (cont bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// No-op, the file is not currently locked by this process.

// No-op, the currently held lock is already the expected lock type.

// We need to unlock before acquiring the new lock.
