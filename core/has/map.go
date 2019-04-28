package has

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/sammiq/matchers"
	"github.com/sammiq/matchers/core/is"
)

func everyKeyInMapMatch(matcher matchers.Matcher, actual reflect.Value) (err error) {
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

func anyKeyInMapMatch(matcher matchers.Matcher, actual reflect.Value) (err error) {
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

func EveryKeyMatching(matcher matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
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

func AnyKeyMatching(matcher matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
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

func Key(expected interface{}) *matchers.MatcherType {
	return AnyKeyMatching(is.EqualTo(expected))
}

func Keys(expected ...interface{}) *matchers.MatcherType {
	return AnyKeyMatching(is.OneOf(expected...))
}
