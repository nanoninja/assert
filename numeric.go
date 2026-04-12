// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"fmt"
	"math"
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
func Between[T Ordered](t testing.TB, actual, min, max T, msgAndArgs ...any) {
	t.Helper()

	if actual < min || actual > max {
		failCompare[any](t,
			fmt.Sprintf("between %v and %v", min, max),
			actual,
			append([]any{"value not within expected range"}, msgAndArgs...)...,
		)
	}
}

// Greater checks if a value is greater than a minimum value.
func Greater[T Ordered](t testing.TB, e1, e2 T, msgAndArgs ...any) {
	t.Helper()

	if e1 <= e2 {
		failCompare[any](t, fmt.Sprintf("> %v", e2), e1, append(
			[]any{"value not greater than minimum"}, msgAndArgs...)...)
	}
}

func Less[T Ordered](t testing.TB, e1, e2 T, msgAndArgs ...any) {
	t.Helper()

	if e1 >= e2 {
		failCompare[any](t, fmt.Sprintf("< %v", e2), e1, append(
			[]any{"value not less than maximum"}, msgAndArgs...)...)
	}
}

// GreaterOrEqual checks if a value is greater than or equal to a minimum.
// Particularly useful for validating minimum requirements or thresholds.
func GreaterOrEqual[T Ordered](t testing.TB, actual, min T, msgAndArgs ...any) {
	t.Helper()

	if actual < min {
		failCompare[any](t, fmt.Sprintf(">= %v", min), actual, msgAndArgs...)
	}
}

// LessOrEqual checks if a value is less than or equal to a maximum.
// This complements our Greater function and is useful for range checks.
func LessOrEqual[T Ordered](t testing.TB, actual, max T, msgAndArgs ...any) {
	t.Helper()

	if actual > max {
		failCompare[any](t, fmt.Sprintf("<= %v", max), actual, msgAndArgs...)
	}
}

// Positive checks if a number is strictly greater than zero
func Positive[T Number](t testing.TB, value T, msgAndArgs ...any) {
	t.Helper()

	if value <= 0 {
		failCompare[any](t, "> 0", value, append([]any{"value is not positive"}, msgAndArgs...)...)
	}
}

// Negative checks if a number is strictly less than zero.
func Negative[T Number](t testing.TB, value T, msgAndArgs ...any) {
	t.Helper()

	if value >= 0 {
		failCompare[any](t, "< 0", value, append([]any{"value is not negative"}, msgAndArgs...)...)
	}
}

// InDelta checks if two numbers are close enough within a given delta.
// Useful for floating-point comparisons where exact equality is unreliable.
func InDelta[T Number](t testing.TB, expected, actual T, delta float64, msgAndArgs ...any) {
	t.Helper()

	diff := math.Abs(float64(expected) - float64(actual))
	if diff > delta {
		failCompare[any](t,
			fmt.Sprintf("%v (±%v)", expected, delta),
			actual,
			append(
				[]any{fmt.Sprintf("values differ by %v, exceeds delta %v", diff, delta)},
				msgAndArgs...)...)
	}
}
