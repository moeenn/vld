package vld

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strings"
)

// NonEmptyString check if provided input is a non-empty string.
func NonEmptyString(input any) error {
	err := errors.New("Please provide a non-empty string")
	asString, ok := input.(string)
	if !ok || asString == "" {
		return err
	}

	return nil
}

// Length check if the provided input is a string and its length is equal to
// the provided length.
func Length(length int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be %d characters in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) != length {
			return err
		}

		return nil
	}
}

// MinLength check if provided input is a string and its length is more than
// or equal to the provided length.
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
// equal to the provided length.
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

// MinFloat check if provided input is a number and its value is less than or
// equal to the provided limit.
func MinFloat(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be more than or equal to %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum < limit {
			return err
		}

		return nil
	}
}

// MaxFloat check if provided input is a number and its value is more than or
// equal to the provided limit.
func MaxFloat(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be less than or equal to %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum > limit {
			return err
		}

		return nil
	}
}

// MinInt check if provided input is a number and its value is less than or
// equal to the provided limit.
func MinInt(limit int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be more than or equal to %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum < limit {
			return err
		}

		return nil
	}
}

// MaxInt check if provided input is a number and its value is more than or
// equal to the provided limit.
func MaxInt(limit int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be less than or equal to %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum > limit {
			return err
		}

		return nil
	}
}

// GreaterThanFloat check if provided input is a number and its value is greater
// than but not equal to the provided limit.
func GreaterThanFloat(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be greater than %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum <= limit {
			return err
		}

		return nil
	}
}

// LessThanFloat check if provided input is a number and its value is less than
// but not equal to the provided limit.
func LessThanFloat(limit float64) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be less than %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum >= limit {
			return err
		}

		return nil
	}
}

// GreaterThanInt check if provided input is a number and its value is greater than
// but not equal to the provided limit.
func GreaterThanInt(limit int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be greater than %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum <= limit {
			return err
		}

		return nil
	}
}

// LessThanInt check if provided input is a number and its value is less than
// but not equal to the provided limit.
func LessThanInt(limit int) Rule {
	return func(input any) error {
		err := fmt.Errorf("The value must be less than %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum >= limit {
			return err
		}

		return nil
	}
}

// Email check if the provide input is a valid email address.
func Email(input any) error {
	err := errors.New("Please provide a valid email address")
	asString, ok := input.(string)
	if !ok {
		return err
	}

	match, errMatch := regexp.MatchString(PATTERN_EMAIL, asString)
	if errMatch != nil || match == false {
		return err
	}

	return nil
}

// StartsWith check if the provided input is a valid string and starts with the
// provided substring
func StartsWith(prefix string) Rule {
	return func(input any) error {
		err := fmt.Errorf("The input must start with '%s'", prefix)
		asString, ok := input.(string)
		if !ok || !strings.HasPrefix(asString, prefix) {
			return err
		}
		return nil
	}
}

// EndsWith check if the provided input is a valid string and ends with the
// provided substring
func EndsWith(suffix string) Rule {
	return func(input any) error {
		err := fmt.Errorf("The input must end with '%s'", suffix)
		asString, ok := input.(string)
		if !ok || !strings.HasSuffix(asString, suffix) {
			return err
		}
		return nil
	}
}

// DoesntStartWith check if the provided input is a valid string and doesn't
// starts with the provided substring
func DoesntStartWith(prefix string) Rule {
	return func(input any) error {
		err := fmt.Errorf("The input must not start with '%s'", prefix)
		asString, ok := input.(string)
		if !ok || strings.HasPrefix(asString, prefix) {
			return err
		}
		return nil
	}
}

// DoesntEndWith check if the provided input is a valid string and ends with the
// provided substring
func DoesntEndWith(suffix string) Rule {
	return func(input any) error {
		err := fmt.Errorf("The input must end with '%s'", suffix)
		asString, ok := input.(string)
		if !ok || strings.HasSuffix(asString, suffix) {
			return err
		}
		return nil
	}
}

// Same check if the provided input is the same as the required input.
func Same(targetName string, targetValue any) Rule {
	return func(input any) error {
		err := fmt.Errorf("The input must be the same as '%s'", targetName)
		if targetValue != input {
			return err
		}
		return nil
	}
}

// Enum check if the provided input matches any of the listed enumerations values
func Enum(enumValues ...string) Rule {
	return func(input any) error {
		err := fmt.Errorf("The input must match values %s", strings.Join(enumValues, ", "))
		asString, ok := input.(string)
		if !ok || !slices.Contains(enumValues, asString) {
			return err
		}
		return nil
	}
}

// URL check if the provided input is a valid string and a valid URL.
func URL(input any) error {
	err := errors.New("The input must be a valid URL")
	asString, ok := input.(string)
	if !ok {
		return err
	}

	_, err = url.ParseRequestURI(asString)
	if err != nil {
		return err
	}
	return nil
}

// Regexp check if the provided input is a valid string and matches the required
// regular expression.
func Regexp(pattern string) Rule {
	return func(input any) error {
		err := errors.New("The input doesn't match the required pattern")
		asString, ok := input.(string)
		if !ok {
			return err
		}

		match, errMatch := regexp.MatchString(pattern, asString)
		if errMatch != nil || !match {
			return err
		}
		return nil
	}
}

// UUID check if the provided input is a valid string and a valid UUID.
func UUID(input any) error {
	err := errors.New("The input must be a valid UUID string")
	asString, ok := input.(string)
	if !ok {
		return err
	}

	match, errMatch := regexp.MatchString(PATTERN_UUID, asString)
	if errMatch != nil || !match {
		return err
	}
	return nil
}

// Password check if the provided input is a valid string and a reasonably strong
// password.
func Password(input any) error {
	err := errors.New("Please provide a stronger password")
	asString, ok := input.(string)
	if !ok {
		return err
	}

	match, errMatch := regexp.MatchString(PATTERN_PASSWORD_STRENGTH, asString)

	// NOTE: regexp pattern being used matches invalid passwords instead of
	// strong passwords. If match is true, it means password was weak.
	if errMatch != nil || match {
		return err
	}
	return nil
}

// JSON check if the provided code is a valid string and a valid json.
func JSON(input any) error {
	err := errors.New("The input must be a valid JSON string")
	asString, ok := input.(string)
	if !ok {
		return err
	}

	var target any
	errUnmarshal := json.Unmarshal([]byte(asString), &target)
	if errUnmarshal != nil {
		return err
	}
	return nil
}
