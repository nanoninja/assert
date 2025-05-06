# Assert

A lightweight and type-safe assertion package for Go tests that prioritizes readability and clear error messages.

[![Golang](https://img.shields.io/badge/Go-%3E%3D%201.18-%2300ADD8.svg)](https://go.dev/)
[![Tests](https://github.com/nanoninja/assert/actions/workflows/tests.yml/badge.svg)](https://github.com/nanoninja/assert/actions)
[![codecov](https://codecov.io/gh/nanoninja/assert/branch/main/graph/badge.svg)](https://codecov.io/gh/nanoninja/assert)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanoninja/assert)](https://goreportcard.com/report/github.com/nanoninja/assert)
[![Go Reference](https://pkg.go.dev/badge/github.com/nanoninja/assert.svg)](https://pkg.go.dev/github.com/nanoninja/assert)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

## Requirements

- Go 1.18 or higher (required for generics support)

## Features

- Type-safe assertions using Go generics
- Clear and detailed error messages
- Seamless integration with Go's testing package
- No external dependencies
- Support for custom error messages
- File and line information for failed assertions

## Installation

```bash
go get github.com/nanoninja/assert
```

## Basic Usage

```go
func TestExample(t *testing.T) {
    // Basic equality
    assert.Equal(t, Calculate(2, 3), 5)

    // Error checking
    err := Process()
    assert.EqualError(t, err, ErrExpected)

    // Collection operations
    users := []string{"alice", "bob"}

    assert.Contains(t, users, "alice")
    assert.Len(t, users, 2)

    // Numeric comparisons
    assert.Greater(t, GetCount(), 0)
    assert.Between(t, GetValue(), 1, 10)

    // Boolean assertions
    assert.True(t, IsValid())
    assert.False(t, HasErrors())

    // Nil checks
    assert.NotNil(t, GetUser())
}
```

## Available Assertions

### Basic Comparisons

```go
assert.Equal(t, Calculate(2, 3), 5)
assert.NotEqual(t, user1, user2)

assert.True(t, IsValid())
assert.False(t, HasErrors())
```

### Error Handling

The package provides comprehensive error handling assertions that work with Go's error wrapping mechanisms:

```go
// Basic error comparison (now for checking the presence of an error)
err := Process()
assert.Error(t, err, "Expected an error")

// Check for the absence of an error
result, err := GetData()
assert.NoError(t, err, "Did not expect an error")


// Assert that the actual error is equal to the expected error (string comparison)
expectedErr := errors.New("file not found")
actualErr := OpenFile("nonexistent.txt")
assert.EqualError(t, actualErr, expectedErr)


// Working with wrapped errors
var ErrNotFound = errors.New("not found")
wrappedErr := fmt.Errorf("failed to fetch: %w", ErrNotFound)

assert.ErrorIs(t, wrappedErr, ErrNotFound)

// Working with custom error types
type ValidationError struct {
    Field string
    Code  int
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed: %s (code: %d)", e.Field, e.Code)
}

// Later in tests...
err := Validate()
var validationErr *ValidationError

assert.ErrorAs(t, err, &validationErr)
```

Different error assertion functions serve different purposes:

* `Error()`: Direct comparison of error values
* `ErrorIs()`: Checks if an error matches a specific error value anywhere in its chain of wrapped errors
* `ErrorAs()`: Checks if an error (or any error it wraps) matches a specific error type and extracts it

### Collection Operations

```go
users := []string{"alice", "bob", "charlie"}

assert.Contains(t, users, "alice")
assert.NotContains(t, users, "dave")

assert.Empty(t, list)
assert.Len(t, users, 3)

userMap := map[string]User{"alice": aliceUser}

assert.HasKey(t, userMap, "alice")
```

### String Operations

```go
assert.StringContains(t, response.Body, "success")

assert.HasPrefix(t, filename, "test-")
assert.HasSuffix(t, filename, ".txt")

assert.MatchRegexp(t, email, `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
```

### Numeric Comparisons

```go
assert.Greater(t, count, 0)
assert.GreaterOrEqual(t, age, 18)

assert.LessOrEqual(t, temperature, 100)
assert.Between(t, value, min, max)
```

## Error Messages

When an assertion fails, you get clear error messages that include:

* File and line number where the failure occurred
* Expected and actual values with their types
* Optional custom message
* Any relevant context for the comparison

Example of an error message:

```bash
file.go:42
Expected: (int) 5
  Actual: (int) 4
```

## License

This project is licensed under the BSD 3-Clause License.

It allows you to:
- Use the software commercially
- Modify the software
- Distribute the software
- Place warranty on the software
- Use the software privately

The only requirements are:
- Include the copyright notice
- Include the license text
- Not use the author's name to promote derived products without permission

For more details, see the LICENSE file in the project repository.