package is

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sammiq/charmset"
)

// Empty returns a matcher that checks whether len(...) of an array, slice, map or string  is 0.
func Empty() *charmset.MatcherType {
	return charmset.NewMatcher(
		"length equal to <0>",
		func(actual interface{}) error {
			actualValue := reflect.ValueOf(actual)
			switch actualValue.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
				actualLength := actualValue.Len()
				if actualLength == 0 {
					return nil
				}
				return fmt.Errorf("length was <%v>", actualLength)
			default:
				return errors.New("was not a array, slice, map or string")
			}
		},
	)
}
