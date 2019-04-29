package is

import (
	"fmt"

	"github.com/sammiq/charmset"
	"github.com/sammiq/charmset/internal"
)

func EqualTo(expected interface{}) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("value equal to <%v>", expected),
		func(actual interface{}) error { return internal.Equal(expected, actual) },
	)
}

func OneOf(expected ...interface{}) *charmset.MatcherType {
	if len(expected) == 0 {
		//panic as there is no reason to continue the test if expected is invalid at construction
		panic("will never match empty slice")
	}
	return charmset.NewMatcher(
		fmt.Sprintf("value equal to any of <%v>", expected),
		func(actual interface{}) (err error) {
			return internal.EqualAny(expected, actual)
		},
	)
}

func NotEqualTo(expected interface{}) *charmset.MatcherType {
	return Not(EqualTo(expected))
}

func Nil() *charmset.MatcherType {
	return EqualTo(nil)
}

func NotNil() *charmset.MatcherType {
	return Not(EqualTo(nil))
}
