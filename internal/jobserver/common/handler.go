// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package common

import (
	"context"

	"github.com/nats-io/nats.go"
)

type (
	Request[Res any] interface {
		ResponseIs(Res)
		ForeachSpanTag(func(key string, val any))
	}
	// RequestHandler is a function that processes a request of a given type, and returns a response or an error to be
	// sent back to the client.
	RequestHandler[Res any, Req Request[Res]] func(context.Context, Req) (Res, error)
)

// HandleRequest returns a NATS subscription target that calls the provided request handler in a new goroutine if the
// NATS message payload can be parsed into the specified request type, and responds to the client appropriately.
func HandleRequest[Res any, Req Request[Res]](ctx context.Context, handler RequestHandler[Res, Req]) func(*nats.Msg) {
	_ = "STUB: not implemented"
	return nil
}

// Spawn the handler in a new goroutine to avoid blocking the NATS subscription poller.

type (
	errorResponse struct {
		Error string `json:"error"`
	}
	successResponse[T any] struct {
		Result T `json:"result"`
	}

	// natsResponse is a marker interface used to make sure the value send to respond is one of the two possible response
	// types, so that it is guaranteed that the UnmarshalResponse function can accept it.
	natsResponse interface {
		isNatsResponse()
	}
)

func (errorResponse) isNatsResponse()      { _ = "STUB: not implemented"; return }
func (successResponse[T]) isNatsResponse() { _ = "STUB: not implemented"; return }

func respond(ctx context.Context, msg *nats.Msg, val natsResponse) {
	_ = "STUB: not implemented"
	return
}

// UnmarshalResponse parses a response received from the job server, either into a result of the specified type, or as
// an error; depending on the response's structure.
func UnmarshalResponse[T any](ctx context.Context, data []byte) (_ T, err error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
