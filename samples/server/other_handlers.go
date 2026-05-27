// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"context"
	"net/http"
)

type foo struct{}

func (foo) fooHandler(rw http.ResponseWriter, _ *http.Request) { _ = "STUB: not implemented"; return }

func buildHandlers() { _ = "STUB: not implemented"; return }

// silly legal things

//dd:span foo:bar type:potato
func myFunc(_ context.Context, name string) {
	_ = "STUB: not implemented"

	//dd:span foo2:bar2 type:request
	return
}

func myFunc2(name string, _ *http.Request) {
	_ = "STUB: not implemented"

	//dd:span foo3:bar3 type:request span.name:customName
	return
}

func myFunc3(name string) error { _ = "STUB: not implemented"; return nil }

func registerHandlers() { _ = "STUB: not implemented"; return }
