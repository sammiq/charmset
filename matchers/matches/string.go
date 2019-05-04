package matches

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/sammiq/charmset"
)

// Pattern returns a matcher that checks whether a string matches a regular expression pattern
func Pattern(regex string) *charmset.MatcherType {
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
