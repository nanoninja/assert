# Assert

A lightweight, type-safe assertion library for Go tests — familiar API, zero dependencies.

[![Golang](https://img.shields.io/badge/Go-%3E%3D%201.21-%2300ADD8.svg)](https://go.dev/)
[![Tests](https://github.com/nanoninja/assert/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/nanoninja/assert/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/nanoninja/assert/branch/main/graph/badge.svg)](https://codecov.io/gh/nanoninja/assert)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanoninja/assert)](https://goreportcard.com/report/github.com/nanoninja/assert)
[![Go Reference](https://pkg.go.dev/badge/github.com/nanoninja/assert.svg)](https://pkg.go.dev/github.com/nanoninja/assert)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

## Features

- Familiar API — same conventions as testify (`expected` first, `actual` second)
- Type-safe using Go generics
- Two packages: `assert` (continues on failure) and `require` (stops immediately)
- No external dependencies
- Clear, structured error messages

## Requirements

Go 1.21 or higher.

## Installation

```bash
go get github.com/nanoninja/assert
```

## assert vs require

The library provides two packages with identical APIs but different failure behavior:

| | `assert` | `require` |
|---|---|---|
| On failure | `t.Error` — test continues | `t.Fatal` — test stops immediately |
| Use for | Validating multiple values at once | Preconditions that must pass |

```go
import (
    "github.com/nanoninja/assert"
    "github.com/nanoninja/assert/require"
)

func TestCreateUser(t *testing.T) {
    user, err := CreateUser("alice")

    require.NoError(t, err)    // stops if err != nil — user would be nil
    require.NotNil(t, user)    // stops if nil — next assertions would panic

    assert.Equal(t, "alice", user.Name)    // continues on failure
    assert.Equal(t, "active", user.Status) // sees all failures in one run
    assert.NotEmpty(t, user.ID)
}
```

## Usage

### Comparisons

```go
assert.Equal(t, expected, actual)
assert.NotEqual(t, expected, actual)

assert.True(t, IsValid())
assert.False(t, HasErrors())

assert.Nil(t, err)
assert.NotNil(t, user)

assert.IsType(t, MyStruct{}, value)
assert.Implements(t, (*io.Reader)(nil), myReader)
```

### Error Handling

```go
result, err := Process()
assert.NoError(t, err)
assert.Error(t, err)

assert.EqualError(t, err, "file not found")

var ErrNotFound = errors.New("not found")
wrappedErr := fmt.Errorf("fetch failed: %w", ErrNotFound)
assert.ErrorIs(t, wrappedErr, ErrNotFound)

var validationErr *ValidationError
assert.ErrorAs(t, err, &validationErr)

assert.Panics(t, fn, "expected panic message")
assert.NotPanics(t, fn)
assert.PanicsWithError(t, fn, "expected error message")
```

### Collections

```go
users := []string{"alice", "bob", "charlie"}

assert.Contains(t, users, "alice")
assert.NotContains(t, users, "dave")
assert.ElementsMatch(t, []int{1, 2, 3}, []int{3, 1, 2})

assert.Subset(t, users, []string{"alice", "bob"})
assert.NotSubset(t, users, []string{"alice", "dave"})

assert.Empty(t, list)
assert.NotEmpty(t, users)
assert.Len(t, users, 3)

userMap := map[string]User{"alice": aliceUser}
assert.HasKey(t, userMap, "alice")
```

### Strings

```go
assert.StringContains(t, response.Body, "success")
assert.HasPrefix(t, filename, "test-")
assert.HasSuffix(t, filename, ".go")
assert.MatchRegexp(t, email, `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
```

### Numeric Comparisons

```go
assert.Greater(t, count, 0)
assert.GreaterOrEqual(t, age, 18)
assert.Less(t, temperature, 100)
assert.LessOrEqual(t, price, maxPrice)
assert.Between(t, value, min, max)

assert.Positive(t, score)
assert.Negative(t, balance)

assert.InDelta(t, 0.3, 0.1+0.2, 0.0001) // float tolerance
```

### JSON

```go
// Compares JSON semantically — key order and whitespace are ignored
assert.JSONEqual(t, `{"a":1,"b":2}`, `{"b":2,"a":1}`)
```

### Time

```go
assert.WithinDuration(t, expected, actual, time.Second)
```

## Error Messages

Failed assertions produce structured output showing file, line, expected and actual values:

```
file_test.go:42
 Message: optional custom message
Expected: (int) 5
  Actual: (int) 4
```

## License

BSD 3-Clause. See [LICENSE](LICENSE).
