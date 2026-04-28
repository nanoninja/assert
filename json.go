// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"encoding/json"
	"testing"
)

// EqualJSON asserts that a Go value marshals to a JSON string semantically equal
// to the expected JSON string, ignoring key order and whitespace.
func EqualJSON(t testing.TB, expected any, actualJSON string, msgAndArgs ...any) {
	t.Helper()

	data, err := json.Marshal(expected)
	if err != nil {
		t.Errorf("\nEqualJSON() failed to marshal expected value: %v", err)
		return
	}

	JSONEqual(t, string(data), actualJSON, msgAndArgs...)
}

// JSONEqual asserts that two JSON strings are semantically equal,
// ignoring key order and whitespace differences.
func JSONEqual(t testing.TB, expected, actual string, msgAndArgs ...any) {
	t.Helper()

	var expectedObj any
	var actualObj any

	if err := json.Unmarshal([]byte(expected), &expectedObj); err != nil {
		t.Errorf("\nJSONEqual() invalid expected JSON: %v", err)
		return
	}

	if err := json.Unmarshal([]byte(actual), &actualObj); err != nil {
		t.Errorf("\nJSONEqual() invalid actual JSON: %v", err)
		return
	}

	if !isEqual(expectedObj, actualObj) {
		failCompare(t, expected, actual,
			append([]any{"JSON values are not equal"}, msgAndArgs...)...)
	}
}
