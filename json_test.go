// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import "testing"

func TestEqualJSON(t *testing.T) {
	tests := []struct {
		name      string
		expected  any
		actual    string
		wantError bool
	}{
		{
			name:      "struct matches JSON",
			expected:  struct{ Name string }{Name: "alice"},
			actual:    `{"Name":"alice"}`,
			wantError: false,
		},
		{
			name:      "map matches JSON",
			expected:  map[string]int{"a": 1, "b": 2},
			actual:    `{"b":2,"a":1}`,
			wantError: false,
		},
		{
			name:      "value mismatch",
			expected:  map[string]int{"a": 1},
			actual:    `{"a":2}`,
			wantError: true,
		},
		{
			name:      "invalid actual JSON",
			expected:  map[string]int{"a": 1},
			actual:    `{invalid}`,
			wantError: true,
		},
		{
			name:      "unmarshalable expected value",
			expected:  make(chan int),
			actual:    `{}`,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			EqualJSON(rec, tt.expected, tt.actual)

			if tt.wantError != rec.HasError() {
				t.Errorf("EqualJSON() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestJSONEqual(t *testing.T) {
	tests := []struct {
		name      string
		expected  string
		actual    string
		wantError bool
	}{
		{
			name:      "identical JSON",
			expected:  `{"name":"alice","age":30}`,
			actual:    `{"name":"alice","age":30}`,
			wantError: false,
		},
		{
			name:      "same content different key order",
			expected:  `{"a":1,"b":2}`,
			actual:    `{"b":2,"a":1}`,
			wantError: false,
		},
		{
			name:      "same content different whitespace",
			expected:  `{"a": 1}`,
			actual:    `{"a":1}`,
			wantError: false,
		},
		{
			name:      "different values",
			expected:  `{"a":1}`,
			actual:    `{"a":2}`,
			wantError: true,
		},
		{
			name:      "different keys",
			expected:  `{"a":1}`,
			actual:    `{"b":1}`,
			wantError: true,
		},
		{
			name:      "JSON arrays same order",
			expected:  `[1,2,3]`,
			actual:    `[1,2,3]`,
			wantError: false,
		},
		{
			name:      "JSON arrays different order",
			expected:  `[1,2,3]`,
			actual:    `[3,2,1]`,
			wantError: true,
		},
		{
			name:      "invalid expected JSON",
			expected:  `{invalid}`,
			actual:    `{"a":1}`,
			wantError: true,
		},
		{
			name:      "invalid actual JSON",
			expected:  `{"a":1}`,
			actual:    `{invalid}`,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			JSONEqual(rec, tt.expected, tt.actual)

			if tt.wantError != rec.HasError() {
				t.Errorf("JSONEqual() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}
