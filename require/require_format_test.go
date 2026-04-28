// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package require_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/nanoninja/assert/require"
)

// -- Comparisons --

func TestEqualf(t *testing.T) {
	rec := newRecorder(t)
	require.Equalf(rec, 1, 2, "values: %d and %d", 1, 2)
	assertStopped(t, rec, "Equalf")

	rec = newRecorder(t)
	require.Equalf(rec, 1, 1, "should not fail")
	assertContinued(t, rec, "Equalf")
}

func TestNotEqualf(t *testing.T) {
	rec := newRecorder(t)
	require.NotEqualf(rec, 1, 1, "got %d", 1)
	assertStopped(t, rec, "NotEqualf")

	rec = newRecorder(t)
	require.NotEqualf(rec, 1, 2, "should not fail")
	assertContinued(t, rec, "NotEqualf")
}

func TestTruef(t *testing.T) {
	rec := newRecorder(t)
	require.Truef(rec, false, "expected %s", "true")
	assertStopped(t, rec, "Truef")

	rec = newRecorder(t)
	require.Truef(rec, true, "should not fail")
	assertContinued(t, rec, "Truef")
}

func TestFalsef(t *testing.T) {
	rec := newRecorder(t)
	require.Falsef(rec, true, "expected %s", "false")
	assertStopped(t, rec, "Falsef")

	rec = newRecorder(t)
	require.Falsef(rec, false, "should not fail")
	assertContinued(t, rec, "Falsef")
}

func TestNilf(t *testing.T) {
	rec := newRecorder(t)
	require.Nilf(rec, "not nil", "got %q", "not nil")
	assertStopped(t, rec, "Nilf")

	rec = newRecorder(t)
	require.Nilf(rec, nil, "should not fail")
	assertContinued(t, rec, "Nilf")
}

func TestNotNilf(t *testing.T) {
	rec := newRecorder(t)
	require.NotNilf(rec, nil, "value is %s", "nil")
	assertStopped(t, rec, "NotNilf")

	rec = newRecorder(t)
	require.NotNilf(rec, "ok", "should not fail")
	assertContinued(t, rec, "NotNilf")
}

func TestIsTypef(t *testing.T) {
	rec := newRecorder(t)
	require.IsTypef(rec, 0, "str", "type: %s", "mismatch")
	assertStopped(t, rec, "IsTypef")

	rec = newRecorder(t)
	require.IsTypef(rec, 0, 1, "should not fail")
	assertContinued(t, rec, "IsTypef")
}

func TestImplementsf(t *testing.T) {
	rec := newRecorder(t)
	require.Implementsf(rec, (*io.Reader)(nil), "not a reader", "check %s", "interface")
	assertStopped(t, rec, "Implementsf")

	rec = newRecorder(t)
	require.Implementsf(rec, (*io.Reader)(nil), strings.NewReader("ok"), "should not fail")
	assertContinued(t, rec, "Implementsf")
}

// -- Errors --

func TestErrorf(t *testing.T) {
	rec := newRecorder(t)
	require.Errorf(rec, nil, "expected error: %s", "none")
	assertStopped(t, rec, "Errorf")

	rec = newRecorder(t)
	require.Errorf(rec, errors.New("err"), "should not fail")
	assertContinued(t, rec, "Errorf")
}

func TestNoErrorf(t *testing.T) {
	rec := newRecorder(t)
	require.NoErrorf(rec, fmt.Errorf("oops"), "unexpected: %s", "oops")
	assertStopped(t, rec, "NoErrorf")

	rec = newRecorder(t)
	require.NoErrorf(rec, nil, "should not fail")
	assertContinued(t, rec, "NoErrorf")
}

func TestEqualErrorf(t *testing.T) {
	rec := newRecorder(t)
	require.EqualErrorf(rec, errors.New("actual"), "expected", "msg: %s", "mismatch")
	assertStopped(t, rec, "EqualErrorf")

	rec = newRecorder(t)
	require.EqualErrorf(rec, errors.New("same"), "same", "should not fail")
	assertContinued(t, rec, "EqualErrorf")
}

