package vld

import (
	"regexp"
	"testing"
)

type testCase struct {
	Input   string
	IsValid bool
}

func TestEmailPattern(t *testing.T) {
	testCases := []testCase{
		{Input: "admin@site.com", IsValid: true},
		{Input: "random@site.co.uk", IsValid: true},
		{Input: "some.random-email@site.com", IsValid: true},
		{Input: "random-site.com", IsValid: false},
		{Input: "random@site-org", IsValid: false},
		{Input: "random.string", IsValid: false},
	}

	for _, testCase := range testCases {
		_, err := regexp.MatchString(PATTERN_EMAIL, testCase.Input)
		if testCase.IsValid && err != nil {
			t.Errorf("pattern failed for valid input: %s", testCase.Input)
		}
	}
}

func TestUUIDPattern(t *testing.T) {
	testCases := []testCase{
		{Input: "bb4b84e0-dd43-11ee-abc5-d31f3e7a8b9a", IsValid: true}, // UUIDv1
		{Input: "c0e219b3-0302-409e-a5d8-f297f789c77a", IsValid: true}, // UUIDv4
		{Input: "342eef96-0452-5b99-ac4d-612f79cb102c", IsValid: true}, // UUIDv5
		{Input: "not-a-uuid-string", IsValid: false},
	}

	for _, testCase := range testCases {
		match, err := regexp.MatchString(PATTERN_UUID, testCase.Input)
		isCurrentValid := err == nil || match

		if testCase.IsValid && !isCurrentValid {
			t.Errorf("pattern failed for valid input: %s", testCase.Input)
		}
	}
}
