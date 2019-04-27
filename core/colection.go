package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/sammiq/matchers"
)

func everyInSliceMatch(matcher matchers.Matcher, actual reflect.Value) (err error) {
	for i := 0; i < actual.Len(); i++ {
		value := actual.Index(i).Interface()
		err = matcher.Match(value)
		if err != nil {
			return fmt.Errorf("an item, where %s", err)
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
		return fmt.Errorf("an item, where %s", errs[0])
	}
	return fmt.Errorf("all items, where [\n         %s\n         ]", strings.Join(errs, ",\n         "))
}

func EveryItemMatching(matcher matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("every item has %s", matcher.Description()),
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

func HasItemMatching(matcher matchers.Matcher) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("any item has %s", matcher.Description()),
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

func HasItem(expected interface{}) *matchers.MatcherType {
	return HasItemMatching(EqualTo(expected))
}
