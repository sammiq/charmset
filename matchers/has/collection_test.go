package has

import (
	"reflect"
	"testing"

	"github.com/sammiq/charmset/internal"
)

func Test_everyInSlice_match_when_no_elements(t *testing.T) {
	matcher := &internal.MockMatcher{}
	actualValue := reflect.ValueOf([]int{})
	err := everyInSliceMatch(matcher, actualValue)
	if err == nil || err.Error() != "was empty" {
		t.Error("should never mismatch with an error")
	}
	if matcher.CallCount != 0 {
		t.Error("should never call matcher")
	}
}

func Test_everyInSliceMatches(t *testing.T) {
	tests := []struct {
		name   string
		actual interface{}
	}{
		{"match when single element", []int{42}},
		{"match when two elements", []int{42, 24}},
		{"match when multiple elements", []int{42, 24, 84}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := &internal.MockMatcher{Matches: true}
			actualValue := reflect.ValueOf(tt.actual)
			if err := everyInSliceMatch(matcher, actualValue); err != nil {
				t.Error("should never mismatch with an error")
			}
			if matcher.CallCount != actualValue.Len() {
				t.Error("should call matcher for every item")
			}
		})
	}
}

func Test_everyInSliceMismatches(t *testing.T) {
	tests := []struct {
		name   string
		actual interface{}
	}{
		{"mismatch when single element", []int{42}},
		{"mismatch when two elements", []int{42, 24}},
		{"mismatch when multiple elements", []int{42, 24, 84}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := &internal.MockMatcher{Matches: false}
			actualValue := reflect.ValueOf(tt.actual)
			err := everyInSliceMatch(matcher, actualValue)
			if err == nil || err.Error() != "contained an item where did not match" {
				t.Error("should always mismatch with an error")
			}
			if matcher.CallCount != 1 {
				t.Error("should call matcher until fails")
			}
		})
	}
}

func Test_anyInSlice_mismatch_when_no_elements(t *testing.T) {
	matcher := &internal.MockMatcher{}
	actualValue := reflect.ValueOf([]int{})
	err := anyInSliceMatch(matcher, actualValue)
	if err == nil || err.Error() != "was empty" {
		t.Error("should always mismatch with an error")
	}
	if matcher.CallCount != 0 {
		t.Error("should never call matcher")
	}
}

func Test_anyInSliceMatches(t *testing.T) {
	tests := []struct {
		name   string
		actual interface{}
	}{
		{"match when single element", []int{42}},
		{"match when two elements", []int{42, 24}},
		{"match when multiple elements", []int{42, 24, 84}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := &internal.MockMatcher{Matches: true}
			actualValue := reflect.ValueOf(tt.actual)
			if err := anyInSliceMatch(matcher, actualValue); err != nil {
				t.Error("should never mismatch with an error")
			}
			if matcher.CallCount != 1 {
				t.Error("should call matcher until match")
			}
		})
	}
}

func Test_anyInSliceMismatches(t *testing.T) {
	tests := []struct {
		name          string
		actual        interface{}
		expectedError string
	}{
		{"mismatch when single element", []int{42}, "contained an item where did not match"},
		{"mismatch when two elements", []int{42, 24}, "no item matched where [\n          did not match,\n          did not match\n          ]"},
		{"mismatch when multiple elements", []int{42, 24, 84}, "no item matched where [\n          did not match,\n          did not match,\n          did not match\n          ]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := &internal.MockMatcher{Matches: false}
			actualValue := reflect.ValueOf(tt.actual)
			err := anyInSliceMatch(matcher, actualValue)
			if err == nil || err.Error() != tt.expectedError {
				t.Error("should always mismatch with an error")
			}
			if matcher.CallCount != actualValue.Len() {
				t.Error("should call matcher for every item")
			}
		})
	}
}

func Test_matchSliceSequence(t *testing.T) {
	tests := []struct {
		name        string
		expected    []interface{}
		actual      interface{}
		expectMatch bool
	}{
		{"should always match empty sequence", []interface{}{}, []int{42}, true},
		{"should always match single matching sequence", []interface{}{42}, []int{42}, true},
		{"should not match single non-matching sequence", []interface{}{43}, []int{42}, false},
		{"should not match out of order sequence", []interface{}{42, 43}, []int{43, 42}, false},
		{"should always match multiple matching sequence", []interface{}{42, 43, 44}, []int{42, 43, 44}, true},
		{"should always match multiple matching sub-sequence", []interface{}{42, 43}, []int{42, 43, 44}, true},
		{"should always match when restarting correct sub-sequence", []interface{}{42, 43, 44}, []int{42, 43, 42, 43, 44}, true},
		{"should never match when restarting incorrect sub-sequence", []interface{}{42, 43, 44}, []int{42, 43, 42, 43, 42}, false},
		{"should never match when too short actual sequence at start", []interface{}{42, 43, 44}, []int{42, 43}, false},
		{"should never match when too short actual sequence at end", []interface{}{42, 43, 44}, []int{43, 44}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := reflect.ValueOf(tt.actual)
			err := matchSliceSequence(tt.expected, actualValue)
			if tt.expectMatch {
				if err != nil {
					t.Errorf("did not expect error. was: %s", err)
				}
			} else {
				if err == nil {
					t.Error("expected error but none was returned")
				}
			}
		})
	}
}
