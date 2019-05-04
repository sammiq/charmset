package internal

import (
	"errors"
	"fmt"
)

// MockTester provides an easy way of mocking the Tester interface
// and counting the times errors are passed to the interface.
type MockTester struct {
	ErrorCount  int
	FatalCount  int
	LastMessage string
}

// Logf mocks the Tester Logf method
func (x *MockTester) Logf(format string, args ...interface{}) {
	x.LastMessage = fmt.Sprintf(format, args...)
}

// Fail mocks the Tester Fail method
func (x *MockTester) Fail() {
	x.ErrorCount++
}

// FailNow mocks the Tester FailNow method
func (x *MockTester) FailNow() {
	x.FatalCount++
}

// MockMatcher provides an easy way of mocking the Matcher interface,
// returning a fixed result, and counting the times the matcher is called.
type MockMatcher struct {
	Matches   bool
	CallCount int
}

// Description mocks the Matcher Description method
func (x *MockMatcher) Description() string {
	return "might match"
}

// Match mocks the Matcher Match method
func (x *MockMatcher) Match(actual interface{}) error {
	x.CallCount++
	if x.Matches {
		return nil
	}
	return errors.New("did not match")
}
