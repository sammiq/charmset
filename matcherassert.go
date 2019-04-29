package matchers

// Tester abstracts away the T struct used in the testing package
type Tester interface {
	Logf(format string, args ...interface{})
	Fail()
	FailNow()
}

// MatcherAssert contains all the state data required for assertions
type MatcherAssert struct {
	Test     Tester
	FailFast bool
}

// That is the method to perform test assertions on a value and matchers
func (ma *MatcherAssert) That(actual interface{}, matcher Matcher) {
	if err := matcher.Match(actual); err != nil {
		ma.Test.Logf("Expected: %s\n     but: %s", matcher.Description(), err)
		if ma.FailFast {
			ma.Test.FailNow()
		} else {
			ma.Test.Fail()
		}
	}
}

// AssertThat is a helper method that creates a MatcherAssert and calls That to
// perform a test assertion in simple cases using default parameters
func AssertThat(tester Tester, actual interface{}, matcher Matcher) {
	assert := &MatcherAssert{Test: tester}
	assert.That(actual, matcher)
}
