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

func (ma *MatcherAssert) AssertThat(actual interface{}, matcher Matcher) {
	if err := matcher.Match(actual); err != nil {
		ma.Test.Logf("Expected: %s\n     but: %s", matcher.Description(), err)
		if ma.FailFast {
			ma.Test.FailNow()
		} else {
			ma.Test.Fail()
		}
	}
}
