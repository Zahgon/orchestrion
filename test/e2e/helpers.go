// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

//go:build e2e

package e2e

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// FindOrchestrionBinary locates or builds the orchestrion binary for testing
func FindOrchestrionBinary(t *testing.T) string {
	_ = "STUB: not implemented"

	// Find the repository root (where go.mod for orchestrion is)
	return ""
}

// Try to build it

// findRepoRoot finds the orchestrion repository root by looking for go.mod
func findRepoRoot() (string, error) {
	_ = "STUB: not implemented"
	// Start from current directory
	return "", nil
}

// Walk up until we find the orchestrion go.mod (not a test one)

// Check if this is the main orchestrion module

// RunAndLog executes a command, logs output to a file, and checks for errors
func RunAndLog(t *testing.T, cmd *exec.Cmd, logPath string, log func(string, ...interface{})) {
	_ = "STUB: not implemented"

	// Capture both stdout and stderr
	return
}

// Write to log file

// Check for "mismatched build ID" error (the bug we fixed)

// CopyDir recursively copies a directory from src to dst
func CopyDir(t *testing.T, src, dst string) { _ = "STUB: not implemented"; return }

// Skip hidden directories and common non-source directories

// Copy file

// Logger creates a logging function that writes to both test output and a log file
func Logger(t *testing.T, logFile *os.File) func(string, ...interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// CreateWorkDir creates a timestamped work directory in the repo and copies test files into it
func CreateWorkDir(t *testing.T, testDir string) string {
	_ = "STUB: not implemented"

	// Find repo root
	return ""
}

// Create work directory in repo's test/e2e/work/

// Create timestamped directory for this test run

// Clean up old directories for this test, keeping only recent 5

// Clean up at test end if test passed

// Get absolute path of test directory

// Copy test files to work directory

// Fix go.mod to use absolute paths in replace directives

// Replace relative path with absolute path
// The testdata/pgo/go.mod uses ../../../../ (4 levels up from test/e2e/testdata/pgo to orchestrion root)

// cleanupOldWorkDirs keeps only the N most recent directories for a given test
func cleanupOldWorkDirs(t *testing.T, workBase, testName string, keep int) {
	_ = "STUB: not implemented"
	return
}

// Directory doesn't exist yet or can't read it

// Filter directories matching this test name

// If we have more than we want to keep, delete the oldest ones

// Sort by name (which includes timestamp, so newest first)
// Directory names are like: TestPGO-2024-10-16-18-30-45

// Sort by name descending (newest first due to timestamp format)

// Delete all but the N most recent

// WaitForCommandWithTimeout runs a command and waits for it to complete with a timeout
func WaitForCommandWithTimeout(t *testing.T, cmd *exec.Cmd, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
