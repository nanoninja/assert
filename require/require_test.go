// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package require_test

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/nanoninja/assert/require"
)

// failNowRecorder wraps testing.T and tracks whether FailNow was called.
type failNowRecorder struct {
	*testing.T
	failNowCalled bool
}

func (r *failNowRecorder) FailNow() {
	r.failNowCalled = true
}

func (r *failNowRecorder) Error(args ...any) {
	// intercept without marking the real test as failed
}

func (r *failNowRecorder) Errorf(format string, args ...any) {
	// intercept without marking the real test as failed
}

func (r *failNowRecorder) Helper() {
	r.T.Helper()
}

func newRecorder(t *testing.T) *failNowRecorder {
	t.Helper()
	return &failNowRecorder{T: t}
}

func assertStopped(t *testing.T, rec *failNowRecorder, name string) {
	t.Helper()
	if !rec.failNowCalled {
		t.Errorf("require.%s() did not call FailNow on failure", name)
	}
}

func assertContinued(t *testing.T, rec *failNowRecorder, name string) {
	t.Helper()
	if rec.failNowCalled {
		t.Errorf("require.%s() called FailNow on success", name)
	}
}

// -- Comparisons --

func TestEqual(t *testing.T) {
	rec := newRecorder(t)
	require.Equal(rec, 1, 2)
	assertStopped(t, rec, "Equal")

	rec = newRecorder(t)
	require.Equal(rec, 1, 1)
	assertContinued(t, rec, "Equal")
}

func TestNotEqual(t *testing.T) {
	rec := newRecorder(t)
	require.NotEqual(rec, 1, 1)
	assertStopped(t, rec, "NotEqual")

	rec = newRecorder(t)
	require.NotEqual(rec, 1, 2)
	assertContinued(t, rec, "NotEqual")
}

func TestTrue(t *testing.T) {
	rec := newRecorder(t)
	require.True(rec, false)
	assertStopped(t, rec, "True")

	rec = newRecorder(t)
	require.True(rec, true)
	assertContinued(t, rec, "True")
}

func TestFalse(t *testing.T) {
	rec := newRecorder(t)
	require.False(rec, true)
	assertStopped(t, rec, "False")

	rec = newRecorder(t)
	require.False(rec, false)
	assertContinued(t, rec, "False")
}

func TestNil(t *testing.T) {
	rec := newRecorder(t)
	require.Nil(rec, "not nil")
	assertStopped(t, rec, "Nil")

	rec = newRecorder(t)
	require.Nil(rec, nil)
	assertContinued(t, rec, "Nil")
}

func TestNotNil(t *testing.T) {
	rec := newRecorder(t)
	require.NotNil(rec, nil)
	assertStopped(t, rec, "NotNil")

	rec = newRecorder(t)
	require.NotNil(rec, "value")
	assertContinued(t, rec, "NotNil")
}

func TestIsType(t *testing.T) {
	rec := newRecorder(t)
	require.IsType(rec, 0, "string")
	assertStopped(t, rec, "IsType")

	rec = newRecorder(t)
	require.IsType(rec, 0, 1)
	assertContinued(t, rec, "IsType")
}

func TestImplements(t *testing.T) {
	rec := newRecorder(t)
	require.Implements(rec, (*io.Reader)(nil), "not a reader")
	assertStopped(t, rec, "Implements")

	rec = newRecorder(t)
	require.Implements(rec, (*io.Reader)(nil), strings.NewReader("ok"))
	assertContinued(t, rec, "Implements")
}

// -- Errors --

func TestNoError(t *testing.T) {
	rec := newRecorder(t)
	require.NoError(rec, fmt.Errorf("unexpected error"))
	assertStopped(t, rec, "NoError")

	rec = newRecorder(t)
	require.NoError(rec, nil)
	assertContinued(t, rec, "NoError")
}

func TestError(t *testing.T) {
	rec := newRecorder(t)
	require.Error(rec, nil)
	assertStopped(t, rec, "Error")

	rec = newRecorder(t)
	require.Error(rec, fmt.Errorf("an error"))
	assertContinued(t, rec, "Error")
}

func TestEqualError(t *testing.T) {
	rec := newRecorder(t)
	require.EqualError(rec, fmt.Errorf("actual"), "expected")
	assertStopped(t, rec, "EqualError")

	rec = newRecorder(t)
	require.EqualError(rec, fmt.Errorf("same"), "same")
	assertContinued(t, rec, "EqualError")
}

func TestErrorIs(t *testing.T) {
	sentinel := errors.New("sentinel")
	rec := newRecorder(t)
	require.ErrorIs(rec, fmt.Errorf("other"), sentinel)
	assertStopped(t, rec, "ErrorIs")

	rec = newRecorder(t)
	require.ErrorIs(rec, fmt.Errorf("wrap: %w", sentinel), sentinel)
	assertContinued(t, rec, "ErrorIs")
}

