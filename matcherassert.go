package matchers

//abstracts away the struct used in the testing package
type Tester interface {
	Logf(format string, args ...interface{})
	Fail()
	FailNow()
}

type MatcherAssert struct {
	Test     Tester
	FailFast bool
}

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

func AssertThat(tester Tester, actual interface{}, matcher Matcher) {
	assert := &MatcherAssert{Test: tester}
	assert.That(actual, matcher)
}