func TestErrorIsf(t *testing.T) {
	sentinel := errors.New("sentinel")

	rec := newRecorder(t)
	require.ErrorIsf(rec, errors.New("other"), sentinel, "err: %s", "not sentinel")
	assertStopped(t, rec, "ErrorIsf")

	rec = newRecorder(t)
	require.ErrorIsf(rec, fmt.Errorf("wrap: %w", sentinel), sentinel, "should not fail")
	assertContinued(t, rec, "ErrorIsf")
}

func TestErrorAsf(t *testing.T) {
	var target *customErr

	rec := newRecorder(t)
	require.ErrorAsf(rec, errors.New("plain"), &target, "err: %s", "wrong type")
	assertStopped(t, rec, "ErrorAsf")

	rec = newRecorder(t)
	require.ErrorAsf(rec, &customErr{"oops"}, &target, "should not fail")
	assertContinued(t, rec, "ErrorAsf")
}

// -- Panics --

func TestNotPanicsf(t *testing.T) {
	rec := newRecorder(t)
	require.NotPanicsf(rec, func() { panic("oops") }, "panic: %s", "unexpected")
	assertStopped(t, rec, "NotPanicsf")

	rec = newRecorder(t)
	require.NotPanicsf(rec, func() {}, "should not fail")
	assertContinued(t, rec, "NotPanicsf")
}

// -- Collections --

func TestContainsf(t *testing.T) {
	rec := newRecorder(t)
	require.Containsf(rec, []string{"a", "b"}, "c", "missing: %s", "c")
	assertStopped(t, rec, "Containsf")

	rec = newRecorder(t)
	require.Containsf(rec, []string{"a", "b"}, "a", "should not fail")
	assertContinued(t, rec, "Containsf")
}

func TestNotContainsf(t *testing.T) {
	rec := newRecorder(t)
	require.NotContainsf(rec, []string{"a", "b"}, "a", "found: %s", "a")
	assertStopped(t, rec, "NotContainsf")

	rec = newRecorder(t)
	require.NotContainsf(rec, []string{"a", "b"}, "c", "should not fail")
	assertContinued(t, rec, "NotContainsf")
}

func TestElementsMatchf(t *testing.T) {
	rec := newRecorder(t)
	require.ElementsMatchf(rec, []int{1, 2, 3}, []int{1, 2}, "len: %d vs %d", 3, 2)
	assertStopped(t, rec, "ElementsMatchf")

	rec = newRecorder(t)
	require.ElementsMatchf(rec, []int{1, 2, 3}, []int{3, 1, 2}, "should not fail")
	assertContinued(t, rec, "ElementsMatchf")
}

func TestSubsetf(t *testing.T) {
	rec := newRecorder(t)
	require.Subsetf(rec, []string{"a", "b"}, []string{"a", "c"}, "not subset: %s", "c missing")
	assertStopped(t, rec, "Subsetf")

	rec = newRecorder(t)
	require.Subsetf(rec, []string{"a", "b", "c"}, []string{"a", "b"}, "should not fail")
	assertContinued(t, rec, "Subsetf")
}

func TestNotSubsetf(t *testing.T) {
	rec := newRecorder(t)
	require.NotSubsetf(rec, []string{"a", "b", "c"}, []string{"a", "b"}, "is subset: %s", "ab")
	assertStopped(t, rec, "NotSubsetf")

	rec = newRecorder(t)
	require.NotSubsetf(rec, []string{"a", "b"}, []string{"a", "c"}, "should not fail")
	assertContinued(t, rec, "NotSubsetf")
}

func TestEmptyf(t *testing.T) {
	rec := newRecorder(t)
	require.Emptyf(rec, []int{1}, "len: %d", 1)
	assertStopped(t, rec, "Emptyf")

	rec = newRecorder(t)
	require.Emptyf(rec, []int{}, "should not fail")
	assertContinued(t, rec, "Emptyf")
}

func TestNotEmptyf(t *testing.T) {
	rec := newRecorder(t)
	require.NotEmptyf(rec, []int{}, "collection: %s", "empty")
	assertStopped(t, rec, "NotEmptyf")

	rec = newRecorder(t)
	require.NotEmptyf(rec, []int{1}, "should not fail")
	assertContinued(t, rec, "NotEmptyf")
}

