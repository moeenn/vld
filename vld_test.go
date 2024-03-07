package vld

import (
	"testing"
)

func TestLoginFormExample(t *testing.T) {
	validations := []Validation{
		{
			Tag:   "email",
			Data:  "admin@site.com",
			Rules: []Rule{NonEmptyString, Email},
		},
		{
			Tag:   "password",
			Data:  "q1w2e3r4",
			Rules: []Rule{NonEmptyString, MinLength(8)},
		},
	}

	err := Validate(validations)
	if err != nil {
		t.Error("valid data returned as invalid")
		return
	}
	// asJSON, err := json.Marshal(validationErrors.(ValidationErrors).Errors)
}
