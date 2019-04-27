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
	Fail      bool
	CallCount int
}

func (x *MockMatcher) Description() string {
	return "should pass"
}

func (x *MockMatcher) Match(actual interface{}) error {
	x.CallCount++
	if x.Fail {
		return errors.New("did fail")
	}
	return nil
}
