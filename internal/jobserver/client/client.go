// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package client

import (
	"context"

	"github.com/DataDog/orchestrion/internal/jobserver/common"
	"github.com/nats-io/nats.go"
)

const (
	Username   = "orchestrion"
	NoPassword = "" // We only use account management to have access to system events, not for security.
)

type Client struct {
	conn *nats.Conn
}

// Connect creates a new client connected to the NATS server at the specified
// address. It implements exponential backoff retry logic to handle temporary
// connection issues, especially on slower CI environments.
func Connect(addr string) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Increased for very slow CI environments
// Start with slightly higher delay
// Increased max backoff for slow systems
// Increased connection timeout

// Don't sleep on the last attempt

// Exponential backoff with cap

func New(conn *nats.Conn) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Close() { _ = "STUB: not implemented"; return }

type (
	request[Res any] interface {
		Subject() string
		common.Request[Res]
	}
)

func Request[Res any, Req request[Res]](ctx context.Context, client *Client, req Req) (Res, error) {
	_ = "STUB: not implemented"
	return *new(Res), nil
}
