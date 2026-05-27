// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package traceutil

import (
	"github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
	"github.com/nats-io/nats.go"
)

type NATSCarrier struct {
	*nats.Msg
}

var _ tracer.TextMapReader = (*NATSCarrier)(nil)
var _ tracer.TextMapWriter = (*NATSCarrier)(nil)

func (c NATSCarrier) Set(key string, value string) { _ = "STUB: not implemented"; return }

func (c NATSCarrier) ForeachKey(handler func(key string, val string) error) error {
	_ = "STUB: not implemented"
	return nil
}
