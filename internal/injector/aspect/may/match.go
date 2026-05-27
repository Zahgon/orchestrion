// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package may

// MatchType is an enumeration of the possible outcomes of a join point
type MatchType byte

const (
	// Unknown indicates that the join point cannot determine whether it matches the given context
	Unknown MatchType = '?'
	// Match indicates that the join point may match the given context
	Match MatchType = 'Y'
	// NeverMatch indicates that the join point DOES NOT match the given context
	NeverMatch MatchType = 'N'
)

// Not returns the logical NOT of a MatchType value
// Truth table:
//
// | A | NOT A |
// |---|-------|
// | N | Y     |
// | ? | ?     |
// | Y | N     |
func (m MatchType) Not() MatchType { _ = "STUB: not implemented"; return *new(MatchType) }

// Or returns the logical OR of two MatchType values
// Truth table:
//
// | A | B  | A OR B |
// |---|---|---------|
// | N | N | N       |
// | N | ? | ?       |
// | N | Y | Y       |
// | ? | ? | ?       |
// | ? | Y | Y       |
// | Y | Y | Y       |
func (m MatchType) Or(other MatchType) MatchType { _ = "STUB: not implemented"; return *new(MatchType) }

// And returns the logical AND of two MatchType values
// Truth table:
//
// | A | B | A AND B |
// |---|---|---------|
// | N | N | N       |
// | N | ? | N       |
// | N | Y | N       |
// | ? | ? | ?       |
// | ? | Y | ?       |
// | Y | Y | Y       |
func (m MatchType) And(other MatchType) MatchType {
	_ = "STUB: not implemented"
	return *new(MatchType)
}
