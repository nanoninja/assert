// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"errors"
	"fmt"
	"testing"
)

// Equals checks if two values are equal using reflection.DeepEqual.
// It provides detailed error messages showing both values and their types when they differ.
func Equals[T any](t testing.TB, actual, expected T, msg ...string) {
	t.Helper()

	compare(t, expected, actual, msg...)
}

// Error checks if an error matches the expected error.
// It handles nil errors appropriately and provides clear error messages.
func Error(t testing.TB, actual, expected error) {
	t.Helper()

	if actual == nil && expected != nil {
		failCompare(t, expected, actual, "expected error but got nil")
	}
	if actual != nil && expected == nil {
		failCompare(t, expected, actual, "expected nil error")
	}
	if actual != expected {
		failCompare(t, expected, actual)
	}
}

// ErrorAs asserts that err can be converted to target type using errors.As.
// The target must be a pointer to an error type.
func ErrorAs(t testing.TB, err error, target any, msg ...string) {
	t.Helper()

	if !errors.As(err, target) {
		failCompare[any](t, err, fmt.Sprintf("error matching type %T", target), msg...)
	}
}

// ErrorIs asserts that err matches target using errors.Is.
// This is particularly useful when working with wrapped errors.
func ErrorIs(t testing.TB, err, target error, msg ...string) {
	t.Helper()

	if !errors.Is(err, target) {
		failCompare[any](t, err, fmt.Sprintf("error chain containing %v", target), msg...)
	}
}

// False asserts that a boolean value is false.
// It provides a clear error message with the source location and optional custom message.
func False(t testing.TB, value bool, msg ...string) {
	t.Helper()

	if value {
		// Using failCompare with any to handle mixed type comparisons in error message
		failCompare(t, false, value, msg...)
	}
}

// Nil checks if a value is nil, handling different types appropriately
// including interfaces, slices, maps, and pointers.
func Nil(t testing.TB, value any) {
	t.Helper()

	if !isNil(value) {
		failCompare(t, value, nil)
	}
}

// NotEquals asserts that two values are not equal.
// This is particularly useful when testing that a value has changed
// or that distinct objects remain separate.
func NotEquals[T any](t testing.TB, actual, expected T, msg ...string) {
	t.Helper()

	if isEqual(actual, expected) {
		failCompare(t,
			"values to be different",
			fmt.Sprintf("both values are equal: %v", actual),
			msg...,
		)
	}
}

// NotNil checks if a value is not nil, providing a clear error message
// when the value is unexpectedly nil.
func NotNil(t testing.TB, value any) {
	t.Helper()

	if isNil(value) {
		t.Error("\nexpected value to not be nil")
	}
}

// Panics verifies that a function panics with an expected message.
func Panics(t testing.TB, fn func(), expectedMsg string) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			actualMsg := fmt.Sprint(r)
			if actualMsg != expectedMsg {
				failCompare(t, expectedMsg, actualMsg, "unexpected panic message")
			}
		} else {
			t.Errorf("\nExpected panic: %v\n  Actual: no panic", expectedMsg)
		}
	}()

	fn()
}

// True asserts that a boolean value is true.
// It provides a clear error message with the source location and optional custom message.
func True(t testing.TB, value bool, msg ...string) {
	t.Helper()

	if !value {
		// Using failCompare with any to handle mixed type comparisons in error message
		failCompare(t, true, value, msg...)
	}
}
