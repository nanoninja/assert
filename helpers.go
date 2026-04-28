// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// TestRecorder wraps a testing.T instance and records error messages
// without failing the actual test. This allows us to verify assertion
// behaviors in our unit tests.
type TestRecorder struct {
	*testing.T
	errorCalled  bool
	errorMessage string
	helperCalled bool
}

// NewTestRecorder creates a new TestRecorder instance.
// It automatically marks itself as a test helper to maintain
// accurate line numbers in test output.
func NewTestRecorder(t *testing.T) *TestRecorder {
	t.Helper()

	return &TestRecorder{T: t}
}

// Error records that an error occurred with the given arguments
func (r *TestRecorder) Error(args ...any) {
	r.errorCalled = true
	r.errorMessage = fmt.Sprint(args...)
}

// Errorf records that an error occurred with the formatted message
func (r *TestRecorder) Errorf(format string, args ...any) {
	r.errorCalled = true
	r.errorMessage = fmt.Sprintf(format, args...)
}

// Helper records that Helper() was called, which is useful
// for verifying our assertions maintain proper stack traces
func (r *TestRecorder) Helper() {
	r.helperCalled = true
}

// HasError checks if any error method was called
func (r *TestRecorder) HasError() bool {
	return r.errorCalled
}

// ErrorMessage returns the recorded error message
func (r *TestRecorder) ErrorMessage() string {
	return r.errorMessage
}

// HelperCalled checks if Helper() was called
func (r *TestRecorder) HelperCalled() bool {
	return r.helperCalled
}

func compare[T any](t testing.TB, expected, actual T, msgAndArgs ...any) {
	t.Helper()

	if !isEqual(expected, actual) {
		failCompare(t, expected, actual, msgAndArgs...)
	}
}

func formatMessage(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	if msg, ok := msgAndArgs[0].(string); ok {
		if len(msgAndArgs) > 1 {
			return fmt.Sprintf(msg, msgAndArgs[1:]...)
		}
		return msg
	}
	return fmt.Sprint(msgAndArgs...)
}

// failCompare formats and outputs a detailed comparison error message for test failures.
// It includes the location of the failure, types of values being compared,
// and a clear visual representation of expected vs actual values.
func failCompare[T any](t testing.TB, expected, actual T, msgAndArgs ...any) {
	t.Helper()

	var builder strings.Builder

	if msg := formatMessage(msgAndArgs...); msg != "" {
		fmt.Fprintf(&builder, "\n Message: %s", msg)
	}

	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

	fmt.Fprintf(&builder, "\nExpected: (%v) %#v\n", expectedType, expected)
	fmt.Fprintf(&builder, "  Actual: (%v) %#v\n", actualType, actual)

	t.Error(builder.String())
}

// isEqual performs a generic equality check between two values of the same type.
// It uses reflection.DeepEqual to handle complex data structures correctly.
func isEqual[T any](x, y T) bool {
	return reflect.DeepEqual(x, y)
}

// isNil is a helper function that properly checks if a value is nil,
// handling special cases like interfaces and slices.
func isNil(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Ptr,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
