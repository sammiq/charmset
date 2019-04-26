package core

import (
	"testing"
)

func assertIsEqualMatched(t *testing.T, expected interface{}, actual interface{}, expectMatch bool) {
	err := isEqual(expected, actual)
	if (err == nil) != expectMatch {
		t.Errorf("Expected err to be returned when match is %v, expected is %v and actual is %v", expectMatch, expected, actual)
	}
}
func TestIsEqualMatches(t *testing.T) {
	assertIsEqualMatched(t, 42, 42, true)
	assertIsEqualMatched(t, 42.0, 42.0, true)
	assertIsEqualMatched(t, "string", "string", true)
	assertIsEqualMatched(t, int16(42), int64(42), true)
	assertIsEqualMatched(t, int64(42), int16(42), true)
	assertIsEqualMatched(t, 42, 42.0, true)
	assertIsEqualMatched(t, struct{ test string }{"string"}, struct{ test string }{"string"}, true)
	assertIsEqualMatched(t, nil, nil, true)
}

func TestIsEqualMismatches(t *testing.T) {
	assertIsEqualMatched(t, 42, 43, false)
	assertIsEqualMatched(t, 42.0, 42.1, false)
	assertIsEqualMatched(t, "string", "strung", false)
	assertIsEqualMatched(t, struct{ test string }{"string"}, struct{ test string }{"strung"}, false)
	assertIsEqualMatched(t, struct{ test string }{"string"}, struct{ test1 string }{"string"}, false)
	assertIsEqualMatched(t, struct{ test string }{"string"}, nil, false)
	assertIsEqualMatched(t, nil, struct{ test string }{"string"}, false)
}
