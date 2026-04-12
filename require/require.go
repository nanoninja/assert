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

func Equal[T any](t testing.TB, expected, actual T, msg ...string) {
	assert.Equal(wrap(t), expected, actual, msg...)
}

func NotEqual[T any](t testing.TB, expected, actual T, msg ...string) {
	assert.NotEqual(wrap(t), expected, actual, msg...)
}

func True(t testing.TB, value bool, msg ...string) {
	assert.True(wrap(t), value, msg...)
}

func False(t testing.TB, value bool, msg ...string) {
	assert.False(wrap(t), value, msg...)
}

func Nil(t testing.TB, value any, msg ...string) {
	assert.Nil(wrap(t), value, msg...)
}

func NotNil(t testing.TB, value any, msg ...string) {
	assert.NotNil(wrap(t), value, msg...)
}

func IsType(t testing.TB, expected, actual any, msg ...string) {
	assert.IsType(wrap(t), expected, actual, msg...)
}

func Implements(t testing.TB, iface, value any, msg ...string) {
	assert.Implements(wrap(t), iface, value, msg...)
}

// -- Errors --

func Error(t testing.TB, err error, msg ...string) {
	assert.Error(wrap(t), err, msg...)
}

func NoError(t testing.TB, err error, msg ...string) {
	assert.NoError(wrap(t), err, msg...)
}

func EqualError(t testing.TB, err error, errString string, msg ...string) {
	assert.EqualError(wrap(t), err, errString, msg...)
}

func ErrorIs(t testing.TB, err, target error, msg ...string) {
	assert.ErrorIs(wrap(t), err, target, msg...)
}

func ErrorAs(t testing.TB, err error, target any, msg ...string) {
	assert.ErrorAs(wrap(t), err, target, msg...)
}

// -- Panics --

func Panics(t testing.TB, fn func(), expectedMsg string) {
	assert.Panics(wrap(t), fn, expectedMsg)
}

func NotPanics(t testing.TB, fn func(), msg ...string) {
	assert.NotPanics(wrap(t), fn, msg...)
}

func PanicsWithError(t testing.TB, fn func(), expectedMsg string) {
	assert.PanicsWithError(wrap(t), fn, expectedMsg)
}

// -- Collections --

func Contains[T any](t testing.TB, slice []T, element T, msg ...string) {
	assert.Contains(wrap(t), slice, element, msg...)
}

func NotContains[T any](t testing.TB, slice []T, element T, msg ...string) {
	assert.NotContains(wrap(t), slice, element, msg...)
}

func ElementsMatch[T any](t testing.TB, expected, actual []T, msg ...string) {
	assert.ElementsMatch(wrap(t), expected, actual, msg...)
}

func Subset[T any](t testing.TB, list, subset []T, msg ...string) {
	assert.Subset(wrap(t), list, subset, msg...)
}

func NotSubset[T any](t testing.TB, list, subset []T, msg ...string) {
	assert.NotSubset(wrap(t), list, subset, msg...)
}

func Empty(t testing.TB, collection any, msg ...string) {
	assert.Empty(wrap(t), collection, msg...)
}

func NotEmpty(t testing.TB, collection any, msg ...string) {
	assert.NotEmpty(wrap(t), collection, msg...)
}

func Len(t testing.TB, collection any, expected int, msg ...string) {
	assert.Len(wrap(t), collection, expected, msg...)
}

func HasKey[K comparable, V any](t testing.TB, m map[K]V, key K, msg ...string) {
	assert.HasKey(wrap(t), m, key, msg...)
}

// -- Strings --

func StringContains(t testing.TB, s, substr string, msg ...string) {
	assert.StringContains(wrap(t), s, substr, msg...)
}

func HasPrefix(t testing.TB, s, prefix string, msg ...string) {
	assert.HasPrefix(wrap(t), s, prefix, msg...)
}

func HasSuffix(t testing.TB, s, suffix string, msg ...string) {
	assert.HasSuffix(wrap(t), s, suffix, msg...)
}

func MatchRegexp(t testing.TB, s, pattern string, msg ...string) {
	assert.MatchRegexp(wrap(t), s, pattern, msg...)
}

// -- Numerics --

func Greater[T assert.Ordered](t testing.TB, e1, e2 T, msg ...string) {
	assert.Greater(wrap(t), e1, e2, msg...)
}

func GreaterOrEqual[T assert.Ordered](t testing.TB, actual, min T, msg ...string) {
	assert.GreaterOrEqual(wrap(t), actual, min, msg...)
}

func Less[T assert.Ordered](t testing.TB, e1, e2 T, msg ...string) {
	assert.Less(wrap(t), e1, e2, msg...)
}

func LessOrEqual[T assert.Ordered](t testing.TB, actual, max T, msg ...string) {
	assert.LessOrEqual(wrap(t), actual, max, msg...)
}

func Between[T assert.Ordered](t testing.TB, actual, min, max T, msg ...string) {
	assert.Between(wrap(t), actual, min, max, msg...)
}

func Positive[T assert.Number](t testing.TB, value T, msg ...string) {
	assert.Positive(wrap(t), value, msg...)
}

func Negative[T assert.Number](t testing.TB, value T, msg ...string) {
	assert.Negative(wrap(t), value, msg...)
}

func InDelta[T assert.Number](t testing.TB, expected, actual T, delta float64, msg ...string) {
	assert.InDelta(wrap(t), expected, actual, delta, msg...)
}

// -- JSON --

func JSONEqual(t testing.TB, expected, actual string, msg ...string) {
	assert.JSONEqual(wrap(t), expected, actual, msg...)
}

// -- Time --

func WithinDuration(t testing.TB, expected, actual time.Time, delta time.Duration, msg ...string) {
	assert.WithinDuration(wrap(t), expected, actual, delta, msg...)
}