func TestLenf(t *testing.T) {
	rec := newRecorder(t)
	require.Lenf(rec, []int{1, 2}, 3, "expected len %d", 3)
	assertStopped(t, rec, "Lenf")

	rec = newRecorder(t)
	require.Lenf(rec, []int{1, 2}, 2, "should not fail")
	assertContinued(t, rec, "Lenf")
}

func TestHasKeyf(t *testing.T) {
	rec := newRecorder(t)
	require.HasKeyf(rec, map[string]int{"a": 1}, "b", "missing key: %s", "b")
	assertStopped(t, rec, "HasKeyf")

	rec = newRecorder(t)
	require.HasKeyf(rec, map[string]int{"a": 1}, "a", "should not fail")
	assertContinued(t, rec, "HasKeyf")
}

func TestNotHasKeyf(t *testing.T) {
	rec := newRecorder(t)
	require.NotHasKeyf(rec, map[string]int{"a": 1}, "a", "found key: %s", "a")
	assertStopped(t, rec, "NotHasKeyf")

	rec = newRecorder(t)
	require.NotHasKeyf(rec, map[string]int{"a": 1}, "b", "should not fail")
	assertContinued(t, rec, "NotHasKeyf")
}

// -- Strings --

func TestStringContainsf(t *testing.T) {
	rec := newRecorder(t)
	require.StringContainsf(rec, "hello world", "xyz", "missing: %s", "xyz")
	assertStopped(t, rec, "StringContainsf")

	rec = newRecorder(t)
	require.StringContainsf(rec, "hello world", "world", "should not fail")
	assertContinued(t, rec, "StringContainsf")
}

func TestHasPrefixf(t *testing.T) {
	rec := newRecorder(t)
	require.HasPrefixf(rec, "hello", "world", "prefix: %s", "world")
	assertStopped(t, rec, "HasPrefixf")

	rec = newRecorder(t)
	require.HasPrefixf(rec, "hello", "hel", "should not fail")
	assertContinued(t, rec, "HasPrefixf")
}

func TestHasSuffixf(t *testing.T) {
	rec := newRecorder(t)
	require.HasSuffixf(rec, "hello", "world", "suffix: %s", "world")
	assertStopped(t, rec, "HasSuffixf")

	rec = newRecorder(t)
	require.HasSuffixf(rec, "hello", "llo", "should not fail")
	assertContinued(t, rec, "HasSuffixf")
}

func TestMatchRegexpf(t *testing.T) {
	rec := newRecorder(t)
	require.MatchRegexpf(rec, "hello", `^\d+$`, "pattern: %s", `^\d+$`)
	assertStopped(t, rec, "MatchRegexpf")

	rec = newRecorder(t)
	require.MatchRegexpf(rec, "hello", `^[a-z]+$`, "should not fail")
	assertContinued(t, rec, "MatchRegexpf")
}

// -- Numerics --

func TestGreaterf(t *testing.T) {
	rec := newRecorder(t)
	require.Greaterf(rec, 1, 2, "%d not greater than %d", 1, 2)
	assertStopped(t, rec, "Greaterf")

	rec = newRecorder(t)
	require.Greaterf(rec, 2, 1, "should not fail")
	assertContinued(t, rec, "Greaterf")
}

func TestGreaterOrEqualf(t *testing.T) {
	rec := newRecorder(t)
	require.GreaterOrEqualf(rec, 1, 2, "%d < %d", 1, 2)
	assertStopped(t, rec, "GreaterOrEqualf")

	rec = newRecorder(t)
	require.GreaterOrEqualf(rec, 2, 2, "should not fail")
	assertContinued(t, rec, "GreaterOrEqualf")
}

func TestLessf(t *testing.T) {
	rec := newRecorder(t)
	require.Lessf(rec, 2, 1, "%d not less than %d", 2, 1)
	assertStopped(t, rec, "Lessf")

	rec = newRecorder(t)
	require.Lessf(rec, 1, 2, "should not fail")
	assertContinued(t, rec, "Lessf")
}

