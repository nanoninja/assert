// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// Contains checks if a slice contains a specific element.
// The comparison is done using reflection.DeepEqual.
func Contains[T any](t testing.TB, slice []T, element T, msg ...string) {
	t.Helper()

	for _, v := range slice {
		if isEqual(v, element) {
			return
		}
	}

	failCompare[any](t, element, slice, append(
		[]string{"slice does not contain expected element"},
		msg...)...)
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
				"empty collection",
				fmt.Sprintf("collection with length %d", v.Len()),
				msg...,
			)
		}
	default:
		t.Errorf("\nEmpty() called with unsupported type: (%T)", collection)
	}
}

// NotEmpty checks if a collection (slice, map, string, or array) is not empty
func NotEmpty(t testing.TB, collection any, msg ...string) {
	t.Helper()

	v := reflect.ValueOf(collection)

	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		if v.Len() == 0 {
			failCompare(t,
				"non-empty collection",
				"empty collection",
				msg...,
			)
		}
	default:
		t.Errorf("\nNotEmpty() called with unsupported type: (%T)", collection)
	}
}

// HasKey checks if a map contains a specific key.
func HasKey[K comparable, V any](t testing.TB, m map[K]V, key K, msg ...string) {
	t.Helper()

	if _, ok := m[key]; !ok {
		failCompare[any](t, key, m, append(
			[]string{"map does not contain expected key"},
			msg...)...)
	}
}

// NotHasKey checks that a map does NOT contain a specific key.
func NotHasKey[K comparable, V any](t testing.TB, m map[K]V, key K, msg ...string) {
	t.Helper()

	if _, ok := m[key]; ok {
		failCompare[any](t, fmt.Sprintf("map without key %v", key), m, append(
			[]string{"map contains unexpected key"},
			msg...)...)
	}
}

// Len checks if a collection (slice, array, map, or string) has the expected length.
func Len(t testing.TB, collection any, expected int, msg ...string) {
	t.Helper()

	v := reflect.ValueOf(collection)
	switch v.Kind() {
	case reflect.Slice,
		reflect.Array,
		reflect.Map,
		reflect.String:
		if v.Len() != expected {
			failCompare(t, expected, v.Len(), append([]string{"unexpected length"}, msg...)...)
		}
	default:
		t.Errorf("\nLen called with unsupported type: %T", collection)
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

// ElementsMatch checks that two slices contain the same elements regardless of order.
func ElementsMatch[T any](t testing.TB, expected, actual []T, msg ...string) {
	t.Helper()

	if len(expected) != len(actual) {
		failCompare[any](t,
			fmt.Sprintf("slice with %d elements", len(expected)),
			fmt.Sprintf("slice with %d elements", len(actual)),
			append([]string{"slices have different lengths"}, msg...)...,
		)
		return
	}

	used := make([]bool, len(actual))

	for _, e := range expected {
		found := false
		for i, a := range actual {
			if !used[i] && isEqual(e, a) {
				used[i] = true
				found = true
				break
			}
		}
		if !found {
			failCompare[any](t, expected, actual,
				append([]string{"slices do not contain the same elements"}, msg...)...)
			return
		}
	}
}

// Subset checks that every element of subset is present in list.
func Subset[T any](t testing.TB, list, subset []T, msg ...string) {
	t.Helper()

	for _, s := range subset {
		found := false
		for _, l := range list {
			if isEqual(s, l) {
				found = true
				break
			}
		}
		if !found {
			failCompare[any](t, list, subset,
				append([]string{"subset contains elements not found in list"}, msg...)...)
			return
		}
	}
}

// NotSubset checks that subset is NOT entirely contained in list.
func NotSubset[T any](t testing.TB, list, subset []T, msg ...string) {
	t.Helper()

	for _, s := range subset {
		found := false
		for _, l := range list {
			if isEqual(s, l) {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}

	failCompare[any](t, list, subset,
		append([]string{"expected subset not to be contained in list"}, msg...)...)
}
