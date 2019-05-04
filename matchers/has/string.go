package has

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sammiq/charmset"
)

func stringOp(actual interface{}, expected string, opFunc func(string, string) bool, opName string) error {
	actualString, ok := actual.(string)
	if ok {
		if opFunc(actualString, expected) {
			return nil
		}
		return fmt.Errorf("string '%s' did not %s '%s'", actualString, opName, expected)
	}
	return errors.New("was not a string")
}

// Prefix returns a matcher that checks whether a string starts with another string
func Prefix(prefix string) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("string starts with '%s'", prefix),
		func(actual interface{}) error {
			return stringOp(actual, prefix, strings.HasPrefix, "start with")
		},
	)
}

// Suffix returns a matcher that checks whether a string ends with another string
func Suffix(suffix string) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("string ends with '%s'", suffix),
		func(actual interface{}) error {
			return stringOp(actual, suffix, strings.HasSuffix, "end with")
		},
	)
}

// Substring returns a matcher that checks whether a string contains another string
func Substring(subString string) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("string containing '%s'", subString),
		func(actual interface{}) error {
			return stringOp(actual, subString, strings.Contains, "contain")
		},
	)
}
