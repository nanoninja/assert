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

func TestLess(t *testing.T) {
	t.Run("numeric comparaison", func(t *testing.T) {
		tests := []struct {
			name      string
			e1        int
			e2        int
			wantError bool
		}{
			{
				name:      "less than maximum",
				e1:        3,
				e2:        5,
				wantError: false,
			},
			{
				name:      "equal to maximum",
				e1:        5,
				e2:        5,
				wantError: true,
			},
			{
				name:      "greater than maximum",
				e1:        10,
				e2:        5,
				wantError: true,
			},
		}

		for _, tt := range tests {
			rec := NewTestRecorder(t)

			Less(rec, tt.e1, tt.e2)

			if tt.wantError != rec.HasError() {
				t.Errorf("Less() error = %v, want %v", rec.HasError(), tt.wantError)
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

func TestPositive(t *testing.T) {
	tests := []struct {
		name      string
		value     int
		wantError bool
	}{
		{
			name:      "positive integer",
			value:     5,
			wantError: false,
		},
		{
			name:      "zero",
			value:     0,
			wantError: true,
		},
		{
			name:      "negative integer",
			value:     -3,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Positive(rec, tt.value)

			if tt.wantError != rec.HasError() {
				t.Errorf("Positive() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestPositive_Float(t *testing.T) {
	floatTests := []struct {
		name      string
		value     float64
		wantError bool
	}{
		{name: "positive float", value: 3.14, wantError: false},
		{name: "zero float", value: 0.0, wantError: true},
		{name: "negative float", value: -1.5, wantError: true},
	}

	for _, tt := range floatTests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Positive(rec, tt.value)

			if tt.wantError != rec.HasError() {
				t.Errorf("Positive() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestNegative(t *testing.T) {
	tests := []struct {
		name      string
		value     int
		wantError bool
	}{
		{
			name:      "negative integer",
			value:     -5,
			wantError: false,
		},
		{
			name:      "zero",
			value:     0,
			wantError: true,
		},
		{
			name:      "positive integer",
			value:     3,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Negative(rec, tt.value)

			if tt.wantError != rec.HasError() {
				t.Errorf("Negative() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestNegative_Float(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		wantError bool
	}{
		{name: "negative float", value: -3.14, wantError: false},
		{name: "zero float", value: 0.0, wantError: true},
		{name: "positive float", value: 1.5, wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			Negative(rec, tt.value)

			if tt.wantError != rec.HasError() {
				t.Errorf("Negative() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestInDelta(t *testing.T) {
	t.Run("float64 values", func(t *testing.T) {
		tests := []struct {
			name      string
			expected  float64
			actual    float64
			delta     float64
			wantError bool
		}{
			{
				name:      "within delta",
				expected:  0.3,
				actual:    0.1 + 0.2,
				delta:     0.0001,
				wantError: false,
			},
			{
				name:      "exactly equal",
				expected:  1.5,
				actual:    1.5,
				delta:     0.0,
				wantError: false,
			},
			{
				name:      "exceeds delta",
				expected:  1.0,
				actual:    1.5,
				delta:     0.1,
				wantError: true,
			},
			{
				name:      "negative values within delta",
				expected:  -1.0,
				actual:    -1.0001,
				delta:     0.001,
				wantError: false,
			},
			{
				name:      "negative values exceeds delta",
				expected:  -1.0,
				actual:    -2.0,
				delta:     0.5,
				wantError: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rec := NewTestRecorder(t)

				InDelta(rec, tt.expected, tt.actual, tt.delta)

				if tt.wantError != rec.HasError() {
					t.Errorf("InDelta() error = %v, want %v", rec.HasError(), tt.wantError)
				}
			})
		}
	})

	t.Run("integer values", func(t *testing.T) {
		tests := []struct {
			name      string
			expected  int
			actual    int
			delta     float64
			wantError bool
		}{
			{
				name:      "within delta",
				expected:  10,
				actual:    11,
				delta:     1.0,
				wantError: false,
			},
			{
				name:      "exceeds delta",
				expected:  10,
				actual:    15,
				delta:     1.0,
				wantError: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rec := NewTestRecorder(t)

				InDelta(rec, tt.expected, tt.actual, tt.delta)

				if tt.wantError != rec.HasError() {
					t.Errorf("InDelta() error = %v, want %v", rec.HasError(), tt.wantError)
				}
			})
		}
	})
}
