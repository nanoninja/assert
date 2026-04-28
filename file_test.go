// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	dir := t.TempDir()

	f, err := os.CreateTemp(dir, "testfile-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		path      string
		wantError bool
	}{
		{
			name:      "file exists",
			path:      f.Name(),
			wantError: false,
		},
		{
			name:      "file does not exist",
			path:      dir + "/nonexistent.txt",
			wantError: true,
		},
		{
			name:      "path is a directory",
			path:      dir,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			FileExists(rec, tt.path)

			if tt.wantError != rec.HasError() {
				t.Errorf("FileExists() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}

func TestDirExists(t *testing.T) {
	dir := t.TempDir()

	f, err := os.CreateTemp(dir, "testfile-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		path      string
		wantError bool
	}{
		{
			name:      "directory exists",
			path:      dir,
			wantError: false,
		},
		{
			name:      "directory does not exist",
			path:      dir + "/nonexistent",
			wantError: true,
		},
		{
			name:      "path is a file",
			path:      f.Name(),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := NewTestRecorder(t)

			DirExists(rec, tt.path)

			if tt.wantError != rec.HasError() {
				t.Errorf("DirExists() error = %v, want %v", rec.HasError(), tt.wantError)
			}
		})
	}
}
