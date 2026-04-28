// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assert

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

// mockT captures whether Errorf was called and the message produced.
type mockT struct {
	testing.TB
	failed  bool
	message string
}

func (m *mockT) Helper() {}

func (m *mockT) Error(args ...any) {
	m.failed = true
	m.message = fmt.Sprint(args...)
}

func (m *mockT) Errorf(format string, args ...any) {
	m.failed = true
	m.message = fmt.Sprintf(format, args...)
}

func newMock(t *testing.T) *mockT {
	t.Helper()
	return &mockT{TB: t}
}

// assertFormattedMessage verifies that the formatted message appears in the output.
func assertFormattedMessage(t *testing.T, m *mockT, substr string) {
	t.Helper()
	if !strings.Contains(m.message, substr) {
		t.Errorf("expected message to contain %q, got: %s", substr, m.message)
	}
}

// -- Comparisons --

func TestEqualf(t *testing.T) {
	m := newMock(t)
	Equalf(m, 1, 2, "values: %d and %d", 1, 2)
	assertFormattedMessage(t, m, "values: 1 and 2")

	m = newMock(t)
	Equalf(m, 1, 1, "should not fail")
	if m.failed {
		t.Error("Equalf() failed on equal values")
	}
}

func TestNotEqualf(t *testing.T) {
	m := newMock(t)
	NotEqualf(m, 1, 1, "got %d", 1)
	assertFormattedMessage(t, m, "got 1")

	m = newMock(t)
	NotEqualf(m, 1, 2, "should not fail")
	if m.failed {
		t.Error("NotEqualf() failed on different values")
	}
}

func TestTruef(t *testing.T) {
	m := newMock(t)
	Truef(m, false, "expected %s", "true")
	assertFormattedMessage(t, m, "expected true")

	m = newMock(t)
	Truef(m, true, "should not fail")
	if m.failed {
		t.Error("Truef() failed on true")
	}
}

func TestFalsef(t *testing.T) {
	m := newMock(t)
	Falsef(m, true, "expected %s", "false")
	assertFormattedMessage(t, m, "expected false")

	m = newMock(t)
	Falsef(m, false, "should not fail")
	if m.failed {
		t.Error("Falsef() failed on false")
	}
}

func TestNilf(t *testing.T) {
	m := newMock(t)
	Nilf(m, "not nil", "got %q", "not nil")
	assertFormattedMessage(t, m, `got "not nil"`)

	m = newMock(t)
	Nilf(m, nil, "should not fail")
	if m.failed {
		t.Error("Nilf() failed on nil")
	}
}

func TestNotNilf(t *testing.T) {
	m := newMock(t)
	NotNilf(m, nil, "value is %s", "nil")
	assertFormattedMessage(t, m, "value is nil")

	m = newMock(t)
	NotNilf(m, "ok", "should not fail")
	if m.failed {
		t.Error("NotNilf() failed on non-nil")
	}
}

func TestIsTypef(t *testing.T) {
	m := newMock(t)
	IsTypef(m, 0, "str", "type mismatch: %s", "int vs string")
	assertFormattedMessage(t, m, "type mismatch: int vs string")

	m = newMock(t)
	IsTypef(m, 0, 1, "should not fail")
	if m.failed {
		t.Error("IsTypef() failed on same type")
	}
}

func TestImplementsf(t *testing.T) {
	m := newMock(t)
	Implementsf(m, (*error)(nil), "not an error", "check %s", "interface")
	assertFormattedMessage(t, m, "check interface")

	m = newMock(t)
	Implementsf(m, (*error)(nil), errors.New("ok"), "should not fail")
	if m.failed {
		t.Error("Implementsf() failed on valid implementation")
	}
}

// -- Errors --

func TestErrorf(t *testing.T) {
	m := newMock(t)
	Errorf(m, nil, "expected error: %s", "none")
	assertFormattedMessage(t, m, "expected error: none")

	m = newMock(t)
	Errorf(m, errors.New("err"), "should not fail")
	if m.failed {
		t.Error("Errorf() failed on non-nil error")
	}
}

func TestNoErrorf(t *testing.T) {
	m := newMock(t)
	NoErrorf(m, errors.New("oops"), "unexpected: %s", "oops")
	assertFormattedMessage(t, m, "unexpected: oops")

	m = newMock(t)
	NoErrorf(m, nil, "should not fail")
	if m.failed {
		t.Error("NoErrorf() failed on nil error")
	}
}

