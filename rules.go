package vld

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"time"
)

// NonEmptyString check if provided input is a non-empty string.
func NonEmptyString(input any) (any, error) {
	err := errors.New("Please provide a non-empty string")
	asString, ok := input.(string)
	if !ok || asString == "" {
		return nil, err
	}

	return asString, nil
}

// Length check if the provided input is a string and its length is equal to
// the provided length.
func Length(length int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be %d characters in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) != length {
			return nil, err
		}

		return asString, nil
	}
}

// MinLength check if provided input is a string and its length is more than
// or equal to the provided length.
func MinLength(length int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be %d characters or more in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) < length {
			return nil, err
		}

		return asString, nil
	}
}

// MaxLength check if provided input is a string and its length is less than or
// equal to the provided length.
func MaxLength(length int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be %d characters or less in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) > length {
			return nil, err
		}

		return asString, nil
	}
}

// MinFloat check if provided input is a number and its value is less than or
// equal to the provided limit.
func MinFloat(limit float64) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be more than or equal to %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum < limit {
			return nil, err
		}

		return asNum, nil
	}
}

// MaxFloat check if provided input is a number and its value is more than or
// equal to the provided limit.
func MaxFloat(limit float64) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be less than or equal to %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum > limit {
			return nil, err
		}

		return asNum, nil
	}
}

// MinInt check if provided input is a number and its value is less than or
// equal to the provided limit.
func MinInt(limit int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be more than or equal to %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum < limit {
			return nil, err
		}

		return asNum, nil
	}
}

// MaxInt check if provided input is a number and its value is more than or
// equal to the provided limit.
func MaxInt(limit int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be less than or equal to %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum > limit {
			return nil, err
		}

		return asNum, nil
	}
}

// GreaterThanFloat check if provided input is a number and its value is greater
// than but not equal to the provided limit.
func GreaterThanFloat(limit float64) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be greater than %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum <= limit {
			return nil, err
		}

		return asNum, nil
	}
}

// LessThanFloat check if provided input is a number and its value is less than
// but not equal to the provided limit.
func LessThanFloat(limit float64) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be less than %f", limit)
		asNum, ok := input.(float64)
		if !ok || asNum >= limit {
			return nil, err
		}

		return asNum, nil
	}
}

// GreaterThanInt check if provided input is a number and its value is greater than
// but not equal to the provided limit.
func GreaterThanInt(limit int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be greater than %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum <= limit {
			return nil, err
		}

		return asNum, nil
	}
}

// LessThanInt check if provided input is a number and its value is less than
// but not equal to the provided limit.
func LessThanInt(limit int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The value must be less than %d", limit)
		asNum, ok := input.(int)
		if !ok || asNum >= limit {
			return nil, err
		}

		return asNum, nil
	}
}

// Email check if the provide input is a valid email address.
func Email(input any) (any, error) {
	err := errors.New("Please provide a valid email address")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	match, errMatch := regexp.MatchString(PATTERN_EMAIL, asString)
	if errMatch != nil || match == false {
		return nil, err
	}

	return asString, nil
}

// StartsWith check if the provided input is a valid string and starts with the
// provided substring
func StartsWith(prefix string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The input must start with '%s'", prefix)
		asString, ok := input.(string)
		if !ok || !strings.HasPrefix(asString, prefix) {
			return nil, err
		}
		return asString, nil
	}
}

// EndsWith check if the provided input is a valid string and ends with the
// provided substring
func EndsWith(suffix string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The input must end with '%s'", suffix)
		asString, ok := input.(string)
		if !ok || !strings.HasSuffix(asString, suffix) {
			return nil, err
		}
		return asString, nil
	}
}

// DoesntStartWith check if the provided input is a valid string and doesn't
// starts with the provided substring
func DoesntStartWith(prefix string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The input must not start with '%s'", prefix)
		asString, ok := input.(string)
		if !ok || strings.HasPrefix(asString, prefix) {
			return nil, err
		}
		return asString, nil
	}
}

