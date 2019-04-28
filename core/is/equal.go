package is

import (
	"fmt"

	"github.com/sammiq/matchers"
	"github.com/sammiq/matchers/internal"
)

func EqualTo(expected interface{}) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("value equal to <%v>", expected),
		func(actual interface{}) error { return internal.Equal(expected, actual) },
	)
}

func OneOf(expected ...interface{}) *matchers.MatcherType {
	if len(expected) == 0 {
		//panic as there is no reason to continue the test if expected is invalid at construction
		panic("will never match empty slice")
	}
	return matchers.NewMatcher(
		fmt.Sprintf("value equal to any of <%v>", expected),
		func(actual interface{}) (err error) {
			return internal.EqualAny(expected, actual)
		},
	)
}

func NotEqualTo(expected interface{}) *matchers.MatcherType {
	return Not(EqualTo(expected))
}

func Nil() *matchers.MatcherType {
	return EqualTo(nil)
}

func NotNil() *matchers.MatcherType {
	return Not(EqualTo(nil))
}
