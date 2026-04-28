// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"os"
	"testing"
)

// FileExists asserts that a file exists at the given path.
func FileExists(t testing.TB, path string, msgAndArgs ...any) {
	t.Helper()

	info, err := os.Stat(path)
	if os.IsNotExist(err) || (err == nil && info.IsDir()) {
		failCompare(t, path, "no file at path",
			append([]any{"expected file does not exist"}, msgAndArgs...)...)
	}
}

// DirExists asserts that a directory exists at the given path.
func DirExists(t testing.TB, path string, msgAndArgs ...any) {
	t.Helper()

	info, err := os.Stat(path)
	if os.IsNotExist(err) || (err == nil && !info.IsDir()) {
		failCompare(t, path, "no directory at path",
			append([]any{"expected directory does not exist"}, msgAndArgs...)...)
	}
}
