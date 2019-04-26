package core

import (
	"fmt"
	"reflect"
)

func isEqual(expected interface{}, actual interface{}) (err error) {
	//check for pointer or nil equality
	match := actual == expected
	if !match {
		// Attempt comparison after type conversion if required
		actualType := reflect.TypeOf(actual)
		expectedValue := reflect.ValueOf(expected)
		if actualType != nil && expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
			match = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
		}
	}
	if !match {
		// Attempt comparison without type conversion as a backup
		match = reflect.DeepEqual(expected, actual)
	}
	if !match {
		err = fmt.Errorf("value was <%v>", actual)
	}
	return err
}
