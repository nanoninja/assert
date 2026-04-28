// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"fmt"
	"testing"
	"time"
)

// -- assert.go --

func Equalf[T any](t testing.TB, expected, actual T, format string, args ...any) {
	t.Helper()
	Equal(t, expected, actual, fmt.Sprintf(format, args...))
}

func EqualErrorf(t testing.TB, err error, errStr, format string, args ...any) {
	t.Helper()
	EqualError(t, err, errStr, fmt.Sprintf(format, args...))
}

func Errorf(t testing.TB, err error, format string, args ...any) {
	t.Helper()
	Error(t, err, fmt.Sprintf(format, args...))
}

func ErrorAsf(t testing.TB, err error, target any, format string, args ...any) {
	t.Helper()
	ErrorAs(t, err, target, fmt.Sprintf(format, args...))
}

func ErrorIsf(t testing.TB, err, target error, format string, args ...any) {
	t.Helper()
	ErrorIs(t, err, target, fmt.Sprintf(format, args...))
}

func Falsef(t testing.TB, val bool, format string, args ...any) {
	t.Helper()
	False(t, val, fmt.Sprintf(format, args...))
}

func Nilf(t testing.TB, val any, format string, args ...any) {
	t.Helper()
	Nil(t, val, fmt.Sprintf(format, args...))
}

func NoErrorf(t testing.TB, err error, format string, args ...any) {
	t.Helper()
	NoError(t, err, fmt.Sprintf(format, args...))
}

func NotEqualf[T any](t testing.TB, expected, actual T, format string, args ...any) {
	t.Helper()
	NotEqual(t, expected, actual, fmt.Sprintf(format, args...))
}

func NotNilf(t testing.TB, val any, format string, args ...any) {
	t.Helper()
	NotNil(t, val, fmt.Sprintf(format, args...))
}

func IsTypef(t testing.TB, expected, actual any, format string, args ...any) {
	t.Helper()
	IsType(t, expected, actual, fmt.Sprintf(format, args...))
}

func Implementsf(t testing.TB, iface, value any, format string, args ...any) {
	t.Helper()
	Implements(t, iface, value, fmt.Sprintf(format, args...))
}

func NotPanicsf(t testing.TB, fn func(), format string, args ...any) {
	t.Helper()
	NotPanics(t, fn, fmt.Sprintf(format, args...))
}

func Truef(t testing.TB, val bool, format string, args ...any) {
	t.Helper()
	True(t, val, fmt.Sprintf(format, args...))
}

// -- collection.go --

func Containsf[T any](t testing.TB, slice []T, element T, format string, args ...any) {
	t.Helper()
	Contains(t, slice, element, fmt.Sprintf(format, args...))
}

func Emptyf(t testing.TB, collection any, format string, args ...any) {
	t.Helper()
	Empty(t, collection, fmt.Sprintf(format, args...))
}

func NotEmptyf(t testing.TB, collection any, format string, args ...any) {
	t.Helper()
	NotEmpty(t, collection, fmt.Sprintf(format, args...))
}

func HasKeyf[K comparable, V any](t testing.TB, m map[K]V, key K, format string, args ...any) {
	t.Helper()
	HasKey(t, m, key, fmt.Sprintf(format, args...))
}

func NotHasKeyf[K comparable, V any](t testing.TB, m map[K]V, key K, format string, args ...any) {
	t.Helper()
	NotHasKey(t, m, key, fmt.Sprintf(format, args...))
}

func Lenf(t testing.TB, collection any, expected int, format string, args ...any) {
	t.Helper()
	Len(t, collection, expected, fmt.Sprintf(format, args...))
}

func NotContainsf[T any](t testing.TB, slice []T, element T, format string, args ...any) {
	t.Helper()
	NotContains(t, slice, element, fmt.Sprintf(format, args...))
}

func ElementsMatchf[T any](t testing.TB, expected, actual []T, format string, args ...any) {
	t.Helper()
	ElementsMatch(t, expected, actual, fmt.Sprintf(format, args...))
}

func Subsetf[T any](t testing.TB, list, subset []T, format string, args ...any) {
	t.Helper()
	Subset(t, list, subset, fmt.Sprintf(format, args...))
}

func NotSubsetf[T any](t testing.TB, list, subset []T, format string, args ...any) {
	t.Helper()
	NotSubset(t, list, subset, fmt.Sprintf(format, args...))
}

// -- file.go --

func FileExistsf(t testing.TB, path string, format string, args ...any) {
	t.Helper()
	FileExists(t, path, fmt.Sprintf(format, args...))
}

func DirExistsf(t testing.TB, path string, format string, args ...any) {
	t.Helper()
	DirExists(t, path, fmt.Sprintf(format, args...))
}

// -- numeric.go --

func Betweenf[T Ordered](t testing.TB, actual, min, max T, format string, args ...any) {
	t.Helper()
	Between(t, actual, min, max, fmt.Sprintf(format, args...))
}

func Greaterf[T Ordered](t testing.TB, e1, e2 T, format string, args ...any) {
	t.Helper()
	Greater(t, e1, e2, fmt.Sprintf(format, args...))
}

func Lessf[T Ordered](t testing.TB, e1, e2 T, format string, args ...any) {
	t.Helper()
	Less(t, e1, e2, fmt.Sprintf(format, args...))
}

func GreaterOrEqualf[T Ordered](t testing.TB, actual, min T, format string, args ...any) {
	t.Helper()
	GreaterOrEqual(t, actual, min, fmt.Sprintf(format, args...))
}

func LessOrEqualf[T Ordered](t testing.TB, actual, max T, format string, args ...any) {
	t.Helper()
	LessOrEqual(t, actual, max, fmt.Sprintf(format, args...))
}

func Positivef[T Number](t testing.TB, value T, format string, args ...any) {
	t.Helper()
	Positive(t, value, fmt.Sprintf(format, args...))
}

func Negativef[T Number](t testing.TB, value T, format string, args ...any) {
	t.Helper()
	Negative(t, value, fmt.Sprintf(format, args...))
}

func InDeltaf[T Number](t testing.TB, expected, actual T, delta float64, format string, args ...any) {
	t.Helper()
	InDelta(t, expected, actual, delta, fmt.Sprintf(format, args...))
}

// -- string.go --

func HasPrefixf(t testing.TB, s, prefix string, format string, args ...any) {
	t.Helper()
	HasPrefix(t, s, prefix, fmt.Sprintf(format, args...))
}

func HasSuffixf(t testing.TB, s, suffix string, format string, args ...any) {
	t.Helper()
	HasSuffix(t, s, suffix, fmt.Sprintf(format, args...))
}

func MatchRegexpf(t testing.TB, s, pattern string, format string, args ...any) {
	t.Helper()
	MatchRegexp(t, s, pattern, fmt.Sprintf(format, args...))
}

func StringContainsf(t testing.TB, s, substr string, format string, args ...any) {
	t.Helper()
	StringContains(t, s, substr, fmt.Sprintf(format, args...))
}

// -- json.go --

func EqualJSONf(t testing.TB, expected any, actualJSON string, format string, args ...any) {
	t.Helper()
	EqualJSON(t, expected, actualJSON, fmt.Sprintf(format, args...))
}

func JSONEqualf(t testing.TB, expected, actual string, format string, args ...any) {
	t.Helper()
	JSONEqual(t, expected, actual, fmt.Sprintf(format, args...))
}

// -- time.go --

func WithinDurationf(t testing.TB, expected, actual time.Time, delta time.Duration, format string, args ...any) {
	t.Helper()
	WithinDuration(t, expected, actual, delta, fmt.Sprintf(format, args...))
}
