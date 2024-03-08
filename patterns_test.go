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
			t.Errorf(errValidFailed, testCase.Input)
			return
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
			t.Errorf(errValidFailed, testCase.Input)
			return
		}
	}
}

func TestPasswordStrengthPattern(t *testing.T) {
	testCases := []testCase{
		{Input: "123_Apple", IsValid: true},
		{Input: "Ax2@11", IsValid: false},                // less than 8 characters in length
		{Input: "vmdk2@vkvqew", IsValid: false},          // upper-case letter missing
		{Input: "543*&SKCM92S0C//", IsValid: false},      // lower-case letter missing
		{Input: "@!Ackanslcksan_scm$", IsValid: false},   // number missing
		{Input: "123cckasnlcKACNslls13", IsValid: false}, // special-character missing
		{Input: "password", IsValid: false},              // multiple-rule violations
		{Input: "abc123", IsValid: false},                // multiple-rule violations
	}

	for _, testCase := range testCases {
		match, err := regexp.MatchString(PATTERN_PASSWORD_STRENGTH, testCase.Input)

		// NOTE: regexp pattern being used matches invalid passwords instead of
		// strong passwords. If match is true, it means password was weak.
		isCurrentValid := err == nil || !match

		if testCase.IsValid && !isCurrentValid {
			t.Errorf(errValidFailed, testCase.Input)
			return
		}
	}
}
