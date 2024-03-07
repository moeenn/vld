package vld

import (
	"testing"
)

func TestNonEmptyStringValid(t *testing.T) {
	input := "Some non-empty string"
	err := NonEmptyString(input)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.(ValidationErrors).Errors)
		return
	}
}

func TestNonEmptyStringInvalidEmpty(t *testing.T) {
	input := ""
	err := NonEmptyString(input)
	if err == nil {
		t.Error("empty string returned as valid")
		return
	}
}

func TestNonEmptyStringInvalidType(t *testing.T) {
	err := NonEmptyString(10)
	if err == nil {
		t.Error("NonEmptyString: invalid type input successfully validated")
		return
	}
}

func TestMinLengthValid(t *testing.T) {
	err := MinLength(8)("q1w2e3r4t5")
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.(ValidationErrors).Errors)
		return
	}

	err = MinLength(4)("q1w2")
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.(ValidationErrors).Errors)
		return
	}
}

func TestMinLengthInvalidInput(t *testing.T) {
	err := MinLength(8)("q1w2e3")
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestMinLengthInvalidTypeInput(t *testing.T) {
	err := MinLength(8)(true)
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}
}
