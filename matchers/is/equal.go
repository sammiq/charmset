package is

import (
	"fmt"

	"github.com/sammiq/charmset"
	"github.com/sammiq/charmset/internal"
)

// EqualTo returns a matcher that checks whether a value is equal to an expected value.
// There are some small type conversions allowed of this expected value, and numbers can be
// converted provided that no truncation occurs that would change the value from the intent.
func EqualTo(expected interface{}) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("value equal to <%v>", expected),
		func(actual interface{}) error { return internal.Equal(expected, actual) },
	)
}

// OneOf returns a matcher that checks whether a value is in a set of expected values.
// As with EqualTo, there are some small type conversions allowed of this expected value,
// and numbers can be converted provided that no truncation occurs that would change the
// value from the intent.
func OneOf(expected ...interface{}) *charmset.MatcherType {
	if len(expected) == 0 {
		//panic as there is no reason to continue the test if expected is invalid at construction
		panic("will never match an empty set of items")
	}
	return charmset.NewMatcher(
		fmt.Sprintf("value equal to any of <%v>", expected),
		func(actual interface{}) (err error) {
			return internal.EqualAny(expected, actual)
		},
	)
}

// NotEqualTo returns a matcher that checks whether a value is not equal to an expected value.
// Equivalent to Not(EqualTo(..))
func NotEqualTo(expected interface{}) *charmset.MatcherType {
	return Not(EqualTo(expected))
}

// Nil returns a matcher that checks whether a value is equal to nil.
// Broadly equivalent to EqualTo(nil) but more efficient.
func Nil() *charmset.MatcherType {
	return charmset.NewMatcher(
		"value equal to <<nil>>",
		func(actual interface{}) error {
			if actual != nil {
				return fmt.Errorf("value was <%v>", actual)
			}
			return nil
		},
	)
}

// NotNil returns a matcher that checks whether a value is not equal to nil.
// Equivalent to Not(Nil())
func NotNil() *charmset.MatcherType {
	return Not(Nil())
}
