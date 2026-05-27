// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package fingerprint

import (
	"crypto/sha512"
	"hash"
	"sync"
)

type Hasher struct {
	hash hash.Hash
}

type Hashable interface {
	Hash(h *Hasher) error
}

var pool = sync.Pool{New: func() any { return &Hasher{hash: sha512.New()} }}

// New returns a [Hasher] from the pool, ready to use.
func New() *Hasher { _ = "STUB: not implemented"; return nil }

// Close returns this [Hasher] to the pool.
func (h *Hasher) Close() { _ = "STUB: not implemented"; return }

// Finish obtains this [Hasher]'s current fingerprint. It does not change the
// underlying state of the [Hasher].
func (h *Hasher) Finish() string { _ = "STUB: not implemented"; return "" }

// Named hashes a named list of values. This creates explicit grouping of the
// values, avoiding that the concatenation of two things has a different hash
// than those same two things one after the other.
func (h *Hasher) Named(name string, vals ...Hashable) error { _ = "STUB: not implemented"; return nil }

// Start of key-value-pair beacon
// Start of value & end of key beacon
// End of key-value-pair beacon

// Fingerprint is a short-hand for creating a new [Hasher], calling
// [Hashable.Hash] on the provided value (unless it is nil), and then returning
// the [Hasher.Finish] result.
func Fingerprint(val Hashable) (string, error) { _ = "STUB: not implemented"; return "", nil }
