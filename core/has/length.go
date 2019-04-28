package has

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sammiq/matchers"
)

func Length(length int) *matchers.MatcherType {
	return matchers.NewMatcher(
		fmt.Sprintf("length equal to %d", length),
		func(actual interface{}) error {
			actualValue := reflect.ValueOf(actual)
			switch actualValue.Kind() {
			case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
				actualLength := actualValue.Len()
				if actualLength == length {
					return nil
				}
				return fmt.Errorf("length was <%v>", actualLength)
			default:
				return errors.New("was not a array, slice, map or string")
			}
		},
	)
}