func TestEqualErrorf(t *testing.T) {
	m := newMock(t)
	EqualErrorf(m, errors.New("actual"), "expected", "msg: %s", "mismatch")
	assertFormattedMessage(t, m, "msg: mismatch")

	m = newMock(t)
	EqualErrorf(m, errors.New("same"), "same", "should not fail")
	if m.failed {
		t.Error("EqualErrorf() failed on matching error")
	}
}

func TestErrorIsf(t *testing.T) {
	sentinel := errors.New("sentinel")
	m := newMock(t)
	ErrorIsf(m, errors.New("other"), sentinel, "err: %s", "not sentinel")
	assertFormattedMessage(t, m, "err: not sentinel")

	m = newMock(t)
	ErrorIsf(m, fmt.Errorf("wrap: %w", sentinel), sentinel, "should not fail")
	if m.failed {
		t.Error("ErrorIsf() failed on wrapped sentinel")
	}
}

func TestErrorAsf(t *testing.T) {
	var target *os.PathError
	m := newMock(t)
	ErrorAsf(m, errors.New("plain"), &target, "err: %s", "wrong type")
	assertFormattedMessage(t, m, "err: wrong type")

	var numErr *os.PathError
	m = newMock(t)
	ErrorAsf(m, &os.PathError{Op: "open"}, &numErr, "should not fail")
	if m.failed {
		t.Error("ErrorAsf() failed on matching target")
	}
}

// -- Panics --

func TestNotPanicsf(t *testing.T) {
	m := newMock(t)
	NotPanicsf(m, func() { panic("oops") }, "panic: %s", "unexpected")
	assertFormattedMessage(t, m, "panic: unexpected")

	m = newMock(t)
	NotPanicsf(m, func() {}, "should not fail")
	if m.failed {
		t.Error("NotPanicsf() failed on no panic")
	}
}

// -- Collections --

func TestContainsf(t *testing.T) {
	m := newMock(t)
	Containsf(m, []string{"a", "b"}, "c", "missing: %s", "c")
	assertFormattedMessage(t, m, "missing: c")

	m = newMock(t)
	Containsf(m, []string{"a", "b"}, "a", "should not fail")
	if m.failed {
		t.Error("Containsf() failed on existing element")
	}
}

func TestNotContainsf(t *testing.T) {
	m := newMock(t)
	NotContainsf(m, []string{"a", "b"}, "a", "found: %s", "a")
	assertFormattedMessage(t, m, "found: a")

	m = newMock(t)
	NotContainsf(m, []string{"a", "b"}, "c", "should not fail")
	if m.failed {
		t.Error("NotContainsf() failed on absent element")
	}
}

func TestEmptyf(t *testing.T) {
	m := newMock(t)
	Emptyf(m, []int{1}, "len: %d", 1)
	assertFormattedMessage(t, m, "len: 1")

	m = newMock(t)
	Emptyf(m, []int{}, "should not fail")
	if m.failed {
		t.Error("Emptyf() failed on empty slice")
	}
}

func TestNotEmptyf(t *testing.T) {
	m := newMock(t)
	NotEmptyf(m, []int{}, "collection: %s", "empty")
	assertFormattedMessage(t, m, "collection: empty")

	m = newMock(t)
	NotEmptyf(m, []int{1}, "should not fail")
	if m.failed {
		t.Error("NotEmptyf() failed on non-empty slice")
	}
}

func TestLenf(t *testing.T) {
	m := newMock(t)
	Lenf(m, []int{1, 2}, 3, "expected len %d", 3)
	assertFormattedMessage(t, m, "expected len 3")

	m = newMock(t)
	Lenf(m, []int{1, 2}, 2, "should not fail")
	if m.failed {
		t.Error("Lenf() failed on correct length")
	}
}

func TestHasKeyf(t *testing.T) {
	m := newMock(t)
	HasKeyf(m, map[string]int{"a": 1}, "b", "missing key: %s", "b")
	assertFormattedMessage(t, m, "missing key: b")

	m = newMock(t)
	HasKeyf(m, map[string]int{"a": 1}, "a", "should not fail")
	if m.failed {
		t.Error("HasKeyf() failed on existing key")
	}
}

