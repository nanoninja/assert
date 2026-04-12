// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import "testing"

func TestContains(t *testing.T) {
	t.Run("string slice", func(t *testing.T) {
		tests := []struct {
			name      string
			slice     []string
			element   string
			wantError bool
		}{
			{
				name:      "element present",
				slice:     []string{"a", "b", "c"},
				element:   "b",
				wantError: false,
			},
			{
				name:      "element not present",
				slice:     []string{"a", "b", "c"},
				element:   "d",
				wantError: true,
			},
			{
				name:      "empty slice",
				slice:     []string{},
				element:   "a",
				wantError: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rec := NewTestRecorder(t)

				Contains(rec, tt.slice, tt.element)

				if tt.wantError != rec.HasError() {
					t.Errorf("Contains() error = %v, want %v", rec.HasError(), tt.wantError)
				}
			})
		}
	})

	t.Run("custom struct slice", func(t *testing.T) {
		type custom struct {
			value string
		}

		slice := []custom{{value: "a"}, {value: "b"}}
		rec := NewTestRecorder(t)

		Contains(rec, slice, custom{value: "a"})

		if rec.HasError() {
			t.Error("Contains() recorded error for present struct")
		}
	})
}

func TestEmpty(t *testing.T) {
	tests := []struct {
		name       string
		collection any
		wantError  bool
	}{
		{
			name:       "empty array",
			collection: [...]int{},
			wantError:  false,
		},
		{
			name:       "non-empty array",
			collection: [...]int{1},
			wantError:  true,
		},
		{
			name:       "empty slice",
			collection: []int{},
			wantError:  false,
		},
		{
			name:       "non-empty slice",
			collection: []int{1},
			wantError:  true,
		},
		{
			name:       "empty map",
			collection: map[string]int{},
			wantError:  false,
		},
		{
			name:       "non-empty map",
			collection: map[string]int{"a": 1},
			wantError:  true,
		},
		{
			name:       "empty string",
			collection: []int{},
			wantError:  false,
		},
		{
			name:       "non-empty string",
			collection: "hello",
			wantError:  true,
		},
		{
			name:       "invalid type",
			collection: 42,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Empty(rec, tt.collection)

			if tt.wantError != rec.HasError() {
				t.Errorf("Empty() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name       string
		collection any
		wantError  bool
	}{
		{
			name:       "non-empty slice",
			collection: []int{1, 2, 3},
			wantError:  false,
		},
		{
			name:       "non-empty map",
			collection: map[string]int{"a": 1},
			wantError:  false,
		},
		{
			name:       "non-empty string",
			collection: "hello",
			wantError:  false,
		},
		{
			name:       "empty slice",
			collection: []int{},
			wantError:  true,
		},
		{
			name:       "empty map",
			collection: map[string]int{},
			wantError:  true,
		},
		{
			name:       "empty string",
			collection: "",
			wantError:  true,
		},
		{
			name:       "invalid type",
			collection: 42,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			NotEmpty(rec, tt.collection)

			if tt.wantError != rec.HasError() {
				rec.Errorf("NotEmpty() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestHasKey(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]int
		key       string
		wantError bool
	}{
		{
			name:      "key exists",
			m:         map[string]int{"test": 1, "other": 2},
			key:       "test",
			wantError: false,
		},
		{
			name:      "key doesn't exist",
			m:         map[string]int{"test": 1, "other": 2},
			key:       "missing",
			wantError: true,
		},
		{
			name:      "empty map",
			m:         map[string]int{},
			key:       "test",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			HasKey(rec, tt.m, tt.key)

			if tt.wantError != rec.HasError() {
				t.Errorf("HasKey() error %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name       string
		collection any
		expected   int
		wantError  bool
	}{
		{
			name:       "correct slice length",
			collection: []int{1, 2, 3},
			expected:   3,
			wantError:  false,
		},
		{
			name:       "incorrect slice length",
			collection: []int{1, 2},
			expected:   3,
			wantError:  true,
		},
		{
			name:       "correct map length",
			collection: map[string]int{"a": 1, "b": 2},
			expected:   2,
			wantError:  false,
		},
		{
			name:       "correct string length",
			collection: "hello",
			expected:   5,
			wantError:  false,
		},
		{
			name:       "invalid type",
			collection: 42,
			expected:   0,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Len(rec, tt.collection, tt.expected)

			if tt.wantError != rec.HasError() {
				t.Errorf("Len() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestElementsMatch(t *testing.T) {
	tests := []struct {
		name      string
		expected  []int
		actual    []int
		wantError bool
	}{
		{
			name:      "same elements same order",
			expected:  []int{1, 2, 3},
			actual:    []int{1, 2, 3},
			wantError: false,
		},
		{
			name:      "same elements different order",
			expected:  []int{1, 2, 3},
			actual:    []int{3, 1, 2},
			wantError: false,
		},
		{
			name:      "different lengths",
			expected:  []int{1, 2, 3},
			actual:    []int{1, 2},
			wantError: true,
		},
		{
			name:      "different elements",
			expected:  []int{1, 2, 3},
			actual:    []int{1, 2, 4},
			wantError: true,
		},
		{
			name:      "duplicate elements handled correctly",
			expected:  []int{1, 1, 2},
			actual:    []int{1, 2, 2},
			wantError: true,
		},
		{
			name:      "empty slices",
			expected:  []int{},
			actual:    []int{},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			ElementsMatch(rec, tt.expected, tt.actual)

			if tt.wantError != rec.HasError() {
				t.Errorf("ElementsMatch() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestSubset(t *testing.T) {
	tests := []struct {
		name      string
		list      []int
		subset    []int
		wantError bool
	}{
		{
			name:      "subset contained in list",
			list:      []int{1, 2, 3, 4, 5},
			subset:    []int{2, 4},
			wantError: false,
		},
		{
			name:      "subset equals list",
			list:      []int{1, 2, 3},
			subset:    []int{1, 2, 3},
			wantError: false,
		},
		{
			name:      "element not in list",
			list:      []int{1, 2, 3},
			subset:    []int{2, 5},
			wantError: true,
		},
		{
			name:      "empty subset",
			list:      []int{1, 2, 3},
			subset:    []int{},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Subset(rec, tt.list, tt.subset)

			if tt.wantError != rec.HasError() {
				t.Errorf("Subset() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestNotSubset(t *testing.T) {
	tests := []struct {
		name      string
		list      []int
		subset    []int
		wantError bool
	}{
		{
			name:      "subset not contained in list",
			list:      []int{1, 2, 3},
			subset:    []int{2, 5},
			wantError: false,
		},
		{
			name:      "subset fully contained",
			list:      []int{1, 2, 3, 4, 5},
			subset:    []int{2, 4},
			wantError: true,
		},
		{
			name:      "empty subset is always contained",
			list:      []int{1, 2, 3},
			subset:    []int{},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			NotSubset(rec, tt.list, tt.subset)

			if tt.wantError != rec.HasError() {
				t.Errorf("NotSubset() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestNotContains(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		element   string
		msg       []string
		wantError bool
	}{
		{
			name:      "element not present in slice",
			slice:     []string{"a", "b", "c"},
			element:   "d",
			wantError: false,
		},
		{
			name:      "element present in slice",
			slice:     []string{"a", "b", "c"},
			element:   "b",
			wantError: true,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			element:   "a",
			wantError: false,
		},
		{
			name:      "element present with custom message",
			slice:     []string{"a", "b", "c"},
			element:   "c",
			msg:       []string{"should not contain c"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			NotContains(rec, tt.slice, tt.element, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("NotContains() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}
