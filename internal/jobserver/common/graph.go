// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package common

import (
	"sync"
)

// Graph keeps track of a directed acyclic graph.
type Graph struct {
	nodes map[string]map[string]struct{}
	mu    sync.Mutex
}

// AddEdge adds a new edge to this graph. Returns an error if the new edge would
// introduce a cycle in the graph.
func (g *Graph) AddEdge(from string, to string) error { _ = "STUB: not implemented"; return nil }

// RemoveEdge removes an edge from this graph.
func (g *Graph) RemoveEdge(from string, to string) { _ = "STUB: not implemented"; return }

func (g *Graph) path(from string, to string) []string { _ = "STUB: not implemented"; return nil }