func TestLessOrEqualf(t *testing.T) {
	rec := newRecorder(t)
	require.LessOrEqualf(rec, 3, 2, "%d > %d", 3, 2)
	assertStopped(t, rec, "LessOrEqualf")

	rec = newRecorder(t)
	require.LessOrEqualf(rec, 2, 2, "should not fail")
	assertContinued(t, rec, "LessOrEqualf")
}

func TestBetweenf(t *testing.T) {
	rec := newRecorder(t)
	require.Betweenf(rec, 10, 1, 5, "%d not in [%d, %d]", 10, 1, 5)
	assertStopped(t, rec, "Betweenf")

	rec = newRecorder(t)
	require.Betweenf(rec, 3, 1, 5, "should not fail")
	assertContinued(t, rec, "Betweenf")
}

func TestPositivef(t *testing.T) {
	rec := newRecorder(t)
	require.Positivef(rec, -1, "got %d", -1)
	assertStopped(t, rec, "Positivef")

	rec = newRecorder(t)
	require.Positivef(rec, 1, "should not fail")
	assertContinued(t, rec, "Positivef")
}

func TestNegativef(t *testing.T) {
	rec := newRecorder(t)
	require.Negativef(rec, 1, "got %d", 1)
	assertStopped(t, rec, "Negativef")

	rec = newRecorder(t)
	require.Negativef(rec, -1, "should not fail")
	assertContinued(t, rec, "Negativef")
}

func TestInDeltaf(t *testing.T) {
	rec := newRecorder(t)
	require.InDeltaf(rec, 1.0, 2.0, 0.5, "delta: %v", 1.0)
	assertStopped(t, rec, "InDeltaf")

	rec = newRecorder(t)
	require.InDeltaf(rec, 1.0, 1.1, 0.5, "should not fail")
	assertContinued(t, rec, "InDeltaf")
}

// -- JSON --

func TestEqualJSONf(t *testing.T) {
	rec := newRecorder(t)
	require.EqualJSONf(rec, map[string]int{"a": 1}, `{"a":2}`, "json: %s", "mismatch")
	assertStopped(t, rec, "EqualJSONf")

	rec = newRecorder(t)
	require.EqualJSONf(rec, map[string]int{"a": 1}, `{"a":1}`, "should not fail")
	assertContinued(t, rec, "EqualJSONf")
}

func TestJSONEqualf(t *testing.T) {
	rec := newRecorder(t)
	require.JSONEqualf(rec, `{"a":1}`, `{"a":2}`, "json: %s", "mismatch")
	assertStopped(t, rec, "JSONEqualf")

	rec = newRecorder(t)
	require.JSONEqualf(rec, `{"a":1,"b":2}`, `{"b":2,"a":1}`, "should not fail")
	assertContinued(t, rec, "JSONEqualf")
}

// -- Files --

func TestFileExistsf(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "testfile-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	rec := newRecorder(t)
	require.FileExistsf(rec, f.Name()+".nonexistent", "file: %s", "missing")
	assertStopped(t, rec, "FileExistsf")

	rec = newRecorder(t)
	require.FileExistsf(rec, f.Name(), "should not fail")
	assertContinued(t, rec, "FileExistsf")
}

func TestDirExistsf(t *testing.T) {
	dir := t.TempDir()

	rec := newRecorder(t)
	require.DirExistsf(rec, dir+"/nonexistent", "dir: %s", "missing")
	assertStopped(t, rec, "DirExistsf")

	rec = newRecorder(t)
	require.DirExistsf(rec, dir, "should not fail")
	assertContinued(t, rec, "DirExistsf")
}

// -- Time --

func TestWithinDurationf(t *testing.T) {
	now := time.Now()

	rec := newRecorder(t)
	require.WithinDurationf(rec, now, now.Add(2*time.Second), time.Second, "delta: %s", "too large")
	assertStopped(t, rec, "WithinDurationf")

	rec = newRecorder(t)
	require.WithinDurationf(rec, now, now.Add(500*time.Millisecond), time.Second, "should not fail")
	assertContinued(t, rec, "WithinDurationf")
}
