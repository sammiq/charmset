package is

import (
	"errors"
	"fmt"
	"regexp"
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

func StringContaining(subString string) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("string containing '%s'", subString),
		func(actual interface{}) error {
			return stringOp(actual, subString, strings.Contains, "contain")
		},
	)
}

func StringStartingWith(prefix string) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("string starting with '%s'", prefix),
		func(actual interface{}) error {
			return stringOp(actual, prefix, strings.HasPrefix, "start with")
		},
	)
}

func StringEndingWith(suffix string) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("string ending with '%s'", suffix),
		func(actual interface{}) error {
			return stringOp(actual, suffix, strings.HasSuffix, "end with")
		},
	)
}

func StringMatchingPattern(regex string) *charmset.MatcherType {
	r, err := regexp.Compile(regex)
	if err != nil {
		//panic as there is no reason to continue the test if pattern is invalid at construction
		panic(fmt.Sprintf("failed to compile regex: %s", err))
	}
	return charmset.NewMatcher(
		fmt.Sprintf("string matching pattern '%s'", regex),
		func(actual interface{}) error {
			actualString, ok := actual.(string)
			if ok {
				if r.MatchString(actualString) {
					return nil
				}
				return fmt.Errorf("string '%s' did not match pattern '%s'", actualString, regex)
			}
			return errors.New("was not a string")
		},
	)
}
