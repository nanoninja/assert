// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"errors"
	"fmt"
	"testing"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		name      string
		actual    any
		expected  any
		msg       []string
		wantError bool
	}{
		{
			name:      "same boolean",
			actual:    true,
			expected:  true,
			wantError: false,
		},
		{
			name:      "same integer",
			actual:    10,
			expected:  10,
			wantError: false,
		},
		{
			name:      "same empty struct",
			actual:    struct{}{},
			expected:  struct{}{},
			wantError: false,
		},
		{
			name:      "same nil values",
			actual:    nil,
			expected:  nil,
			wantError: false,
		},
		{
			name:      "same string different quotes",
			actual:    `test`,
			expected:  "test",
			wantError: false,
		},
		{
			name:      "same float",
			actual:    0.5405278,
			expected:  0.5405278,
			wantError: false,
		},
		{
			name:      "same error message",
			actual:    errors.New("error"),
			expected:  errors.New("error"),
			wantError: false,
		},
		{
			name:      "same slice",
			actual:    []string{"a", "b", "c"},
			expected:  []string{"a", "b", "c"},
			wantError: false,
		},
		{
			name:      "same map",
			actual:    map[int]string{1: "x", 2: "y"},
			expected:  map[int]string{1: "x", 2: "y"},
			wantError: false,
		},
		{
			name:      "different values",
			actual:    42,
			expected:  43,
			wantError: true,
		},
		{
			name:      "different types",
			actual:    42,
			expected:  "42",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Equals(rec, tt.actual, tt.expected, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("Equals() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestEqualError(t *testing.T) {
	errOne := errors.New("error one")
	errTwo := errors.New("error two")

	tests := []struct {
		name      string
		actual    error
		expected  error
		wantError bool
	}{
		{
			name:      "same errors",
			actual:    errOne,
			expected:  errOne,
			wantError: false,
		},
		{
			name:      "different errors",
			actual:    errOne,
			expected:  errTwo,
			wantError: true,
		},
		{
			name:      "actual is nil",
			actual:    nil,
			expected:  errOne,
			wantError: true,
		},
		{
			name:      "expected is nil",
			actual:    errOne,
			expected:  nil,
			wantError: true,
		},
		{
			name:      "both nil",
			actual:    nil,
			expected:  nil,
			wantError: false,
		},
	}

	for _, tt := range tests {
		rec := NewTestRecorder(t)

		EqualError(rec, tt.actual, tt.expected)

		if tt.wantError != rec.HasError() {
			t.Errorf("Error() error %v, want %v", rec.HasError(), tt.wantError)
		}
	}
}

// Custom error types.
type testError struct {
	code int
}

func (e *testError) Error() string {
	return fmt.Sprintf("error code: %d", e.code)
}

func TestErrorAs(t *testing.T) {
	baseErr := &testError{code: 404}
	wrappedErr := fmt.Errorf("wrapped: %w", baseErr)

	tests := []struct {
		name      string
		err       error
		wantError bool
	}{
		{
			name:      "direct error match",
			err:       baseErr,
			wantError: false,
		},
		{
			name:      "wrapped error match",
			err:       wrappedErr,
			wantError: false,
		},
		{
			name:      "no match with basic error",
			err:       errors.New("basic error"),
			wantError: true,
		},
		{
			name:      "nil error",
			err:       nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)
			var target *testError

			ErrorAs(rec, tt.err, &target)

			if tt.wantError != rec.HasError() {
				t.Errorf("ErrorAs() error = %v, want %v", rec.HasError(), tt.wantError)
			}

			if !tt.wantError && target.code != 404 {
				t.Errorf("ErrorAs() extracted wrong value = %v, want %v", target.code, 404)
			}
		})
	}
}

func TestErrorIs(t *testing.T) {
	baseErr := errors.New("base error")
	wrappedErr := fmt.Errorf("wrapped: %w", baseErr)

	tests := []struct {
		name      string
		err       error
		target    error
		msg       []string
		wantError bool
	}{
		{
			name:      "direct error match",
			err:       baseErr,
			target:    baseErr,
			wantError: false,
		},
		{
			name:      "wrapped error match",
			err:       wrappedErr,
			target:    baseErr,
			wantError: false,
		},
		{
			name:      "no match",
			err:       baseErr,
			target:    errors.New("different error"),
			wantError: true,
		},
		{
			name:      "nil error",
			err:       nil,
			target:    baseErr,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			ErrorIs(rec, tt.err, tt.target, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("ErrorIs() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestFalse(t *testing.T) {
	tests := []struct {
		name      string
		value     bool
		msg       []string
		wantError bool
	}{
		{
			name:      "false value passes",
			value:     false,
			wantError: false,
		},
		{
			name:      "true value fails",
			value:     true,
			wantError: true,
		},
		{
			name:      "true value fails with message",
			value:     true,
			msg:       []string{"custom failure message"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			False(rec, tt.value, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("False() error = %v, want %v", rec.HasError(), tt.wantError)
			}

			if tt.wantError && !rec.HelperCalled() {
				t.Error("Helper() was not called")
			}
		})
	}
}

func TestNil(t *testing.T) {
	// Define test values
	var nilPointer *string
	var nilSlice []string
	var nilMap map[string]int
	nonNilValue := "test"

	tests := []struct {
		name      string
		value     any
		wantError bool
	}{
		{
			name:      "nil value",
			value:     nil,
			wantError: false,
		},
		{
			name:      "nil pointer",
			value:     nilPointer,
			wantError: false,
		},
		{
			name:      "nil slice",
			value:     nilSlice,
			wantError: false,
		},
		{
			name:      "nil map",
			value:     nilMap,
			wantError: false,
		},
		{
			name:      "non-nil value",
			value:     nonNilValue,
			wantError: true,
		},
		{
			name:      "empty slice",
			value:     []string{},
			wantError: true,
		},
		{
			name:      "empty map",
			value:     map[string]int{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Nil(rec, tt.value)

			if tt.wantError != rec.HasError() {
				t.Errorf("Nil() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestNotEquals(t *testing.T) {
	tests := []struct {
		name      string
		actual    any
		expected  any
		msg       []string
		wantError bool
	}{
		{
			name:      "different integers",
			actual:    42,
			expected:  43,
			wantError: false,
		},
		{
			name:      "same integer",
			actual:    42,
			expected:  42,
			wantError: true,
		},
		{
			name:      "different strings with message",
			actual:    "hello",
			expected:  "world",
			msg:       []string{"string should differ"},
			wantError: false,
		},
		{
			name:      "same string with message",
			actual:    "hello",
			expected:  "hello",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			NotEquals(rec, tt.actual, tt.expected, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("NotEquals() error %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestNotNil(t *testing.T) {
	// Define some test values
	var nilPointer *string
	var nilSlice []string
	var nilMap map[string]int

	nonNilValue := "test"

	tests := []struct {
		name      string
		value     any
		wantError bool
	}{
		{
			name:      "nil value",
			value:     nil,
			wantError: true,
		},
		{
			name:      "nil pointer",
			value:     nilPointer,
			wantError: true,
		},
		{
			name:      "nil slice",
			value:     nilSlice,
			wantError: true,
		},
		{
			name:      "nil map",
			value:     nilMap,
			wantError: true,
		},
		{
			name:      "non-nil value",
			value:     nonNilValue,
			wantError: false,
		},
		{
			name:      "empty slice",
			value:     []string{},
			wantError: false,
		},
		{
			name:      "empty map",
			value:     map[string]int{},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			NotNil(rec, tt.value)

			if tt.wantError != rec.HasError() {
				t.Errorf("NotNil() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestPanics(t *testing.T) {
	tests := []struct {
		name      string
		fn        func()
		expectMsg string
		wantError bool
	}{
		{
			name: "matching panic message",
			fn: func() {
				panic("expected panic")
			},
			expectMsg: "expected panic",
			wantError: false,
		},
		{
			name: "different panic message",
			fn: func() {
				panic("unexpected panic")
			},
			expectMsg: "expected panic",
			wantError: true,
		},
		{
			name: "no panic when expected",
			fn: func() {
				// Function that doesn't panic
			},
			expectMsg: "should panic",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Panics(rec, tt.fn, tt.expectMsg)

			if tt.wantError != rec.HasError() {
				t.Errorf("Panics() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestTrue(t *testing.T) {
	tests := []struct {
		name      string
		value     bool
		msg       []string
		wantError bool
	}{
		{
			name:      "true value passes",
			value:     true,
			wantError: false,
		},
		{
			name:      "false value fails",
			value:     false,
			wantError: true,
		},
		{
			name:      "false value fails with message",
			value:     false,
			msg:       []string{"custom failure message"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			True(rec, tt.value, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("True() error = %v, want %v", rec.HasError(), tt.wantError)
			}

			if tt.wantError && !rec.HelperCalled() {
				t.Error("Helper() was not called")
			}
		})
	}
}
