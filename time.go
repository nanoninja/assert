// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"fmt"
	"testing"
	"time"
)

// WithinDuration asserts that two time values are within a given duration of each other.
func WithinDuration(t testing.TB, expected, actual time.Time, delta time.Duration, msgAndArgs ...any) {
	t.Helper()

	diff := expected.Sub(actual)
	if diff < 0 {
		diff = -diff
	}

	if diff > delta {
		failCompare[any](t,
			fmt.Sprintf("%v (±%v)", expected.Format(time.RFC3339), delta),
			actual.Format(time.RFC3339),
			append([]any{fmt.Sprintf("time difference %v exceeds delta %v", diff, delta)}, msgAndArgs...)...,
		)
	}
}
