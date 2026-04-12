// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package require provides the same assertions as the assert package
// but stops the test immediately on failure using t.FailNow().
//
// Use require for preconditions — assertions that must pass before
// the rest of the test makes sense:
//
//	func TestExample(t *testing.T) {
//	    result, err := DoSomething()
//	    require.NoError(t, err)     // stops immediately if err != nil
//	    require.NotNil(t, result)   // stops immediately if nil
//
//	    assert.Equal(t, "expected", result.Name)   // continues on failure
//	    assert.Equal(t, 42, result.Count)          // sees all failures
//	}
package require

import (
	"testing"
	"time"

	"github.com/nanoninja/assert"
)

// fatalT wraps testing.TB and calls FailNow after each error,
// stopping the test immediately on failure.
type fatalT struct {
	testing.TB
}

func (f *fatalT) Error(args ...any) {
	f.TB.Error(args...)
	f.FailNow()
}

func (f *fatalT) Errorf(format string, args ...any) {
	f.TB.Errorf(format, args...)
	f.FailNow()
}

func wrap(t testing.TB) testing.TB {
	return &fatalT{TB: t}
}

// -- Comparisons --

func Equal[T any](t testing.TB, expected, actual T, msgAndArgs ...any) {
	assert.Equal(wrap(t), expected, actual, msgAndArgs...)
}

func NotEqual[T any](t testing.TB, expected, actual T, msgAndArgs ...any) {
	assert.NotEqual(wrap(t), expected, actual, msgAndArgs...)
}

func True(t testing.TB, value bool, msgAndArgs ...any) {
	assert.True(wrap(t), value, msgAndArgs...)
}

func False(t testing.TB, value bool, msgAndArgs ...any) {
	assert.False(wrap(t), value, msgAndArgs...)
}

func Nil(t testing.TB, value any, msgAndArgs ...any) {
	assert.Nil(wrap(t), value, msgAndArgs...)
}

func NotNil(t testing.TB, value any, msgAndArgs ...any) {
	assert.NotNil(wrap(t), value, msgAndArgs...)
}

func IsType(t testing.TB, expected, actual any, msgAndArgs ...any) {
	assert.IsType(wrap(t), expected, actual, msgAndArgs...)
}

func Implements(t testing.TB, iface, value any, msgAndArgs ...any) {
	assert.Implements(wrap(t), iface, value, msgAndArgs...)
}

// -- Errors --

func Error(t testing.TB, err error, msgAndArgs ...any) {
	assert.Error(wrap(t), err, msgAndArgs...)
}

func NoError(t testing.TB, err error, msgAndArgs ...any) {
	assert.NoError(wrap(t), err, msgAndArgs...)
}

func EqualError(t testing.TB, err error, errString string, msgAndArgs ...any) {
	assert.EqualError(wrap(t), err, errString, msgAndArgs...)
}

func ErrorIs(t testing.TB, err, target error, msgAndArgs ...any) {
	assert.ErrorIs(wrap(t), err, target, msgAndArgs...)
}

func ErrorAs(t testing.TB, err error, target any, msgAndArgs ...any) {
	assert.ErrorAs(wrap(t), err, target, msgAndArgs...)
}

// -- Panics --

func Panics(t testing.TB, fn func(), expectedMsg string) {
	assert.Panics(wrap(t), fn, expectedMsg)
}

func NotPanics(t testing.TB, fn func(), msgAndArgs ...any) {
	assert.NotPanics(wrap(t), fn, msgAndArgs...)
}

func PanicsWithError(t testing.TB, fn func(), expectedMsg string) {
	assert.PanicsWithError(wrap(t), fn, expectedMsg)
}

// -- Collections --

func Contains[T any](t testing.TB, slice []T, element T, msgAndArgs ...any) {
	assert.Contains(wrap(t), slice, element, msgAndArgs...)
}

func NotContains[T any](t testing.TB, slice []T, element T, msgAndArgs ...any) {
	assert.NotContains(wrap(t), slice, element, msgAndArgs...)
}

func ElementsMatch[T any](t testing.TB, expected, actual []T, msgAndArgs ...any) {
	assert.ElementsMatch(wrap(t), expected, actual, msgAndArgs...)
}

