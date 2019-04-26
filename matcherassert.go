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

func (ma *MatcherAssert) AssertThat(reason string, actual interface{}, matcher Matcher) {
	err := matcher.Match(actual)
	if err != nil {
		ma.Test.Logf("%s\nExpected: %s\nbut: %s", reason, matcher.Description(), err)
		if ma.FailFast {
			ma.Test.FailNow()
		} else {
			ma.Test.Fail()
		}
	}
}
