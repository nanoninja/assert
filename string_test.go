// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import "testing"

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		prefix    string
		msg       []any
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
		msg       []any
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

func TestMatchRegexp(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		pattern   string
		msg       []any
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
			name:      "invalid pattern",
			s:         "test",
			pattern:   "[",
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
