// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

// Contains checks if a slice contains a specific element.
// The comparison is done using reflection.DeepEqual.
func Contains[T any](t testing.TB, slice []T, element T) {
	t.Helper()

	for _, v := range slice {
		if isEqual(v, element) {
			return
		}
	}

	failCompare[any](t, element, slice, "slice does not contain expected element")
}

// Empty checks if a collection (slice, map, string, or array) is empty.
// It provides a clear error message if the collection contains elements.
func Empty(t testing.TB, collection any, msg ...string) {
	t.Helper()
	v := reflect.ValueOf(collection)

	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		if v.Len() != 0 {
			failCompare(t,
				"empty colleciton",
				fmt.Sprintf("collection with length %d", v.Len()),
				msg...,
			)
		}
	default:
		t.Errorf("\n%s\nEmpty called with unsupported type: %T", location(), collection)
	}
}

// HasKey checks if a map contains a specific key.
func HasKey[K comparable, V any](t testing.TB, m map[K]V, key K) {
	t.Helper()

	if _, ok := m[key]; !ok {
		failCompare[any](t, key, m, "map does not contain expected key")
	}
}

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

// Len checks if a collection (slice, array, map, or string) has the expected length.
func Len(t testing.TB, collection any, expected int) {
	t.Helper()

	v := reflect.ValueOf(collection)
	switch v.Kind() {
	case reflect.Slice,
		reflect.Array,
		reflect.Map,
		reflect.String:
		if v.Len() != expected {
			failCompare(t, expected, v.Len(), "unexpected length")
		}
	default:
		t.Errorf("\n%s\nLen called with unsupported type: %T", location(), collection)
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

// NotContains verifies that a slice does NOT contain an element.
// Useful for ensuring exclusion of specific values.
func NotContains[T any](t testing.TB, slice []T, element T, msg ...string) {
	t.Helper()

	for _, v := range slice {
		if isEqual(v, element) {
			failCompare[any](t, slice, fmt.Sprintf("should not contain %v", element), msg...)
			return
		}
	}
}

// StringContains checks if a string contains an expected substring.
func StringContains(t testing.TB, s, substr string) {
	t.Helper()

	if !strings.Contains(s, substr) {
		failCompare(t, substr, s, "string does not contain expected substring")
	}
}
