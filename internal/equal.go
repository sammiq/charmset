package internal

import (
	"fmt"
	"reflect"
)

// the rules for Type.ConvertibleTo and Value.Convert in the reflect package are not the same
// as in the specification (https://golang.org/ref/spec#Conversions) as they allow truncating
// conversions to take place. This is my attempt to stop the truncation from taking place.
func convertValue(expectedValue reflect.Value, actualType reflect.Type) interface{} {
	expectedType := expectedValue.Type()
	expectedKind := expectedType.Kind()
	convertedValue := expectedValue.Convert(actualType)
	//this relies on order of this enumeration to get the numeric types only
	if expectedKind > reflect.Invalid && expectedKind < reflect.Array {
		restoredValue := convertedValue.Convert(expectedType)
		//failed to convert back, or we are not equal (truncated), do not allow the conversion
		if !restoredValue.IsValid() ||
			restoredValue.Interface() != expectedValue.Interface() {
			return expectedValue
		}
	}
	return convertedValue.Interface()
}

func shouldConvert(expectedValue reflect.Value, actualType reflect.Type) bool {
	expectedType := expectedValue.Type()
	return expectedType != nil &&
		actualType != nil &&
		expectedType != actualType &&
		expectedType.ConvertibleTo(actualType)

}

func Equal(expected interface{}, actual interface{}) (err error) {
	//check for pointer or nil equality
	match := actual == expected
	if !match {
		// Attempt comparison after type conversion if required
		actualType := reflect.TypeOf(actual)
		expectedValue := reflect.ValueOf(expected)
		if expectedValue.IsValid() && shouldConvert(expectedValue, actualType) {
			match = reflect.DeepEqual(convertValue(expectedValue, actualType), actual)
		} else {
			match = reflect.DeepEqual(expected, actual)
		}
	}
	if !match {
		err = fmt.Errorf("value was <%v>", actual)
	}
	return err
}

func EqualAny(expected []interface{}, actual interface{}) (err error) {
	for _, ex := range expected {
		err = Equal(ex, actual)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("value was <%v>", actual)
}
