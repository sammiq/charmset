package has

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/sammiq/matchers"
	"github.com/sammiq/matchers/core/is"
)

func everyInSliceMatch(matcher matchers.Matcher, actual reflect.Value) (err error) {
	for i := 0; i < actual.Len(); i++ {
		value := actual.Index(i).Interface()
		err = matcher.Match(value)
		if err != nil {
			return fmt.Errorf("contained an item where %s", err)
		}
	}
	return nil
}

func anyInSliceMatch(matcher matchers.Matcher, actual reflect.Value) (err error) {
	if actual.Len() == 0 {
		return errors.New("was empty")
	}
	errs := make([]string, 0, actual.Len())
	for i := 0; i < actual.Len(); i++ {
		value := actual.Index(i).Interface()
		err = matcher.Match(value)
		if err == nil {
			return nil
		}
		errs = append(errs, err.Error())
	}
	if len(errs) == 1 {
		return fmt.Errorf("contained an item where %s", errs[0])
	}
	return fmt.Errorf("no item matched [\n          %s\n          ]", strings.Join(errs, ",\n          "))
}

func EveryItemMatching(matcher matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("every item to have %s", matcher.Description()),
		func(actual interface{}) error {
			actualValue := reflect.ValueOf(actual)
			switch actualValue.Kind() {
			case reflect.Array, reflect.Slice:
				return everyInSliceMatch(matcher, actualValue)
			default:
				panic("invalid value type to check")
			}
		},
	)
}

func AnyItemMatching(matcher matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("any item to have %s", matcher.Description()),
		func(actual interface{}) error {
			actualValue := reflect.ValueOf(actual)
			switch actualValue.Kind() {
			case reflect.Array, reflect.Slice:
				return anyInSliceMatch(matcher, actualValue)
			default:
				panic("invalid value type to check")
			}
		},
	)
}

func AnyItem(expected interface{}) *matchers.MatcherType {
	return AnyItemMatching(is.EqualTo(expected))
}
