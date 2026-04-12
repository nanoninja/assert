// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

// HasPrefix checks if a string starts with an expected prefix.
// Useful for testing string formatting, paths, or URLs.
func HasPrefix(t testing.TB, s, prefix string, msg ...string) {
	t.Helper()

	if !strings.HasPrefix(s, prefix) {
		failCompare(t, s, fmt.Sprintf("should start with %q", prefix), msg...)
	}
}

// HasSuffix checks if a string ends with an expected suffix.
// Useful for testing file extensions, domains, etc.
func HasSuffix(t testing.TB, s, suffix string, msg ...string) {
	t.Helper()

	if !strings.HasSuffix(s, suffix) {
		failCompare(t, s, fmt.Sprintf("should end with %q", suffix), msg...)
	}
}

// MatchRegexp checks if a string matches a regular expression pattern.
// Powerful for testing string patterns and formats.
func MatchRegexp(t testing.TB, s, pattern string, msg ...string) {
	t.Helper()

	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		failCompare(t, pattern, "valid regexp pattern",
			append([]string{fmt.Sprintf("invalid regexp: %v", err)}, msg...)...)
		return
	}

	if !matched {
		failCompare(t, s, fmt.Sprintf("should match pattern %q", pattern), msg...)
	}
}

// StringContains checks if a string contains an expected substring.
func StringContains(t testing.TB, s, substr string, msg ...string) {
	t.Helper()

	if !strings.Contains(s, substr) {
		failCompare(t, substr, s, append(
			[]string{"string does not contain expected substring"},
			msg...)...)
	}
}
