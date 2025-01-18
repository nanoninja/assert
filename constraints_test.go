// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import "testing"

func TestBetween(t *testing.T) {
	t.Run("numeric values", func(t *testing.T) {
		tests := []struct {
			name      string
			actual    int
			min       int
			max       int
			wantError bool
		}{
			{
				name:      "value within bound",
				actual:    5,
				min:       1,
				max:       10,
				wantError: false,
			},
			{
				name:      "value at upper bound",
				actual:    10,
				min:       1,
				max:       10,
				wantError: false,
			},
			{
				name:      "value below bounds",
				actual:    0,
				min:       1,
				max:       10,
				wantError: true,
			},
			{
				name:      "value above bounds",
				actual:    11,
				min:       1,
				max:       10,
				wantError: true,
			},
		}

		for _, tt := range tests {
			rec := NewTestRecorder(t)

			Between(rec, tt.actual, tt.min, tt.max)

			if tt.wantError != rec.HasError() {
				t.Errorf("Between() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		}
	})

	t.Run("strings values", func(t *testing.T) {
		rec := NewTestRecorder(t)

		Between(rec, "b", "a", "c")

		if rec.HasError() {
			t.Error("Between() recorded error for value within string bounds")
		}
	})
}

func TestGreater(t *testing.T) {
	t.Run("numeric comparisons", func(t *testing.T) {
		tests := []struct {
			name      string
			actual    int
			min       int
			wantError bool
		}{
			{
				name:      "greater than minimum",
				actual:    10,
				min:       5,
				wantError: false,
			},
			{
				name:      "equal to minimum",
				actual:    5,
				min:       5,
				wantError: true,
			},
			{
				name:      "less than minimum",
				actual:    3,
				min:       5,
				wantError: true,
			},
		}

		for _, tt := range tests {
			rec := NewTestRecorder(t)

			Greater(rec, tt.actual, tt.min)

			if tt.wantError != rec.HasError() {
				t.Errorf("Greater() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		}
	})
}

func TestGreaterOrEqual(t *testing.T) {
	tests := []struct {
		name      string
		actual    int
		min       int
		msg       []string
		wantError bool
	}{
		{
			name:      "value greater than min",
			actual:    15,
			min:       10,
			wantError: false,
		},
		{
			name:      "value equal to min",
			actual:    10,
			min:       10,
			wantError: false,
		},
		{
			name:      "value less than min",
			actual:    5,
			min:       10,
			wantError: true,
		},
		{
			name:      "zero values",
			actual:    0,
			min:       0,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			GreaterOrEqual(rec, tt.actual, tt.min, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("GreaterOrEqual() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestLessOrEqual(t *testing.T) {
	tests := []struct {
		name      string
		actual    int
		max       int
		msg       []string
		wantError bool
	}{
		{
			name:      "value less than max",
			actual:    5,
			max:       10,
			wantError: false,
		},
		{
			name:      "value equal to max",
			actual:    10,
			max:       10,
			wantError: false,
		},
		{
			name:      "value greater than max",
			actual:    15,
			max:       10,
			wantError: true,
		},
		{
			name:      "zero values",
			actual:    0,
			max:       0,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			LessOrEqual(rec, tt.actual, tt.max, tt.msg...)

			if tt.wantError != rec.HasError() {
				t.Errorf("LessOrEqual() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}