type customErr struct{ msg string }

func (e *customErr) Error() string { return e.msg }

func TestErrorAs(t *testing.T) {
	var target *customErr
	rec := newRecorder(t)
	require.ErrorAs(rec, errors.New("plain"), &target)
	assertStopped(t, rec, "ErrorAs")

	rec = newRecorder(t)
	require.ErrorAs(rec, &customErr{"oops"}, &target)
	assertContinued(t, rec, "ErrorAs")
}

// -- Panics --

func TestPanics(t *testing.T) {
	rec := newRecorder(t)
	require.Panics(rec, func() {}, "expected panic")
	assertStopped(t, rec, "Panics")

	rec = newRecorder(t)
	require.Panics(rec, func() { panic("expected panic") }, "expected panic")
	assertContinued(t, rec, "Panics")
}

func TestNotPanics(t *testing.T) {
	rec := newRecorder(t)
	require.NotPanics(rec, func() { panic("oops") })
	assertStopped(t, rec, "NotPanics")

	rec = newRecorder(t)
	require.NotPanics(rec, func() {})
	assertContinued(t, rec, "NotPanics")
}

func TestPanicsWithError(t *testing.T) {
	rec := newRecorder(t)
	require.PanicsWithError(rec, func() { panic("wrong") }, "expected message")
	assertStopped(t, rec, "PanicsWithError")

	rec = newRecorder(t)
	require.PanicsWithError(rec, func() { panic(errors.New("expected message")) }, "expected message")
	assertContinued(t, rec, "PanicsWithError")
}

// -- Collections --

func TestContains(t *testing.T) {
	rec := newRecorder(t)
	require.Contains(rec, []string{"a", "b"}, "c")
	assertStopped(t, rec, "Contains")

	rec = newRecorder(t)
	require.Contains(rec, []string{"a", "b"}, "a")
	assertContinued(t, rec, "Contains")
}

func TestNotContains(t *testing.T) {
	rec := newRecorder(t)
	require.NotContains(rec, []string{"a", "b"}, "a")
	assertStopped(t, rec, "NotContains")

	rec = newRecorder(t)
	require.NotContains(rec, []string{"a", "b"}, "c")
	assertContinued(t, rec, "NotContains")
}

func TestElementsMatch(t *testing.T) {
	rec := newRecorder(t)
	require.ElementsMatch(rec, []int{1, 2, 3}, []int{1, 2})
	assertStopped(t, rec, "ElementsMatch")

	rec = newRecorder(t)
	require.ElementsMatch(rec, []int{1, 2, 3}, []int{3, 1, 2})
	assertContinued(t, rec, "ElementsMatch")
}

func TestSubset(t *testing.T) {
	rec := newRecorder(t)
	require.Subset(rec, []string{"a", "b"}, []string{"a", "c"})
	assertStopped(t, rec, "Subset")

	rec = newRecorder(t)
	require.Subset(rec, []string{"a", "b", "c"}, []string{"a", "b"})
	assertContinued(t, rec, "Subset")
}

func TestNotSubset(t *testing.T) {
	rec := newRecorder(t)
	require.NotSubset(rec, []string{"a", "b", "c"}, []string{"a", "b"})
	assertStopped(t, rec, "NotSubset")

	rec = newRecorder(t)
	require.NotSubset(rec, []string{"a", "b"}, []string{"a", "c"})
	assertContinued(t, rec, "NotSubset")
}

func TestEmpty(t *testing.T) {
	rec := newRecorder(t)
	require.Empty(rec, []int{1})
	assertStopped(t, rec, "Empty")

	rec = newRecorder(t)
	require.Empty(rec, []int{})
	assertContinued(t, rec, "Empty")
}

func TestNotEmpty(t *testing.T) {
	rec := newRecorder(t)
	require.NotEmpty(rec, []int{})
	assertStopped(t, rec, "NotEmpty")

	rec = newRecorder(t)
	require.NotEmpty(rec, []int{1})
	assertContinued(t, rec, "NotEmpty")
}

func TestLen(t *testing.T) {
	rec := newRecorder(t)
	require.Len(rec, []int{1, 2}, 3)
	assertStopped(t, rec, "Len")

	rec = newRecorder(t)
	require.Len(rec, []int{1, 2}, 2)
	assertContinued(t, rec, "Len")
}

func TestHasKey(t *testing.T) {
	rec := newRecorder(t)
	require.HasKey(rec, map[string]int{"a": 1}, "b")
	assertStopped(t, rec, "HasKey")

	rec = newRecorder(t)
	require.HasKey(rec, map[string]int{"a": 1}, "a")
	assertContinued(t, rec, "HasKey")
}

