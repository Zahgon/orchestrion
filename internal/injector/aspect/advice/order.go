// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package advice

const (
	// DefaultNamespace is used when no namespace is specified.
	// Uses the highest Unicode code point to ensure default advice
	// always executes last, giving users full control over execution
	// order by choosing any explicit namespace.
	DefaultNamespace = "\U0010ffff"

	// DefaultOrder is used when no order is specified
	DefaultOrder = 0
)

// OrderedAdvice wraps an Advice with ordering information for deterministic
// execution order across aspects.
type OrderedAdvice struct {
	Advice

	AspectID string
	Index    int // Original definition order for stable sorting

	namespace string
	order     int
}

// NewOrderedAdvice creates a new OrderedAdvice with default values
func NewOrderedAdvice(aspectID string, advice Advice, index int) *OrderedAdvice {
	_ = "STUB: not implemented"
	return nil
}

// Sort sorts advice from multiple aspects and returns them in execution order.
// It handles both orderable and non-orderable advice, providing deterministic sorting
// based on namespace, order, and original definition order.
func Sort(orderedAdvice []*OrderedAdvice) { _ = "STUB: not implemented"; return }

func adviceSorter(a *OrderedAdvice, b *OrderedAdvice) int { _ = "STUB: not implemented"; return 0 }
