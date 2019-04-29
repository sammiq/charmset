package internal

import (
	"fmt"
	"reflect"
)

// The rules for Type.ConvertibleTo and Value.Convert in the reflect package are not the same
// as in the specification (https://golang.org/ref/spec#Conversions) as they allow truncating
// conversions to take place. This is my attempt to stop the truncation from taking place.
func convertValue(expectedValue reflect.Value, actualType reflect.Type) interface{} {
	expectedType := expectedValue.Type()
	expectedKind := expectedType.Kind()
	convertedValue := expectedValue.Convert(actualType)
	//this relies on order of this enumeration to get the numeric types only
	if expectedKind > reflect.Invalid && expectedKind < reflect.Array {
		//attempt to convert back to ensure we lost no information
		if actualType.ConvertibleTo(expectedType) {
			restoredValue := convertedValue.Convert(expectedType)
			if restoredValue.IsValid() &&
				restoredValue.Interface() == expectedValue.Interface() {
				// we converted back and lost no information to allow conversion
				return convertedValue.Interface()
			}
		}
		// we do not allow conversion of one-way conversions
		return expectedValue
	}
	//not a numeric type so allow the conversion
	return convertedValue.Interface()
}

func shouldConvert(expectedValue reflect.Value, actualType reflect.Type) bool {
	expectedType := expectedValue.Type()
	return expectedType != nil &&
		actualType != nil &&
		expectedType != actualType &&
		expectedType.ConvertibleTo(actualType)

}

// Equal tests for equality while allowing easy conversions to take place
// such as converting between integer types before checking equality
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

// EqualAny tests that a value is equal to one of an array of other values
// using the rules in Equal above. Returns early if a match is found.
func EqualAny(expected []interface{}, actual interface{}) (err error) {
	for _, ex := range expected {
		err = Equal(ex, actual)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("value was <%v>", actual)
}
