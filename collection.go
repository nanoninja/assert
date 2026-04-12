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
