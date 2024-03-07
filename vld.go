package vld

import (
	"errors"
	"fmt"
	"strings"
)

type Rule func(any) error

type Validation struct {
	Tag   string
	Data  any
	Rules []Rule
}

type ValidationErrors struct {
	Errors map[string]string `json:"errors"`
}

func (v ValidationErrors) Error() string {
	return "Validation of provided data failed"
}

func Validate(validations []Validation) error {
	errors := ValidationErrors{
		Errors: make(map[string]string),
	}

	for _, validation := range validations {
		if _, exists := errors.Errors[validation.Tag]; exists {
			continue
		}

		for _, rule := range validation.Rules {
			err := rule(validation.Data)
			if err != nil {
				errors.Errors[validation.Tag] = err.Error()
			}
		}
	}

	if len(errors.Errors) != 0 {
		return errors
	}

	return nil
}

/**
 * rule functions implementation
 *
 *
 */

// NonEmptyString check if provided input is a non-empty string
func NonEmptyString(input any) error {
	err := errors.New("Please provide a non-empty string")
	asString, ok := input.(string)
	if !ok || asString == "" {
		return err
	}

	return nil
}

// Length check if the provided input is a string and its length is equal to
// the provided length
func Length(length int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be %d characters in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) == length {
			return err
		}

		return nil
	}
}

// MinLength check if provided input is a string and its length is more than
// or equal to the provided length
func MinLength(length int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be %d characters or more in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) < length {
			return err
		}

		return nil
	}
}

// MaxLength check if provided input is a string and its length is less than or
// equal to the provided length
func MaxLength(length int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be %d characters or less in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) > length {
			return err
		}

		return nil
	}
}

// Min check if provided input is a number and its value is less than or equal
// to the provided limit
func Min(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be less than or equal to %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum > limit {
			return err
		}

		return nil
	}
}

// Max check if provided input is a number and its value is more than or equal
// to the provided limit
func Max(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be more than or equal to %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum < limit {
			return err
		}

		return nil
	}
}

// GreaterThan check if provided input is a number and its value is greater than
// but not equal to the provided limit
func GreaterThan(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be greater than %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum <= limit {
			return err
		}

		return nil
	}
}

// LessThan check if provided input is a number and its value is less than
// but not equal to the provided limit
func LessThan(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be less than %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum >= limit {
			return err
		}

		return nil
	}
}

// Email check if the provide input is a valid email address
func Email(input any) error {
	err := errors.New("Please provide a valid email address")
	asString, ok := input.(string)
	if !ok || !strings.Contains(asString, "@") {
		return err
	}

	return nil
}

/*
reference implementation: implement as test

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (f LoginForm) Validate() error {
	validations := []Validation{
		{
			Tag:   "email",
			Data:  f.Email,
			Rules: []Rule{NonEmptyString, Email},
		},
		{
			Tag:   "password",
			Data:  f.Password,
			Rules: []Rule{NonEmptyString, MinLength(8)},
		},
	}

	return Validate(validations)
}

func main() {
	loginForm := LoginForm{
		Email:    "admin-site.com",
		Password: "q1w",
	}

	validationErrors := loginForm.Validate()

	asJSON, err := json.Marshal(validationErrors.(ValidationErrors).Errors)
	if err != nil {
		fmt.Println("failed to serialize errors as json")
		return
	}

	fmt.Printf("%s\n", asJSON)
}
*/