// DoesntEndWith check if the provided input is a valid string and ends with the
// provided substring
func DoesntEndWith(suffix string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The input must end with '%s'", suffix)
		asString, ok := input.(string)
		if !ok || strings.HasSuffix(asString, suffix) {
			return nil, err
		}
		return asString, nil
	}
}

// Same check if the provided input is the same as the required input.
func Same(targetName string, targetValue any) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The input must be the same as '%s'", targetName)
		if targetValue != input {
			return nil, err
		}
		return targetValue, nil
	}
}

// Enum check if the provided input matches any of the listed enumerations values
func Enum(enumValues ...string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("The input must match values %s", strings.Join(enumValues, ", "))
		asString, ok := input.(string)
		if !ok || !slices.Contains(enumValues, asString) {
			return nil, err
		}
		return asString, nil
	}
}

// URL check if the provided input is a valid string and a valid URL.
func URL(input any) (any, error) {
	err := errors.New("The input must be a valid URL")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	_, err = url.ParseRequestURI(asString)
	if err != nil {
		return nil, err
	}
	return asString, nil
}

// Regexp check if the provided input is a valid string and matches the required
// regular expression.
func Regexp(pattern string) Rule {
	return func(input any) (any, error) {
		err := errors.New("The input doesn't match the required pattern")
		asString, ok := input.(string)
		if !ok {
			return nil, err
		}

		match, errMatch := regexp.MatchString(pattern, asString)
		if errMatch != nil || !match {
			return nil, err
		}
		return asString, nil
	}
}

// UUID check if the provided input is a valid string and a valid UUID.
func UUID(input any) (any, error) {
	err := errors.New("The input must be a valid UUID string")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	match, errMatch := regexp.MatchString(PATTERN_UUID, asString)
	if errMatch != nil || !match {
		return nil, err
	}
	return asString, nil
}

// Password check if the provided input is a valid string and a reasonably strong
// password.
func Password(input any) (any, error) {
	err := errors.New("Please provide a stronger password")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	match, errMatch := regexp.MatchString(PATTERN_PASSWORD_STRENGTH, asString)

	// NOTE: regexp pattern being used matches invalid passwords instead of
	// strong passwords. If match is true, it means password was weak.
	if errMatch != nil || match {
		return nil, err
	}
	return asString, nil
}

// JSON check if the provided code is a valid string and a valid json.
func JSON(input any) (any, error) {
	err := errors.New("The input must be a valid JSON string")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	var target any
	errUnmarshal := json.Unmarshal([]byte(asString), &target)
	if errUnmarshal != nil {
		return nil, err
	}
	return asString, nil
}

// ISODate check if the provided input is a valid string and a valid ISO
// timestamp according to RFC3339: [Link](https://pkg.go.dev/time#pkg-constants).
func ISODate(input any) (any, error) {
	err := errors.New("Please provide a valid date time string")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	asDate, parseErr := time.Parse(time.RFC3339, asString)
	if parseErr != nil {
		return nil, err
	}
	return asDate, nil
}

// Date check if the provided input is a valid string and a valid date-only
// string. Date string must be in format e.g. 2023-10-05
// [Link](https://pkg.go.dev/time#pkg-constants).
func Date(input any) (any, error) {
	err := errors.New("Please provide a valid date")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	asDate, parseErr := time.Parse(time.DateOnly, asString)
	if parseErr != nil {
		return nil, err
	}
	return asDate, nil
}

// Time check if the provided input is a valid string and a valid time-only
// string. Time string must be in 24-hours format: e.g. 10:20:00
// [Link](https://pkg.go.dev/time#pkg-constants).
func Time(input any) (any, error) {
	err := errors.New("Please provide a valid time")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	asTime, parseErr := time.Parse(time.TimeOnly, asString)
	if parseErr != nil {
		return nil, err
	}
	return asTime, nil
}