func TestNotHasKeyf(t *testing.T) {
	m := newMock(t)
	NotHasKeyf(m, map[string]int{"a": 1}, "a", "found key: %s", "a")
	assertFormattedMessage(t, m, "found key: a")

	m = newMock(t)
	NotHasKeyf(m, map[string]int{"a": 1}, "b", "should not fail")
	if m.failed {
		t.Error("NotHasKeyf() failed on absent key")
	}
}

func TestElementsMatchf(t *testing.T) {
	m := newMock(t)
	ElementsMatchf(m, []int{1, 2, 3}, []int{1, 2}, "len mismatch: %d vs %d", 3, 2)
	assertFormattedMessage(t, m, "len mismatch: 3 vs 2")

	m = newMock(t)
	ElementsMatchf(m, []int{1, 2, 3}, []int{3, 1, 2}, "should not fail")
	if m.failed {
		t.Error("ElementsMatchf() failed on matching elements")
	}
}

func TestSubsetf(t *testing.T) {
	m := newMock(t)
	Subsetf(m, []string{"a", "b"}, []string{"a", "c"}, "not subset: %s", "c missing")
	assertFormattedMessage(t, m, "not subset: c missing")

	m = newMock(t)
	Subsetf(m, []string{"a", "b", "c"}, []string{"a", "b"}, "should not fail")
	if m.failed {
		t.Error("Subsetf() failed on valid subset")
	}
}

func TestNotSubsetf(t *testing.T) {
	m := newMock(t)
	NotSubsetf(m, []string{"a", "b", "c"}, []string{"a", "b"}, "is subset: %s", "ab")
	assertFormattedMessage(t, m, "is subset: ab")

	m = newMock(t)
	NotSubsetf(m, []string{"a", "b"}, []string{"a", "c"}, "should not fail")
	if m.failed {
		t.Error("NotSubsetf() failed on non-subset")
	}
}

// -- Strings --

func TestStringContainsf(t *testing.T) {
	m := newMock(t)
	StringContainsf(m, "hello world", "xyz", "missing: %s", "xyz")
	assertFormattedMessage(t, m, "missing: xyz")

	m = newMock(t)
	StringContainsf(m, "hello world", "world", "should not fail")
	if m.failed {
		t.Error("StringContainsf() failed on existing substring")
	}
}

func TestHasPrefixf(t *testing.T) {
	m := newMock(t)
	HasPrefixf(m, "hello", "world", "prefix: %s", "world")
	assertFormattedMessage(t, m, "prefix: world")

	m = newMock(t)
	HasPrefixf(m, "hello", "hel", "should not fail")
	if m.failed {
		t.Error("HasPrefixf() failed on valid prefix")
	}
}

func TestHasSuffixf(t *testing.T) {
	m := newMock(t)
	HasSuffixf(m, "hello", "world", "suffix: %s", "world")
	assertFormattedMessage(t, m, "suffix: world")

	m = newMock(t)
	HasSuffixf(m, "hello", "llo", "should not fail")
	if m.failed {
		t.Error("HasSuffixf() failed on valid suffix")
	}
}

func TestMatchRegexpf(t *testing.T) {
	m := newMock(t)
	MatchRegexpf(m, "hello", `^\d+$`, "pattern: %s", `^\d+$`)
	assertFormattedMessage(t, m, `pattern: ^\d+$`)

	m = newMock(t)
	MatchRegexpf(m, "hello", `^[a-z]+$`, "should not fail")
	if m.failed {
		t.Error("MatchRegexpf() failed on matching pattern")
	}
}

// -- Numerics --

func TestGreaterf(t *testing.T) {
	m := newMock(t)
	Greaterf(m, 1, 2, "%d not greater than %d", 1, 2)
	assertFormattedMessage(t, m, "1 not greater than 2")

	m = newMock(t)
	Greaterf(m, 2, 1, "should not fail")
	if m.failed {
		t.Error("Greaterf() failed on greater value")
	}
}

func TestGreaterOrEqualf(t *testing.T) {
	m := newMock(t)
	GreaterOrEqualf(m, 1, 2, "%d < %d", 1, 2)
	assertFormattedMessage(t, m, "1 < 2")

	m = newMock(t)
	GreaterOrEqualf(m, 2, 2, "should not fail")
	if m.failed {
		t.Error("GreaterOrEqualf() failed on equal values")
	}
}