// -- Strings --

func TestStringContains(t *testing.T) {
	rec := newRecorder(t)
	require.StringContains(rec, "hello world", "xyz")
	assertStopped(t, rec, "StringContains")

	rec = newRecorder(t)
	require.StringContains(rec, "hello world", "world")
	assertContinued(t, rec, "StringContains")
}

func TestHasPrefix(t *testing.T) {
	rec := newRecorder(t)
	require.HasPrefix(rec, "hello", "world")
	assertStopped(t, rec, "HasPrefix")

	rec = newRecorder(t)
	require.HasPrefix(rec, "hello", "hel")
	assertContinued(t, rec, "HasPrefix")
}

func TestHasSuffix(t *testing.T) {
	rec := newRecorder(t)
	require.HasSuffix(rec, "hello", "world")
	assertStopped(t, rec, "HasSuffix")

	rec = newRecorder(t)
	require.HasSuffix(rec, "hello", "llo")
	assertContinued(t, rec, "HasSuffix")
}

func TestMatchRegexp(t *testing.T) {
	rec := newRecorder(t)
	require.MatchRegexp(rec, "hello", `^\d+$`)
	assertStopped(t, rec, "MatchRegexp")

	rec = newRecorder(t)
	require.MatchRegexp(rec, "hello", `^[a-z]+$`)
	assertContinued(t, rec, "MatchRegexp")
}

// -- Numerics --

func TestGreater(t *testing.T) {
	rec := newRecorder(t)
	require.Greater(rec, 1, 2)
	assertStopped(t, rec, "Greater")

	rec = newRecorder(t)
	require.Greater(rec, 2, 1)
	assertContinued(t, rec, "Greater")
}

func TestGreaterOrEqual(t *testing.T) {
	rec := newRecorder(t)
	require.GreaterOrEqual(rec, 1, 2)
	assertStopped(t, rec, "GreaterOrEqual")

	rec = newRecorder(t)
	require.GreaterOrEqual(rec, 2, 2)
	assertContinued(t, rec, "GreaterOrEqual")
}

func TestLess(t *testing.T) {
	rec := newRecorder(t)
	require.Less(rec, 2, 1)
	assertStopped(t, rec, "Less")

	rec = newRecorder(t)
	require.Less(rec, 1, 2)
	assertContinued(t, rec, "Less")
}

func TestLessOrEqual(t *testing.T) {
	rec := newRecorder(t)
	require.LessOrEqual(rec, 3, 2)
	assertStopped(t, rec, "LessOrEqual")

	rec = newRecorder(t)
	require.LessOrEqual(rec, 2, 2)
	assertContinued(t, rec, "LessOrEqual")
}

func TestBetween(t *testing.T) {
	rec := newRecorder(t)
	require.Between(rec, 10, 1, 5)
	assertStopped(t, rec, "Between")

	rec = newRecorder(t)
	require.Between(rec, 3, 1, 5)
	assertContinued(t, rec, "Between")
}

func TestPositive(t *testing.T) {
	rec := newRecorder(t)
	require.Positive(rec, -1)
	assertStopped(t, rec, "Positive")

	rec = newRecorder(t)
	require.Positive(rec, 1)
	assertContinued(t, rec, "Positive")
}

func TestNegative(t *testing.T) {
	rec := newRecorder(t)
	require.Negative(rec, 1)
	assertStopped(t, rec, "Negative")

	rec = newRecorder(t)
	require.Negative(rec, -1)
	assertContinued(t, rec, "Negative")
}

func TestInDelta(t *testing.T) {
	rec := newRecorder(t)
	require.InDelta(rec, 1.0, 2.0, 0.5)
	assertStopped(t, rec, "InDelta")

	rec = newRecorder(t)
	require.InDelta(rec, 1.0, 1.1, 0.5)
	assertContinued(t, rec, "InDelta")
}

// -- JSON --

func TestJSONEqual(t *testing.T) {
	rec := newRecorder(t)
	require.JSONEqual(rec, `{"a":1}`, `{"a":2}`)
	assertStopped(t, rec, "JSONEqual")

	rec = newRecorder(t)
	require.JSONEqual(rec, `{"a":1,"b":2}`, `{"b":2,"a":1}`)
	assertContinued(t, rec, "JSONEqual")
}

// -- Time --

func TestWithinDuration(t *testing.T) {
	now := time.Now()
	rec := newRecorder(t)
	require.WithinDuration(rec, now, now.Add(2*time.Second), time.Second)
	assertStopped(t, rec, "WithinDuration")

	rec = newRecorder(t)
	require.WithinDuration(rec, now, now.Add(500*time.Millisecond), time.Second)
	assertContinued(t, rec, "WithinDuration")
}
