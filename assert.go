// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// Equal checks if two values are equal using reflection.DeepEqual.
// Parameters follow the same convention as testify: expected first, actual second.
func Equal[T any](t testing.TB, expected, actual T, msgAndArgs ...any) {
	t.Helper()

	compare(t, expected, actual, msgAndArgs...)
}

// EqualError checks if an error's message matches the expected string.
// It handles nil errors appropriately and provides clear error messages.
func EqualError(t testing.TB, err error, errString string, msgAndArgs ...any) {
	t.Helper()

	if err == nil {
		failCompare[any](t, errString, nil, append([]any{"expected an error"}, msgAndArgs...)...)
		return
	}
	if err.Error() != errString {
		failCompare(t, errString, err.Error(), msgAndArgs...)
	}
}

// Error asserts that an error occurred (i.e., the error is not nil).
func Error(t testing.TB, err error, msgAndArgs ...any) {
	t.Helper()

	if isNil(err) {
		failCompare[any](t, "non-nil error", nil, append([]any{"expected an error"}, msgAndArgs...)...)
	}
}

// ErrorAs asserts that err can be converted to target type using errors.As.
func ErrorAs(t testing.TB, err error, target any, msgAndArgs ...any) {
	t.Helper()

	if !errors.As(err, target) {
		failCompare[any](t, fmt.Sprintf("error matching type %T", target), err, msgAndArgs...)
	}
}

// ErrorIs asserts that err matches target using errors.Is.
func ErrorIs(t testing.TB, err, target error, msgAndArgs ...any) {
	t.Helper()

	if !errors.Is(err, target) {
		failCompare[any](t, fmt.Sprintf("error chain containing %v", target), err, msgAndArgs...)
	}
}

// False asserts that a boolean value is false.
func False(t testing.TB, value bool, msgAndArgs ...any) {
	t.Helper()

	if value {
		failCompare(t, false, value, msgAndArgs...)
	}
}

// Nil checks if a value is nil.
func Nil(t testing.TB, value any, msgAndArgs ...any) {
	t.Helper()

	if !isNil(value) {
		failCompare(t, nil, value, msgAndArgs...)
	}
}

// NoError asserts that no error occurred (i.e., the error is nil).
func NoError(t testing.TB, err error, msgAndArgs ...any) {
	t.Helper()

	if !isNil(err) {
		failCompare[any](t, nil, err, append([]any{"unexpected error"}, msgAndArgs...)...)
	}
}

// NotEqual asserts that two values are not equal.
// Parameters follow the same convention as testify: expected first, actual second.
func NotEqual[T any](t testing.TB, expected, actual T, msgAndArgs ...any) {
	t.Helper()

	if isEqual(expected, actual) {
		failCompare(t,
			"values to be different",
			fmt.Sprintf("both values are equal: %v", actual),
			msgAndArgs...,
		)
	}
}

// NotNil checks if a value is not nil.
func NotNil(t testing.TB, value any, msgAndArgs ...any) {
	t.Helper()

	if isNil(value) {
		failCompare[any](t, "non-nil value", nil, msgAndArgs...)
	}
}

// IsType asserts that the actual value is of the same type as the expected value.
func IsType(t testing.TB, expected, actual any, msgAndArgs ...any) {
	t.Helper()

	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

	if expectedType != actualType {
		failCompare(t, expectedType, actualType,
			append([]any{"unexpected type"}, msgAndArgs...)...)
	}
}

// Implements asserts that a value implements a given interface.
// The iface parameter must be a pointer to an interface type: (*MyInterface)(nil).
//
// Example:
//
//	assert.Implements(t, (*io.Reader)(nil), myReader)
func Implements(t testing.TB, iface, value any, msgAndArgs ...any) {
	t.Helper()

	ifaceType := reflect.TypeOf(iface).Elem()
	valueType := reflect.TypeOf(value)

	if !valueType.Implements(ifaceType) {
		failCompare[any](t,
			fmt.Sprintf("value implementing %s", ifaceType),
			valueType,
			append([]any{"value does not implement interface"}, msgAndArgs...)...)
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

// NotPanics verifies that a function does not panic.
func NotPanics(t testing.TB, fn func(), msgAndArgs ...any) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			failCompare[any](t,
				"no panic",
				fmt.Sprint(r),
				append([]any{"unexpected panic"}, msgAndArgs...)...,
			)
		}
	}()

	fn()
}

// PanicsWithError verifies that a function panics with an error matching the expected message.
func PanicsWithError(t testing.TB, fn func(), expectedMsg string) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				failCompare[any](t, expectedMsg, fmt.Sprint(r), "panic value is not an error")
				return
			}
			if err.Error() != expectedMsg {
				failCompare(t, expectedMsg, err.Error(), "unexpected panic error message")
			}
		} else {
			t.Errorf("\nExpected panic: %v\n  Actual: no panic", expectedMsg)
		}
	}()

	fn()
}

// True asserts that a boolean value is true.
func True(t testing.TB, value bool, msgAndArgs ...any) {
	t.Helper()

	if !value {
		failCompare(t, true, value, msgAndArgs...)
	}
}