func Subset[T any](t testing.TB, list, subset []T, msgAndArgs ...any) {
	assert.Subset(wrap(t), list, subset, msgAndArgs...)
}

func NotSubset[T any](t testing.TB, list, subset []T, msgAndArgs ...any) {
	assert.NotSubset(wrap(t), list, subset, msgAndArgs...)
}

func Empty(t testing.TB, collection any, msgAndArgs ...any) {
	assert.Empty(wrap(t), collection, msgAndArgs...)
}

func NotEmpty(t testing.TB, collection any, msgAndArgs ...any) {
	assert.NotEmpty(wrap(t), collection, msgAndArgs...)
}

func Len(t testing.TB, collection any, expected int, msgAndArgs ...any) {
	assert.Len(wrap(t), collection, expected, msgAndArgs...)
}

func HasKey[K comparable, V any](t testing.TB, m map[K]V, key K, msgAndArgs ...any) {
	assert.HasKey(wrap(t), m, key, msgAndArgs...)
}

func NotHasKey[K comparable, V any](t testing.TB, m map[K]V, key K, msgAndArgs ...any) {
	assert.NotHasKey(wrap(t), m, key, msgAndArgs...)
}

// -- Strings --

func StringContains(t testing.TB, s, substr string, msgAndArgs ...any) {
	assert.StringContains(wrap(t), s, substr, msgAndArgs...)
}

func HasPrefix(t testing.TB, s, prefix string, msgAndArgs ...any) {
	assert.HasPrefix(wrap(t), s, prefix, msgAndArgs...)
}

func HasSuffix(t testing.TB, s, suffix string, msgAndArgs ...any) {
	assert.HasSuffix(wrap(t), s, suffix, msgAndArgs...)
}

func MatchRegexp(t testing.TB, s, pattern string, msgAndArgs ...any) {
	assert.MatchRegexp(wrap(t), s, pattern, msgAndArgs...)
}

// -- Numerics --

func Greater[T assert.Ordered](t testing.TB, e1, e2 T, msgAndArgs ...any) {
	assert.Greater(wrap(t), e1, e2, msgAndArgs...)
}

func GreaterOrEqual[T assert.Ordered](t testing.TB, actual, min T, msgAndArgs ...any) {
	assert.GreaterOrEqual(wrap(t), actual, min, msgAndArgs...)
}

func Less[T assert.Ordered](t testing.TB, e1, e2 T, msgAndArgs ...any) {
	assert.Less(wrap(t), e1, e2, msgAndArgs...)
}

func LessOrEqual[T assert.Ordered](t testing.TB, actual, max T, msgAndArgs ...any) {
	assert.LessOrEqual(wrap(t), actual, max, msgAndArgs...)
}

func Between[T assert.Ordered](t testing.TB, actual, min, max T, msgAndArgs ...any) {
	assert.Between(wrap(t), actual, min, max, msgAndArgs...)
}

func Positive[T assert.Number](t testing.TB, value T, msgAndArgs ...any) {
	assert.Positive(wrap(t), value, msgAndArgs...)
}

func Negative[T assert.Number](t testing.TB, value T, msgAndArgs ...any) {
	assert.Negative(wrap(t), value, msgAndArgs...)
}

func InDelta[T assert.Number](t testing.TB, expected, actual T, delta float64, msgAndArgs ...any) {
	assert.InDelta(wrap(t), expected, actual, delta, msgAndArgs...)
}

// -- JSON --

func EqualJSON(t testing.TB, expected any, actualJSON string, msgAndArgs ...any) {
	assert.EqualJSON(wrap(t), expected, actualJSON, msgAndArgs...)
}

func JSONEqual(t testing.TB, expected, actual string, msgAndArgs ...any) {
	assert.JSONEqual(wrap(t), expected, actual, msgAndArgs...)
}

// -- Files --

func FileExists(t testing.TB, path string, msgAndArgs ...any) {
	assert.FileExists(wrap(t), path, msgAndArgs...)
}

func DirExists(t testing.TB, path string, msgAndArgs ...any) {
	assert.DirExists(wrap(t), path, msgAndArgs...)
}

// -- Time --

func WithinDuration(t testing.TB, expected, actual time.Time, delta time.Duration, msgAndArgs ...any) {
	assert.WithinDuration(wrap(t), expected, actual, delta, msgAndArgs...)
}
