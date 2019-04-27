package core

import (
	"reflect"
	"testing"

	"github.com/sammiq/matchers/internal"
)

func Test_everyInSlice_match_when_no_elements(t *testing.T) {
	matcher := &internal.MockMatcher{}
	actualValue := reflect.ValueOf([]int{})
	if err := everyInSliceMatch(matcher, actualValue); err != nil {
		t.Error("should never fail with an error")
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
			matcher := &internal.MockMatcher{Fail: false}
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
			matcher := &internal.MockMatcher{Fail: true}
			actualValue := reflect.ValueOf(tt.actual)
			err := everyInSliceMatch(matcher, actualValue)
			if err == nil || err.Error() != "contained an item where did fail" {
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
			matcher := &internal.MockMatcher{Fail: false}
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
		{"mismatch when single element", []int{42}, "contained an item where did fail"},
		{"mismatch when two elements", []int{42, 24}, "no item matched [\n          did fail,\n          did fail\n          ]"},
		{"mismatch when multiple elements", []int{42, 24, 84}, "no item matched [\n          did fail,\n          did fail,\n          did fail\n          ]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := &internal.MockMatcher{Fail: true}
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
