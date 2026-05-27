// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package main

import (
	"flag"
	"log"
	"path/filepath"
	"regexp"
)

func main() {
	var (
		licensesDir string
		outputFile  string
	)
	flag.StringVar(&licensesDir, "licenses", "/tmp/licenses", "directory containing all pages license information")
	flag.StringVar(&outputFile, "output", "LICENSE-3rdparty.csv", "output file to write")
	flag.Parse()

	licensesDir = filepath.Clean(licensesDir)
	outputFile = filepath.Clean(outputFile)

	var store Licenses
	for _, filename := range flag.Args() {
		log.Printf("Loading data from %s\n", filename)
		if err := store.LoadFile(filename); err != nil {
			log.Fatalf("Failed to load %q: %v\n", filename, err)
		}
	}

	log.Println("Adding copyright information")
	store.AddCopyrights(licensesDir)

	log.Printf("Writing data to %s\n", outputFile)
	if err := store.WriteFile(outputFile); err != nil {
		log.Fatalf("Error writing %q: %v\n", outputFile, err)
	}
}

type (
	Licenses struct {
		data map[string]map[string]*License
	}
	License struct {
		spdx      []string
		copyright string
	}
)

func (l *Licenses) LoadFile(filename string) error { _ = "STUB: not implemented"; return nil }

func (l *Licenses) AddCopyrights(pkgDir string) { _ = "STUB: not implemented"; return }

func (l *Licenses) WriteFile(filename string) error { _ = "STUB: not implemented"; return nil }

func scanPkg(pkg string) string { _ = "STUB: not implemented"; return "" }

var hasDigits = regexp.MustCompile(`\d`)

func isCopyright(line string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func scanFile(fname string) []string { _ = "STUB: not implemented"; return nil }
