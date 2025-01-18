// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"fmt"
	"testing"
)

// Number represents any numeric type in Go.
// This constraint allows us to work with any built-in numeric type
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 |
		uint16 | uint32 | uint64 | float32 | float64
}

// Ordered represents any type that can be ordered (compared with <, >, <=, >=).
// This includes all numeric types and strings, staying within Go's native types.
type Ordered interface {
	Number | string
}

// Between checks if a value falls within an inclusive range.
// It works with any type that can be ordered (numbers and strings).
func Between[T Ordered](t testing.TB, actual, min, max T) {
	t.Helper()

	if actual < min || actual > max {
		failCompare[any](t,
			fmt.Sprintf("Between %v and %v", min, max),
			actual,
			"value not within expected range",
		)
	}
}

// Greater checks if a value is greater than a minimum value.
func Greater[T Ordered](t testing.TB, actual, min T) {
	t.Helper()

	if actual <= min {
		failCompare[any](t, fmt.Sprintf("> %v", min), actual, "value not greater than minimum")
	}
}

// GreaterOrEqual checks if a value is greater than or equal to a minimum.
// Particularly useful for validating minimum requirements or thresholds.
func GreaterOrEqual[T Ordered](t testing.TB, actual, min T, msg ...string) {
	t.Helper()

	if actual < min {
		failCompare[any](t, actual, fmt.Sprintf(">= %v", min), msg...)
	}
}

// LessOrEqual checks if a value is less than or equal to a maximum.
// This complements our Greater function and is useful for range checks.
func LessOrEqual[T Ordered](t testing.TB, actual, max T, msg ...string) {
	t.Helper()

	if actual > max {
		failCompare[any](t, actual, fmt.Sprintf("<= %v", max), msg...)
	}
}
