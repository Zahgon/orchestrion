// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package jobserver

import (
	"context"
	"sync"
	"time"

	"github.com/DataDog/orchestrion/internal/jobserver/client"
	"github.com/DataDog/orchestrion/internal/jobserver/common"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

const (
	serverUsername = "server" // User for the server itself
	sysUser        = "admin"  // User for system event access
	noPassword     = ""       // We don't need passwords, this is only to have access to system events, not for security.
)

var (
	loopback = "127.0.0.1"
)

type (
	Server struct {
		server     *server.Server     // The underlying NATS server
		CacheStats *common.CacheStats // Cache statistics
		clientURL  string             // The client URL to use for connecting to this server
		log        zerolog.Logger

		shutdownHooks []func(context.Context) error

		// Tracking connected clients for automatic shutdown on inactivity...
		clients           map[uint64]string
		shutdownTimer     *time.Timer
		clientsMu         sync.Mutex
		inactivityTimeout time.Duration
	}

	Options struct {
		// Port is the port on which the server will listen for connections. If
		// zero, a random available port is used.
		Port int
		// StartTimeout is the maximum time to wait for the server to become ready
		// to accept connections. If zero, the default timeout of 5 seconds is used.
		StartTimeout time.Duration
		// EnableLogging enables server logging.
		EnableLogging bool
		// InactivityTimeout is the maximum time to wait for a ping from a client
		// before automatically shutting down. If zero, the server will not shut
		// down automatically.
		InactivityTimeout time.Duration
		// NoListener disables the network listener, only allowing in-process
		// connections to be made to this server instead.
		NoListener bool
	}
)

// New initializes and starts a new NATS server with the provided options. The
// server only listens on the loopback interface.
func New(ctx context.Context, opts *Options) (srv *Server, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Computing defaults

// Creating the server instance

// Starting the server, and waiting for it to be ready

// Shut down the server immediately, as we are returning an error and no
// server, so the caller wouldn't be able to do this by themselves.

// "Any" URL will do here, it's not actually used...

// We don't use `server.ClientURL()` here because it currently returns an invalid URL is the
// listener address is IPv6 (see: https://github.com/nats-io/nats-server/issues/5721)

// Obtaining the local server connection

// Installing the handlers

// Wait until all subscriptions have been processed by the server...

// Wait until all subscriptions have been processed by the server...

// We don't have any (external) client just yet (we've not yet advertised our URL!), so we can start the inactivity
// timer right away.

// Ready for business

func (s *Server) onShutdown(cb func(context.Context) error) { _ = "STUB: not implemented"; return }

// Connect returns a client using the in-process connection to the server.
func (s *Server) Connect() (*client.Client, error) { _ = "STUB: not implemented"; return nil, nil }

// ClientURL returns the URL connection string clients should use to connect to
// this NATS server.
func (s *Server) ClientURL() string {
	_ = "STUB: not implemented"

	// Shutdown initiates the shutdown of this server.
	return ""
}

func (s *Server) Shutdown() { _ = "STUB: not implemented"; return }

// WaitForShutdown waits indefinitely for this server to have shut down.
func (s *Server) WaitForShutdown() { _ = "STUB: not implemented"; return }

func (s *Server) handleClients(msg *nats.Msg) { _ = "STUB: not implemented"; return }

func (s *Server) handleClientConnect(msg *nats.Msg) {
	_ = "STUB: not implemented"

	// Acquire the lock early, so that we process the request before the automatic shutdown happens,
	// since we are most likely going to be cancelling it.
	return
}

// We don't count the server user (it shouldn't disconnect, ever!)

func (s *Server) handleClientDisconnect(msg *nats.Msg) { _ = "STUB: not implemented"; return }

// We don't count the server user (it shouldn't disconnect, ever!)

// startShutdownTimer initiates the automated shutdown timer. The caller must guarantee it has exclusive access to the
// underlying server instance (e.g, during initialization), or have acquired `s.clientsMu`.
func (s *Server) startShutdownTimer() { _ = "STUB: not implemented"; return }

var getLoopbackOnce sync.Once

// Tries to identify a loopback IP address from available interfaces. This is
// done to ensure the server will work even if the host runs an IPv6-only
// stack, as we would discover `::1` appropriately.
func getLoopback(log zerolog.Logger) string { _ = "STUB: not implemented"; return "" }
