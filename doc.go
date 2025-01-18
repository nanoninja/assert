// Package assert provides a simple yet powerful testing helper library for Go.
// It offers a set of assertion functions that make tests more readable and
// maintainable while providing clear error messages when tests fail.
//
// The package follows Go's testing conventions and works seamlessly with the
// standard testing package. It uses generics to provide type safety while
// maintaining a clean and intuitive API.
//
// Basic Usage:
//
//	func TestExample(t *testing.T) {
//	    result := Calculate(2, 3)
//	    assert.Equals(t, result, 5)
//
//	    user := GetUser()
//	    assert.NotNil(t, user)
//	    assert.True(t, user.IsActive())
//	}
//
// The package includes several categories of assertions:
//
// Basic Comparisons:
//   - Equals/NotEquals: Compare values of any type
//   - True/False: Boolean assertions
//   - Nil/NotNil: Check for nil values
//
// Error Handling:
//   - Error: Compare error values directly
//   - ErrorIs: Check if an error matches a specific error value, even when wrapped
//   - ErrorAs: Check and extract typed errors from error chains
//   - Panics: Test for panic conditions
//
// Collection Operations:
//   - Contains/NotContains: Check if a slice contains (or not) an element
//   - Empty: Verify if a collection is empty
//   - Len: Check collection length
//   - HasKey: Verify map key existence
//
// String Operations:
//   - StringContains: Check string containment
//   - HasPrefix: Verify if a string starts with a prefix
//   - HasSuffix: Verify if a string ends with a suffix
//   - MatchRegexp: Check if a string matches a regular expression pattern
//
// Numeric Comparisons:
//   - Greater: Compare if a value is strictly greater
//   - GreaterOrEqual: Compare if a value is greater or equal
//   - LessOrEqual: Compare if a value is less or equal
//   - Between: Check if a value falls within a range
//
// Each assertion function provides clear error messages that include:
//   - The file and line number where the assertion failed
//   - The expected and actual values
//   - The types of the compared values
//   - An optional custom message
//
// The error messages are designed to be clear and helpful for debugging:
//
//	file.go:42
//	Message: custom error message
//	Expected: (int) 5
//	  Actual: (int) 4
//
// For more information and examples, see the README.md file.
package assert
