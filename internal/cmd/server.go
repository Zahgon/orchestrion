// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package cmd

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/DataDog/orchestrion/internal/goflags"
	"github.com/DataDog/orchestrion/internal/jobserver"
	"github.com/goccy/go-yaml"
	"github.com/rs/zerolog"

	"github.com/urfave/cli/v2"
)

var Server = &cli.Command{
	Name:        "server",
	Usage:       "Start an Objectsrion job server.",
	Description: "The job server is used to remove duplicated processing that can occur when instrumenting large applications, due to how Orchestrion injects new dependencies that the go toolchain was initially not aware of.\n\nUsers do not normally need to use this command directly, as Orchestrion automatically manages servers during runtime.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "url-file",
			Usage: "Write a file containing the ClientURL for this server once it is ready to accept connections. The server automatically shuts down when the URL file is deleted.",
		},
		&cli.IntFlag{
			Name:        "port",
			Usage:       "Choose a port to listen on.",
			Value:       -1,
			DefaultText: "random",
		},
		&cli.DurationFlag{
			Name:  "inactivity-timeout",
			Usage: "Automatically shut down after a period without any connected client.",
			Value: time.Minute,
		},
		&cli.BoolFlag{
			Name:  "nats-logging",
			Usage: "Enable NATS server logging.",
		},
		&cli.StringFlag{
			Name:        "build-flags",
			Usage:       "Specify the 'go build' flags to use when resolving packages. This is specified as a YAML array and must start with a valid go subcommand (e.g, 'build').",
			DefaultText: "Looked up the process hierarchy",
			Action: func(ctx *cli.Context, val string) error {
				var args []string
				if err := yaml.Unmarshal([]byte(val), &args); err != nil {
					return cli.Exit(fmt.Errorf("invalid -build-flags value: %w", err), 2)
				}
				goflags.SetFlags(ctx.Context, ".", args)
				return nil
			},
		},
		&cli.IntFlag{
			Name:        "parent-pid",
			Usage:       "Specify which process created this server. This is useful when the server is started as a daemon, as it needs to be able to resolve the top-level go command line.",
			DefaultText: "This process' parent (may be inaccurate if the process is daemonized)",
			Action: func(ctx *cli.Context, val int) error {
				return goflags.SetFlagsFromPid(ctx.Context, val)
			},
		},
	},
	Hidden: true,
	Action: func(ctx *cli.Context) error {
		log := zerolog.Ctx(ctx.Context)

		opts := jobserver.Options{
			Port:              ctx.Int("port"),
			InactivityTimeout: ctx.Duration("inactivity-timeout"),
			EnableLogging:     ctx.Bool("nats-logging"),
		}

		if urlFile := ctx.String("url-file"); urlFile != "" {
			if err := startWithURLFile(ctx.Context, &opts, urlFile); err != nil {
				log.Error().Err(err).Str("url-file", urlFile).Msg("Failed to start job server")
			}
			return nil
		}
		_, err := start(ctx.Context, &opts, true)
		if err != nil {
			log.Error().Err(err).Msg("Failed to start job server")
		}
		return err
	},
}

// start starts a new job server, and waits for it to have completely shut down if `wait` is true.
// When `wait` is true, the server is always returned as `nil`.
func start(ctx context.Context, opts *jobserver.Options, wait bool) (*jobserver.Server, cli.ExitCoder) {
	_ = "STUB: not implemented"
	return nil, *new(cli.ExitCoder)
}

// startWithURLFile starts a new job server using the provided URL file (unless the file contains the URL to a still
// running server), and waits for it to have completely shut down.
func startWithURLFile(ctx context.Context, opts *jobserver.Options, urlFile string) cli.ExitCoder {
	_ = "STUB: not implemented"
	return *new(cli.ExitCoder)
}

// Check if there is already a server running...

// No existing server, so now we're actually going to try starting our own

// Check again whether there is a running server; as a concurrent process might have acquired the write lock first.

// This process "owns" the URL file, so it'll try had to remove it when it terminates...

// Start the server normally...

// Write the ClientURL into the urlFile

// Release the URL File lock

// Shut the server down, as we won't actually be returning it...

// Try to watch for removal of the URL file, so we can shut down the server eagerly when that happens.

// deleteOnInterrupt attempts to deletes the provided file when an interrupt signal is received. It returns a
// cancellation function that can be used to uninstall the signal handler.
func deleteOnInterrupt(ctx context.Context, path string) func() {
	_ = "STUB: not implemented"
	return nil
}

// hasURLToRunningServer checks whether the provided URL file contains the URL to a running server,
// by trying to connect to it. If that is the case, it returns the URL to the running server.
func hasURLToRunningServer(file io.ReadSeeker) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// shutdownOnRemove shuts the server down when the designated file is removed. It returns a cancellation function that
// can be used to cancel the file watcher. Since fsnotify support is highly dependent on platform/kernel support, this
// function ignores any error and emits WARN log entries describing the problem.
func shutdownOnRemove(ctx context.Context, server *jobserver.Server, urlFile string) func() error {
	_ = "STUB: not implemented"
	return nil

	// noCancel is returned when there is nothing to cancel...
}
