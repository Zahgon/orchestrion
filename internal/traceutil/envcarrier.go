// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package traceutil

import (
	"github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
)

type EnvVarCarrier struct {
	Env *[]string
}

var _ tracer.TextMapReader = (*EnvVarCarrier)(nil)
var _ tracer.TextMapWriter = (*EnvVarCarrier)(nil)

const envVarPrefix = "DD_X_"

func (c EnvVarCarrier) ForeachKey(handler func(key string, val string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c EnvVarCarrier) Set(key string, value string) { _ = "STUB: not implemented"; return }

func envVarStyle(key string) string { _ = "STUB: not implemented"; return "" }

func headerStyle(key string) string { _ = "STUB: not implemented"; return "" }