func TestLessf(t *testing.T) {
	m := newMock(t)
	Lessf(m, 2, 1, "%d not less than %d", 2, 1)
	assertFormattedMessage(t, m, "2 not less than 1")

	m = newMock(t)
	Lessf(m, 1, 2, "should not fail")
	if m.failed {
		t.Error("Lessf() failed on lesser value")
	}
}

func TestLessOrEqualf(t *testing.T) {
	m := newMock(t)
	LessOrEqualf(m, 3, 2, "%d > %d", 3, 2)
	assertFormattedMessage(t, m, "3 > 2")

	m = newMock(t)
	LessOrEqualf(m, 2, 2, "should not fail")
	if m.failed {
		t.Error("LessOrEqualf() failed on equal values")
	}
}

func TestBetweenf(t *testing.T) {
	m := newMock(t)
	Betweenf(m, 10, 1, 5, "%d not in [%d, %d]", 10, 1, 5)
	assertFormattedMessage(t, m, "10 not in [1, 5]")

	m = newMock(t)
	Betweenf(m, 3, 1, 5, "should not fail")
	if m.failed {
		t.Error("Betweenf() failed on in-range value")
	}
}

func TestPositivef(t *testing.T) {
	m := newMock(t)
	Positivef(m, -1, "got %d", -1)
	assertFormattedMessage(t, m, "got -1")

	m = newMock(t)
	Positivef(m, 1, "should not fail")
	if m.failed {
		t.Error("Positivef() failed on positive value")
	}
}

func TestNegativef(t *testing.T) {
	m := newMock(t)
	Negativef(m, 1, "got %d", 1)
	assertFormattedMessage(t, m, "got 1")

	m = newMock(t)
	Negativef(m, -1, "should not fail")
	if m.failed {
		t.Error("Negativef() failed on negative value")
	}
}

func TestInDeltaf(t *testing.T) {
	m := newMock(t)
	InDeltaf(m, 1.0, 2.0, 0.5, "delta too large: %v", 1.0)
	assertFormattedMessage(t, m, "delta too large: 1")

	m = newMock(t)
	InDeltaf(m, 1.0, 1.1, 0.5, "should not fail")
	if m.failed {
		t.Error("InDeltaf() failed on in-delta values")
	}
}

// -- JSON --

func TestEqualJSONf(t *testing.T) {
	m := newMock(t)
	EqualJSONf(m, map[string]int{"a": 1}, `{"a":2}`, "json: %s", "mismatch")
	assertFormattedMessage(t, m, "json: mismatch")

	m = newMock(t)
	EqualJSONf(m, map[string]int{"a": 1}, `{"a":1}`, "should not fail")
	if m.failed {
		t.Error("EqualJSONf() failed on matching JSON")
	}
}

func TestJSONEqualf(t *testing.T) {
	m := newMock(t)
	JSONEqualf(m, `{"a":1}`, `{"a":2}`, "json: %s", "mismatch")
	assertFormattedMessage(t, m, "json: mismatch")

	m = newMock(t)
	JSONEqualf(m, `{"a":1}`, `{"a":1}`, "should not fail")
	if m.failed {
		t.Error("JSONEqualf() failed on equal JSON")
	}
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

	m := newMock(t)
	FileExistsf(m, f.Name()+".nonexistent", "file: %s", "missing")
	assertFormattedMessage(t, m, "file: missing")

	m = newMock(t)
	FileExistsf(m, f.Name(), "should not fail")
	if m.failed {
		t.Error("FileExistsf() failed on existing file")
	}
}

func TestDirExistsf(t *testing.T) {
	dir := t.TempDir()

	m := newMock(t)
	DirExistsf(m, dir+"/nonexistent", "dir: %s", "missing")
	assertFormattedMessage(t, m, "dir: missing")

	m = newMock(t)
	DirExistsf(m, dir, "should not fail")
	if m.failed {
		t.Error("DirExistsf() failed on existing dir")
	}
}

// -- Time --

func TestWithinDurationf(t *testing.T) {
	now := time.Now()

	m := newMock(t)
	WithinDurationf(m, now, now.Add(2*time.Second), time.Second, "delta: %s", "too large")
	assertFormattedMessage(t, m, "delta: too large")

	m = newMock(t)
	WithinDurationf(m, now, now.Add(500*time.Millisecond), time.Second, "should not fail")
	if m.failed {
		t.Error("WithinDurationf() failed on in-range duration")
	}
}
