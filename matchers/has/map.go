package has

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/sammiq/charmset"
	"github.com/sammiq/charmset/matchers/is"
)

func everyKeyInMapMatch(matcher charmset.Matcher, actual reflect.Value) (err error) {
	if actual.Len() == 0 {
		return errors.New("was empty")
	}
	iter := reflect.ValueOf(actual).MapRange()
	for iter.Next() {
		key := iter.Key()
		err = matcher.Match(key.Interface())
		if err != nil {
			return fmt.Errorf("contained an key where %s", err)
		}
	}
	return nil
}

func anyKeyInMapMatch(matcher charmset.Matcher, actual reflect.Value) (err error) {
	if actual.Len() == 0 {
		return errors.New("was empty")
	}
	errs := make([]string, 0, actual.Len())
	iter := reflect.ValueOf(actual).MapRange()
	for iter.Next() {
		key := iter.Key()
		err = matcher.Match(key.Interface())
		if err == nil {
			return nil
		}
		errs = append(errs, err.Error())
	}
	if len(errs) == 1 {
		return fmt.Errorf("contained a key where %s", errs[0])
	}
	return fmt.Errorf("no key matched where [\n          %s\n          ]", strings.Join(errs, ",\n          "))
}

// EveryKeyMatching returns a matcher that checks whether every key in a map
// matches a given matcher. Returns early if an item does not match.
func EveryKeyMatching(matcher charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("any key to have %s", matcher.Description()),
		func(actual interface{}) error {
			actualValue := reflect.ValueOf(actual)
			switch actualValue.Kind() {
			case reflect.Map:
				return everyKeyInMapMatch(matcher, actualValue)
			default:
				return errors.New("was not a map")
			}
		},
	)
}

// AnyKeyMatching returns a matcher that checks whether any key in a map
// matches a given matcher. Returns early if an item matches.
func AnyKeyMatching(matcher charmset.Matcher) *charmset.MatcherType {
	return charmset.NewMatcher(
		fmt.Sprintf("any key to have %s", matcher.Description()),
		func(actual interface{}) error {
			actualValue := reflect.ValueOf(actual)
			switch actualValue.Kind() {
			case reflect.Map:
				return anyKeyInMapMatch(matcher, actualValue)
			default:
				return errors.New("was not a map")
			}
		},
	)
}

// Key returns a matcher that checks whether any key in a map
// matches a given value. Returns early if an item matches.
func Key(expected interface{}) *charmset.MatcherType {
	return AnyKeyMatching(is.EqualTo(expected))
}

// KeyIn returns a matcher that checks whether any key in a map
// matches any of a set of given values. Returns early if an item matches.
func KeyIn(expected ...interface{}) *charmset.MatcherType {
	return AnyKeyMatching(is.OneOf(expected...))
}

// Keys returns a matcher that checks whether all elements of a given set of values
// are keys in a map. Returns early if a match is not found.
func Keys(expected ...interface{}) *charmset.MatcherType {
	if len(expected) == 0 {
		//panic as there is no reason to continue the test if expected is invalid at construction
		panic("will never match an empty set of items")
	}
	return charmset.NewMatcher(
		fmt.Sprintf("keys equal to <%v> in any order", expected),
		func(actual interface{}) error {
			actualValue := reflect.ValueOf(actual)
			switch actualValue.Kind() {
			case reflect.Map:
				for _, ex := range expected {
					if err := anyKeyInMapMatch(is.EqualTo(ex), actualValue); err != nil {
						return err
					}
				}
				return nil
			default:
				return errors.New("was not a map")
			}
		},
	)
}
