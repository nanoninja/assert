// Copyright 2025 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package require_test

import (
	"fmt"
	"testing"

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

func TestRequire_StopsOnFailure(t *testing.T) {
	t.Run("Equal stops on failure", func(t *testing.T) {
		rec := newRecorder(t)
		require.Equal(rec, 1, 2)
		if !rec.failNowCalled {
			t.Error("require.Equal() did not call FailNow on failure")
		}
	})

	t.Run("Equal does not stop on success", func(t *testing.T) {
		rec := newRecorder(t)
		require.Equal(rec, 1, 1)
		if rec.failNowCalled {
			t.Error("require.Equal() called FailNow on success")
		}
	})

	t.Run("NoError stops on failure", func(t *testing.T) {
		rec := newRecorder(t)
		require.NoError(rec, fmt.Errorf("unexpected error"))
		if !rec.failNowCalled {
			t.Error("require.NoError() did not call FailNow on failure")
		}
	})

	t.Run("NoError does not stop on success", func(t *testing.T) {
		rec := newRecorder(t)
		require.NoError(rec, nil)
		if rec.failNowCalled {
			t.Error("require.NoError() called FailNow on success")
		}
	})

	t.Run("Nil stops on failure", func(t *testing.T) {
		rec := newRecorder(t)
		require.Nil(rec, "not nil")
		if !rec.failNowCalled {
			t.Error("require.Nil() did not call FailNow on failure")
		}
	})
}
