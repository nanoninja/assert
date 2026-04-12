// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"encoding/json"
	"testing"
)

// JSONEqual asserts that two JSON strings are semantically equal,
// ignoring key order and whitespace differences.
func JSONEqual(t testing.TB, expected, actual string, msg ...string) {
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
			append([]string{"JSON values are not equal"}, msg...)...)
	}
}
