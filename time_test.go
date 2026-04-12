// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"testing"
	"time"
)

func TestWithinDuration(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		expected  time.Time
		actual    time.Time
		delta     time.Duration
		wantError bool
	}{
		{
			name:      "same time",
			expected:  now,
			actual:    now,
			delta:     time.Second,
			wantError: false,
		},
		{
			name:      "within delta",
			expected:  now,
			actual:    now.Add(500 * time.Millisecond),
			delta:     time.Second,
			wantError: false,
		},
		{
			name:      "exactly at delta boundary",
			expected:  now,
			actual:    now.Add(time.Second),
			delta:     time.Second,
			wantError: false,
		},
		{
			name:      "exceeds delta",
			expected:  now,
			actual:    now.Add(2 * time.Second),
			delta:     time.Second,
			wantError: true,
		},
		{
			name:      "actual before expected within delta",
			expected:  now,
			actual:    now.Add(-500 * time.Millisecond),
			delta:     time.Second,
			wantError: false,
		},
		{
			name:      "actual before expected exceeds delta",
			expected:  now,
			actual:    now.Add(-2 * time.Second),
			delta:     time.Second,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			WithinDuration(rec, tt.expected, tt.actual, tt.delta)

			if tt.wantError != rec.HasError() {
				t.Errorf("WithinDuration() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}
