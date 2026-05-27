// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package client

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	EnvVarJobserverURL = "ORCHESTRION_JOBSERVER_URL"
	urlFileName        = ".orchestrion-jobserver"
)

var (
	client *Client

	ErrNoServerAvailable = errors.New("no job server is available")
)

// FromEnvironment returns a client connected to the current environment's
// job server, using the following process:
//   - If the ORCHESTRION_JOBSERVER_URL environment variable is set, a client
//     connected to this URL is returned.
//   - Otherwise, if workDir is not empty, a server will be identified based on
//     a `.orchestrion-jobserver` file; or a new server will be started using
//     that url file, and a connection will be established to it. The started
//     job server will automatically shut itself down once it no longer has any
//     active client for a period of time.
//   - Otherwise, the ErrNoServerAvailable error is returned.
//
// The returned client is re-used, so callers should NOT call [Client.Close] on
// it.
func FromEnvironment(ctx context.Context, workDir string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to start a server. The server process is idempotent if the `-url-file` flag is used, so we do not check the
// command's exit status, because another process might act as our server down the line.

// Make sure go doesn't wait for this to exit...
// Suppress the TOOLEXEC_IMPORTPATH variable if it's set.

// Connect to `os.DevNull`

// Wait for the URL file to exist...

// Kill the process if it's still running...

// Detach the process, so it survives this one if needed...

func clientFromURLFile(ctx context.Context, path string) (*Client, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func waitForURLFile(ctx context.Context, path string, cmd *exec.Cmd, exitChan <-chan error) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First, try to connect to the client from the URL file.

// There was no error, so we are good to go!

// Set it in the current environment so that child processes don't have to go through the same dance again.

//revive:disable:defer This happens only once in the loop, and is to avoid starting a timer we don't use

//revive:enable:defer

// If the context is Done, we should not be waiting any longer...

// Attempt to kill the process if it hasn't died by itself...

// If the process has exited, there is no use to waiting any longer...

// The job server exits with status 0 if another process has written to the URL file; in
// which case we should be able to connect to it on the next attempt!

// The retry timer has elapsed, we shall try again!

var jobserverStartTimeout = 5 * time.Second

func init() {
	const envVarName = "ORCHESTRION_JOB_SERVER_START_TIMEOUT_SECONDS"
	val := os.Getenv(envVarName)
	if val == "" {
		return
	}

	sec, err := strconv.Atoi(val)
	if err != nil {
		_, _ = fmt.Fprintf(
			os.Stderr,
			"Warning: unable to parse value of "+envVarName+"=%q due to %v, will use default value of %s instead\n",
			val,
			err,
			jobserverStartTimeout,
		)
		return
	}

	jobserverStartTimeout = time.Duration(sec) * time.Second
}
