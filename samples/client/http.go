// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"context"
	"net/http"
)

func shortHandsWithContext(context.Context) { _ = "STUB: not implemented"; return }

func shortHandsWithRequest(_ *http.Request /* for context */) { _ = "STUB: not implemented"; return }
