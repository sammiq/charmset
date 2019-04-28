package internal

import (
	"errors"
	"fmt"
)

type MockTester struct {
	ErrorCount  int
	FatalCount  int
	LastMessage string
}

func (x *MockTester) Logf(format string, args ...interface{}) {
	x.LastMessage = fmt.Sprintf(format, args...)
}

func (x *MockTester) Fail() {
	x.ErrorCount++
}

func (x *MockTester) FailNow() {
	x.FatalCount++
}

type MockMatcher struct {
	Matches   bool
	CallCount int
}

func (x *MockMatcher) Description() string {
	return "might match"
}

func (x *MockMatcher) Match(actual interface{}) error {
	x.CallCount++
	if x.Matches {
		return nil
	}
	return errors.New("did not match")
}
