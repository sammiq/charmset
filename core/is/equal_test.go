package is

import (
	"testing"
)

func assertEqualMatched(t *testing.T, expected interface{}, actual interface{}, expectMatch bool) {
	err := equal(expected, actual)
	if (err == nil) != expectMatch {
		t.Errorf("Expected err to be returned when match is %v, expected is %v and actual is %v", expectMatch, expected, actual)
	}
}

func Test_equalMatches(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"should be equal with two integers the same", 42, 42},
		{"should be equal with two doubles the same", 42.0, 42.0},
		{"should be equal with two strings the same", "string", "string"},
		{"should be equal with two integers the same number but expected less bytes", int16(42), int64(42)},
		{"should be equal with two integers the same number but expected more bytes", int64(42), int16(42)},
		{"should be equal with expected as integer but actual as double", 42, 42.0},
		{"should be equal with expected as double but actual as integer", 42.0, 42},
		{"should be equal with two structs with the same contents", struct{ test string }{"string"}, struct{ test string }{"string"}},
		{"should be equal with two nil values", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqualMatched(t, tt.expected, tt.actual, true)
		})
	}
}

func Test_equalMismatches(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
	}{
		{"should not be equal if two integers different values", 42, 43},
		{"should not be equal if two doubles different values", 42.0, 42.1},
		{"should not be equal if two strings different values", "string", "strung"},
		{"should not be equal if two integers different values when expected is less bytes", int16(42), int64(43)},
		{"should not be equal if two integers different values when expected is more bytes", int64(42), int16(43)},
		{"should not be equal with expected as integer but actual as double with different values", 42, 42.1},
		{"should not be equal with expected as double but actual as integer with different values", 42.1, 42},
		{"should not be equal with two structs with the same field names but different contents", struct{ test string }{"string"}, struct{ test string }{"strung"}},
		{"should not be equal with two structs with the different field names but same contents", struct{ test string }{"string"}, struct{ test1 string }{"string"}},
		{"should not be equal expected integer but actual nil", 42, nil},
		{"should not be equal expected double but actual nil", 42.0, nil},
		{"should not be equal expected string but actual nil", "string", nil},
		{"should not be equal expected struct but actual nil", struct{ test string }{"string"}, nil},
		{"should not be equal expected nil but actual integer", nil, 42},
		{"should not be equal expected nil but actual double", nil, 42.0},
		{"should not be equal expected nil but actual string", nil, "string"},
		{"should not be equal expected nil but actual struct", nil, struct{ test string }{"string"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertEqualMatched(t, tt.expected, tt.actual, false)
		})
	}
}
