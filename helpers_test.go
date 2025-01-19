// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Run("compare with equal values", func(t *testing.T) {
		rec := NewTestRecorder(t)

		compare(rec, 42, 42)

		if rec.HasError() {
			t.Error("compare() recorded error for equal values")
		}

		if !rec.HelperCalled() {
			t.Error("Helper() was not called")
		}
	})

	t.Run("compare with different values should call failCompare", func(t *testing.T) {
		rec := NewTestRecorder(t)
		msg := "optional message"

		compare(rec, 42, 43, msg)

		if !rec.HasError() {
			t.Error("compare() did not record error for different values")
		}

		errorMsg := rec.ErrorMessage()

		// Check essential parts of the error message matching the actual output format
		expectedParts := []string{
			"Message: optional message", // Should match exactly how failCompare formats it
			"Expected: (int) 42",        // Order matches failCompare implementation
			"Actual: (int) 43",          // Order matches failCompare implementation
		}

		for _, part := range expectedParts {
			if !strings.Contains(errorMsg, part) {
				t.Errorf("incorrect error message\nwant: %q\nin: %q", part, errorMsg)
			}
		}
	})
}

func TestFailCompare(t *testing.T) {
	tests := []struct {
		name      string
		actual    any
		expected  any
		msg       []string
		wantParts []string
	}{
		{
			name:     "basic types with message",
			actual:   42,
			expected: 43,
			msg:      []string{"test message"},
			wantParts: []string{
				"Message: test message",
				"Actual: (int) 42",
				"Expected: (int) 43",
			},
		},
		{
			name:     "without message",
			actual:   "hello",
			expected: "world",
			msg:      nil,
			wantParts: []string{
				"Actual: (string) \"hello\"",
				"Expected: (string) \"world\"",
			},
		},
		{
			name:     "complex types",
			actual:   []int{1, 2},
			expected: []int{3, 4},
			msg:      []string{"slice comparison"},
			wantParts: []string{
				"Message: slice comparison",
				"Actual: ([]int) []int{1, 2}",
				"Expected: ([]int) []int{3, 4}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			failCompare(rec, tt.actual, tt.expected, tt.msg...)

			if !rec.HasError() {
				t.Error("failCompare() did not record error")
			}

			errorMsg := rec.ErrorMessage()

			for _, part := range tt.wantParts {
				if !strings.Contains(errorMsg, part) {
					t.Errorf("failCompare() message missing %q\ngot: %s", part, errorMsg)
				}
			}
		})
	}
}

func TestIsEqual(t *testing.T) {
	tests := []struct {
		name string
		x    any
		y    any
		want bool
	}{
		{
			name: "same integers",
			x:    42,
			y:    42,
			want: true,
		},
		{
			name: "different integers",
			x:    42,
			y:    43,
			want: false,
		},
		{
			name: "same strings",
			x:    "test",
			y:    "test",
			want: true,
		},
		{
			name: "complex structs same",
			x:    struct{ A []string }{A: []string{"a", "b"}},
			y:    struct{ A []string }{A: []string{"a", "b"}},
			want: true,
		},
		{
			name: "complex structs different",
			x:    struct{ A []string }{A: []string{"a", "b"}},
			y:    struct{ A []string }{A: []string{"b", "a"}},
			want: false,
		},
		{
			name: "nil values",
			x:    nil,
			y:    nil,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.x, tt.y); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	// Prepare test values
	var nilPtr *string
	var nilSlice []string
	var nilMap map[string]int
	var nilInterface interface{}
	nonNilPtr := new(string)

	tests := []struct {
		name  string
		value any
		want  bool
	}{
		{
			name:  "nil value",
			value: nil,
			want:  true,
		},
		{
			name:  "nil pointer",
			value: nilPtr,
			want:  true,
		},
		{
			name:  "non-nil pointer",
			value: nonNilPtr,
			want:  false,
		},
		{
			name:  "nil slice",
			value: nilSlice,
			want:  true,
		},
		{
			name:  "empty slice",
			value: []string{},
			want:  false,
		},
		{
			name:  "nil map",
			value: nilMap,
			want:  true,
		},
		{
			name:  "empty map",
			value: map[string]int{},
			want:  false,
		},
		{
			name:  "nil interface",
			value: nilInterface,
			want:  true,
		},
		{
			name:  "non-nil basic type",
			value: 42,
			want:  false,
		},
		{
			name:  "non-nil string",
			value: "test",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNil(tt.value); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
