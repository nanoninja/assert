// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package require

import (
	"testing"
	"time"

	"github.com/nanoninja/assert"
)

// -- Comparisons --

func Equalf[T any](t testing.TB, expected, actual T, format string, args ...any) {
	assert.Equalf(wrap(t), expected, actual, format, args...)
}

func NotEqualf[T any](t testing.TB, expected, actual T, format string, args ...any) {
	assert.NotEqualf(wrap(t), expected, actual, format, args...)
}

func Truef(t testing.TB, value bool, format string, args ...any) {
	assert.Truef(wrap(t), value, format, args...)
}

func Falsef(t testing.TB, value bool, format string, args ...any) {
	assert.Falsef(wrap(t), value, format, args...)
}

func Nilf(t testing.TB, value any, format string, args ...any) {
	assert.Nilf(wrap(t), value, format, args...)
}

func NotNilf(t testing.TB, value any, format string, args ...any) {
	assert.NotNilf(wrap(t), value, format, args...)
}

func IsTypef(t testing.TB, expected, actual any, format string, args ...any) {
	assert.IsTypef(wrap(t), expected, actual, format, args...)
}

func Implementsf(t testing.TB, iface, value any, format string, args ...any) {
	assert.Implementsf(wrap(t), iface, value, format, args...)
}

// -- Errors --

func Errorf(t testing.TB, err error, format string, args ...any) {
	assert.Errorf(wrap(t), err, format, args...)
}

func NoErrorf(t testing.TB, err error, format string, args ...any) {
	assert.NoErrorf(wrap(t), err, format, args...)
}

func EqualErrorf(t testing.TB, err error, errString, format string, args ...any) {
	assert.EqualErrorf(wrap(t), err, errString, format, args...)
}

func ErrorIsf(t testing.TB, err, target error, format string, args ...any) {
	assert.ErrorIsf(wrap(t), err, target, format, args...)
}

func ErrorAsf(t testing.TB, err error, target any, format string, args ...any) {
	assert.ErrorAsf(wrap(t), err, target, format, args...)
}

// -- Panics --

func NotPanicsf(t testing.TB, fn func(), format string, args ...any) {
	assert.NotPanicsf(wrap(t), fn, format, args...)
}

// -- Collections --

func Containsf[T any](t testing.TB, slice []T, element T, format string, args ...any) {
	assert.Containsf(wrap(t), slice, element, format, args...)
}

func NotContainsf[T any](t testing.TB, slice []T, element T, format string, args ...any) {
	assert.NotContainsf(wrap(t), slice, element, format, args...)
}

func ElementsMatchf[T any](t testing.TB, expected, actual []T, format string, args ...any) {
	assert.ElementsMatchf(wrap(t), expected, actual, format, args...)
}

func Subsetf[T any](t testing.TB, list, subset []T, format string, args ...any) {
	assert.Subsetf(wrap(t), list, subset, format, args...)
}

func NotSubsetf[T any](t testing.TB, list, subset []T, format string, args ...any) {
	assert.NotSubsetf(wrap(t), list, subset, format, args...)
}

func Emptyf(t testing.TB, collection any, format string, args ...any) {
	assert.Emptyf(wrap(t), collection, format, args...)
}

func NotEmptyf(t testing.TB, collection any, format string, args ...any) {
	assert.NotEmptyf(wrap(t), collection, format, args...)
}

func Lenf(t testing.TB, collection any, expected int, format string, args ...any) {
	assert.Lenf(wrap(t), collection, expected, format, args...)
}

func HasKeyf[K comparable, V any](t testing.TB, m map[K]V, key K, format string, args ...any) {
	assert.HasKeyf(wrap(t), m, key, format, args...)
}

func NotHasKeyf[K comparable, V any](t testing.TB, m map[K]V, key K, format string, args ...any) {
	assert.NotHasKeyf(wrap(t), m, key, format, args...)
}

// -- Strings --

func StringContainsf(t testing.TB, s, substr string, format string, args ...any) {
	assert.StringContainsf(wrap(t), s, substr, format, args...)
}

func HasPrefixf(t testing.TB, s, prefix string, format string, args ...any) {
	assert.HasPrefixf(wrap(t), s, prefix, format, args...)
}

func HasSuffixf(t testing.TB, s, suffix string, format string, args ...any) {
	assert.HasSuffixf(wrap(t), s, suffix, format, args...)
}

func MatchRegexpf(t testing.TB, s, pattern string, format string, args ...any) {
	assert.MatchRegexpf(wrap(t), s, pattern, format, args...)
}

// -- Numerics --

func Greaterf[T assert.Ordered](t testing.TB, e1, e2 T, format string, args ...any) {
	assert.Greaterf(wrap(t), e1, e2, format, args...)
}

func GreaterOrEqualf[T assert.Ordered](t testing.TB, actual, min T, format string, args ...any) {
	assert.GreaterOrEqualf(wrap(t), actual, min, format, args...)
}

func Lessf[T assert.Ordered](t testing.TB, e1, e2 T, format string, args ...any) {
	assert.Lessf(wrap(t), e1, e2, format, args...)
}

func LessOrEqualf[T assert.Ordered](t testing.TB, actual, max T, format string, args ...any) {
	assert.LessOrEqualf(wrap(t), actual, max, format, args...)
}

func Betweenf[T assert.Ordered](t testing.TB, actual, min, max T, format string, args ...any) {
	assert.Betweenf(wrap(t), actual, min, max, format, args...)
}

func Positivef[T assert.Number](t testing.TB, value T, format string, args ...any) {
	assert.Positivef(wrap(t), value, format, args...)
}

func Negativef[T assert.Number](t testing.TB, value T, format string, args ...any) {
	assert.Negativef(wrap(t), value, format, args...)
}

func InDeltaf[T assert.Number](t testing.TB, expected, actual T, delta float64, format string, args ...any) {
	assert.InDeltaf(wrap(t), expected, actual, delta, format, args...)
}

// -- JSON --

func EqualJSONf(t testing.TB, expected any, actualJSON string, format string, args ...any) {
	assert.EqualJSONf(wrap(t), expected, actualJSON, format, args...)
}

func JSONEqualf(t testing.TB, expected, actual string, format string, args ...any) {
	assert.JSONEqualf(wrap(t), expected, actual, format, args...)
}

// -- Files --

func FileExistsf(t testing.TB, path string, format string, args ...any) {
	assert.FileExistsf(wrap(t), path, format, args...)
}

func DirExistsf(t testing.TB, path string, format string, args ...any) {
	assert.DirExistsf(wrap(t), path, format, args...)
}

// -- Time --

func WithinDurationf(t testing.TB, expected, actual time.Time, delta time.Duration, format string, args ...any) {
	assert.WithinDurationf(wrap(t), expected, actual, delta, format, args...)
}
