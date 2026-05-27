// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package fingerprint

type Bool bool

func (b Bool) Hash(h *Hasher) error { _ = "STUB: not implemented"; return nil }

type Int int

func (i Int) Hash(h *Hasher) error { _ = "STUB: not implemented"; return nil }

type List[T Hashable] []T

func (l List[T]) Hash(h *Hasher) error { _ = "STUB: not implemented"; return nil }

type String string

func (s String) Hash(h *Hasher) error { _ = "STUB: not implemented"; return nil }

func Cast[E any, T ~[]E, H Hashable](slice T, fn func(E) H) List[H] {
	_ = "STUB: not implemented"
	return nil
}

type (
	mapped[T Hashable]     []mappedItem[T]
	mappedItem[T Hashable] struct {
		key string
		val T
	}
)

func Map[K comparable, V any, H Hashable](m map[K]V, fn func(K, V) (string, H)) mapped[H] {
	_ = "STUB: not implemented"
	return nil
}

func (m mapped[T]) Hash(h *Hasher) error { _ = "STUB: not implemented"; return nil }

func (m mappedItem[T]) Hash(h *Hasher) error { _ = "STUB: not implemented"; return nil }
