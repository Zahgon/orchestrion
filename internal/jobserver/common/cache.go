// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package common

import (
	"sync"
	"sync/atomic"
)

type (
	Cache[V any] struct {
		mu    sync.RWMutex
		data  map[string]*slot[V]
		stats *CacheStats
	}
	slot[V any] struct {
		mu     sync.Mutex
		loaded bool
		value  V
	}

	CacheStats struct {
		total atomic.Uint64
		hits  atomic.Uint64
	}
)

func NewCache[V any](stats *CacheStats) Cache[V] { _ = "STUB: not implemented"; return nil }

// Load retrieves the value for the provided key in the cache, loading it using
// the provided callback function if it is not already present. If the loader
// returns an error, the cache slot is not marked as loaded, and the error is
// returned as-is.
func (c *Cache[V]) Load(key string, loader func() (V, error)) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

// slot returns the cache slot for a given key. If the slot does not exist, a
// new, un-loaded slot is allocated and stored.
func (c *Cache[V]) slot(key string) *slot[V] { _ = "STUB: not implemented"; return nil }

// Check again, since we non-atomically upgraded to a write lock

// Allocate the new slot and return it

func (c *CacheStats) RecordHit() { _ = "STUB: not implemented"; return }

func (c *CacheStats) RecordMiss() {
	_ = "STUB: not implemented"

	// Hits returns the count of cache accesses that resulted in a hit.
	return
}

func (c *CacheStats) Hits() uint64 { _ = "STUB: not implemented"; return 0 }

// Count returns the total count of cache accesses.
func (c *CacheStats) Count() uint64 { _ = "STUB: not implemented"; return 0 }

func (c *CacheStats) String() string { _ = "STUB: not implemented"; return "" }
