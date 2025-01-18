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

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		prefix    string
		msg       []string
		wantError bool
	}{
		{
			name:      "string has prefix",
			s:         "hello world",
			prefix:    "hello",
			wantError: false,
		},
		{
			name:      "string does not have prefix",
			s:         "hello world",
			prefix:    "world",
			wantError: true,
		},
		{
			name:      "empty string",
			s:         "",
			prefix:    "test",
			wantError: true,
		},
		{
			name:      "empty prefix",
			s:         "hello",
			prefix:    "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			HasPrefix(rec, tt.s, tt.prefix, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("HasPrefix() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		suffix    string
		msg       []string
		wantError bool
	}{
		{
			name:      "string has suffix",
			s:         "hello world",
			suffix:    "world",
			wantError: false,
		},
		{
			name:      "string does not have suffix",
			s:         "hello world",
			suffix:    "hello",
			wantError: true,
		},
		{
			name:      "empty string",
			s:         "",
			suffix:    "test",
			wantError: true,
		},
		{
			name:      "empty suffix",
			s:         "hello",
			suffix:    "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			HasSuffix(rec, tt.s, tt.suffix, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("HasSuffix() error = %v, want %v", rec.HasError(), tt.wantError)
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

func TestMatchRegexp(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		pattern   string
		msg       []string
		wantError bool
	}{
		{
			name:      "string matches pattern",
			s:         "hello123",
			pattern:   `^[a-z]+\d+$`,
			wantError: false,
		},
		{
			name:      "string does not match pattern",
			s:         "hello",
			pattern:   `^\d+$`,
			wantError: true,
		},
		{
			name:      "empty string with pattern",
			s:         "",
			pattern:   `^$`,
			wantError: false,
		},
		{
			name:      "invalid pattern should panic",
			s:         "test",
			pattern:   "[", // Invalid regexp pattern
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			MatchRegexp(rec, tt.s, tt.pattern, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("MatchRegexp() error = %v, want %v", rec.HasError(), tt.wantError)
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

func TestStringContains(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		substr    string
		wantError bool
	}{
		{
			name:      "substring present",
			s:         "hello world",
			substr:    "world",
			wantError: false,
		},
		{
			name:      "substring not present",
			s:         "hello world",
			substr:    "golang",
			wantError: true,
		},
		{
			name:      "empty string",
			s:         "",
			substr:    "test",
			wantError: true,
		},
		{
			name:      "empty substring",
			s:         "hello world",
			substr:    "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			StringContains(rec, tt.s, tt.substr)

			if tt.wantError != rec.HasError() {
				t.Errorf("StringContains() error %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}
