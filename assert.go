// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"errors"
	"fmt"
	"testing"
)

// Equal checks if two values are equal using reflection.DeepEqual.
// Parameters follow the same convention as testify: expected first, actual second.
func Equal[T any](t testing.TB, expected, actual T, msg ...string) {
	t.Helper()

	compare(t, expected, actual, msg...)
}

// EqualError checks if an error's message matches the expected string.
// It handles nil errors appropriately and provides clear error messages.
func EqualError(t testing.TB, err error, errString string, msg ...string) {
	t.Helper()

	if err == nil {
		failCompare[any](t, errString, nil, append([]string{"expected an error"}, msg...)...)
		return
	}
	if err.Error() != errString {
		failCompare(t, errString, err.Error(), msg...)
	}
}

// Error asserts that an error occurred (i.e., the error is not nil).
func Error(t testing.TB, err error, msg ...string) {
	t.Helper()

	if isNil(err) {
		failCompare[any](t, "non-nil error", nil, append([]string{"expected an error"}, msg...)...)
	}
}

// ErrorAs asserts that err can be converted to target type using errors.As.
func ErrorAs(t testing.TB, err error, target any, msg ...string) {
	t.Helper()

	if !errors.As(err, target) {
		failCompare[any](t, fmt.Sprintf("error matching type %T", target), err, msg...)
	}
}

// ErrorIs asserts that err matches target using errors.Is.
func ErrorIs(t testing.TB, err, target error, msg ...string) {
	t.Helper()

	if !errors.Is(err, target) {
		failCompare[any](t, fmt.Sprintf("error chain containing %v", target), err, msg...)
	}
}

// False asserts that a boolean value is false.
func False(t testing.TB, value bool, msg ...string) {
	t.Helper()

	if value {
		failCompare(t, false, value, msg...)
	}
}

// Nil checks if a value is nil.
func Nil(t testing.TB, value any, msg ...string) {
	t.Helper()

	if !isNil(value) {
		failCompare(t, nil, value, msg...)
	}
}

// NoError asserts that no error occurred (i.e., the error is nil).
func NoError(t testing.TB, err error, msg ...string) {
	t.Helper()

	if !isNil(err) {
		failCompare[any](t, nil, err, append([]string{"unexpected error"}, msg...)...)
	}
}

// NotEqual asserts that two values are not equal.
// Parameters follow the same convention as testify: expected first, actual second.
func NotEqual[T any](t testing.TB, expected, actual T, msg ...string) {
	t.Helper()

	if isEqual(expected, actual) {
		failCompare(t,
			"values to be different",
			fmt.Sprintf("both values are equal: %v", actual),
			msg...,
		)
	}
}

// NotNil checks if a value is not nil.
func NotNil(t testing.TB, value any, msg ...string) {
	t.Helper()

	if isNil(value) {
		failCompare[any](t, "non-nil value", nil, msg...)
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
func True(t testing.TB, value bool, msg ...string) {
	t.Helper()

	if !value {
		failCompare(t, true, value, msg...)
	}
}
