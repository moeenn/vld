package vld

import (
	"regexp"
	"testing"
)

type testCase struct {
	Input   string
	IsValid bool
}

func TestEmailPatternTests(t *testing.T) {
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
