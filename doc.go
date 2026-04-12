// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package assert provides a lightweight, type-safe assertion library for Go tests.
// It follows the same conventions as testify while remaining simple and dependency-free.
//
// # Basic Usage
//
//	func TestExample(t *testing.T) {
//	    result, err := Compute(2, 3)
//	    assert.NoError(t, err)
//	    assert.Equal(t, 5, result)
//
//	    users := GetUsers()
//	    assert.NotEmpty(t, users)
//	    assert.Len(t, users, 3)
//	}
//
// # Parameter Convention
//
// All assertion functions follow the testify convention: expected value first, actual value second.
//
//	assert.Equal(t, expected, actual)
//
// # Comparisons
//
//   - [Equal] / [NotEqual]: compare values of any type using reflect.DeepEqual
//   - [True] / [False]: boolean assertions
//   - [Nil] / [NotNil]: nil checks for any type (pointers, slices, maps, channels...)
//   - [IsType]: assert that two values share the same type
//   - [Implements]: assert that a value implements an interface
//
// # Error Handling
//
//   - [Error] / [NoError]: assert presence or absence of an error
//   - [EqualError]: assert an error's message matches an expected string
//   - [ErrorIs]: check error chain with errors.Is
//   - [ErrorAs]: check error chain with errors.As
//   - [Panics] / [NotPanics]: assert panic behavior
//   - [PanicsWithError]: assert a function panics with a specific error message
//
// # Collections
//
//   - [Contains] / [NotContains]: check slice membership
//   - [ElementsMatch]: same elements regardless of order
//   - [Subset] / [NotSubset]: subset relationship between slices
//   - [Empty] / [NotEmpty]: check if a collection is empty
//   - [Len]: check collection length
//   - [HasKey]: check map key existence
//
// # Strings
//
//   - [StringContains]: check substring
//   - [HasPrefix] / [HasSuffix]: check string prefix or suffix
//   - [MatchRegexp]: check against a regular expression
//
// # Numeric Comparisons
//
//   - [Greater] / [Less]: strict comparison
//   - [GreaterOrEqual] / [LessOrEqual]: non-strict comparison
//   - [Between]: check inclusive range
//   - [Positive] / [Negative]: sign assertions
//   - [InDelta]: floating-point comparison with tolerance
//
// # JSON
//
//   - [JSONEqual]: semantic JSON comparison (ignores key order and whitespace)
//
// # Time
//
//   - [WithinDuration]: assert two time values are close within a given duration
//
// # Require Package
//
// The [require] sub-package mirrors the assert API but stops the test immediately
// on failure using t.FailNow(). Use it for preconditions.
//
//	import "github.com/nanoninja/assert/require"
//
//	func TestExample(t *testing.T) {
//	    result, err := GetUser(42)
//	    require.NoError(t, err)      // stops if err != nil
//	    require.NotNil(t, result)    // stops if nil
//
//	    assert.Equal(t, "alice", result.Name)   // continues on failure
//	    assert.True(t, result.Active)           // sees all failures
//	}
//
// # Error Messages
//
// All failed assertions produce structured output:
//
//	file.go:42
//	 Message: optional custom message
//	Expected: (int) 5
//	  Actual: (int) 4
package assert
