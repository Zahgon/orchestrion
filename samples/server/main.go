// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"log"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(myHandler),
	}

	log.Fatal(s.ListenAndServe())
}

// myHandler comment on function
func myHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// test comment in function

func instrumentedHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// test comment in function

// comment that is just hanging out unattached
