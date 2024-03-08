package vld

import (
	"testing"
)

func TestLoginFormExample(t *testing.T) {
	form := struct {
		Email    string
		Password string
	}{
		Email:    "admin@site.com",
		Password: "q1w2e3r4",
	}

	validations := []Validation{
		{
			Tag:   "email",
			Data:  form.Email,
			Rules: []Rule{NonEmptyString, Email},
		},
		{
			Tag:   "password",
			Data:  form.Password,
			Rules: []Rule{NonEmptyString, MinLength(8)},
		},
	}

	err := Validate(validations)
	if err != nil {
		t.Error("valid data returned as invalid")
		return
	}
}

func TestResetPasswordExample(t *testing.T) {
	form := struct {
		Password        string
		ConfirmPassword string
	}{
		Password:        "q1w2e3r4",
		ConfirmPassword: "q1w2e3r4",
	}

	validations := []Validation{
		{
			Tag:   "password",
			Data:  form.Password,
			Rules: []Rule{MinLength(8)},
		},
		{
			Tag:   "confirm_password",
			Data:  form.ConfirmPassword,
			Rules: []Rule{Same("Password", form.Password)},
		},
	}

	err := Validate(validations)
	if err != nil {
		t.Error("valid data returned as invalid")
		return
	}
}
