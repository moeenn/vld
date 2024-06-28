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
	asString, ok := input.(string)
	if !ok || asString == "" {
		return nil, errors.New("please provide a non-empty string")
	}

	return asString, nil
}

// Length check if the provided input is a string and its length is equal to
// the provided length.
func Length(length int) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("the value must be %d characters in length", length)
		asString, ok := input.(string)
		if !ok || len(asString) != length {
			return nil, err
		}

		return asString, nil
	}
}

// Min if the provided number is an int / float(64), check input is greater than
// or equal to the target. If the provided input is a string, check its length
// is more than or equal to the target.
func Min(target any) Rule {
	return func(input any) (any, error) {
		targetAsInt, okTargetIntCast := target.(int)
		targetAsFloat, okTargetFloatCast := target.(float64)

		if !okTargetIntCast && !okTargetFloatCast {
			return nil, errors.New("invalid target type provided")
		}

		asInt, okIntCast := input.(int)
		if okIntCast {
			if okTargetIntCast {
				if asInt < targetAsInt {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asInt, nil
			}

			if okTargetFloatCast {
				if float64(asInt) < targetAsFloat {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asInt, nil
			}
		}

		asFloat, okFloatCast := input.(float64)
		if okFloatCast {
			if okTargetIntCast {
				if asFloat < float64(targetAsInt) {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asFloat, nil
			}

			if okTargetFloatCast {
				if asFloat < targetAsFloat {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asFloat, nil
			}
		}

		asString, okStringCast := input.(string)
		if okStringCast {
			if okTargetIntCast {
				if len(asString) < targetAsInt {
					return nil, fmt.Errorf("length must be more than %d characters", target)
				}
				return asString, nil
			}

			if okTargetFloatCast {
				return nil, errors.New("string length cannot be a floating point number")
			}
		}

		if !okIntCast && !okFloatCast && !okStringCast {
			return nil, errors.New("invalid data type provided")
		}

		return nil, nil
	}
}

// Max if the provided number is an int / float(64), check input is less than
// or equal to the target. If the provided input is a string, check its length
// is less than or equal to the target.
func Max(target any) Rule {
	return func(input any) (any, error) {
		targetAsInt, okTargetIntCast := target.(int)
		targetAsFloat, okTargetFloatCast := target.(float64)

		if !okTargetIntCast && !okTargetFloatCast {
			return nil, errors.New("invalid target type provided")
		}

		asInt, okIntCast := input.(int)
		if okIntCast {
			if okTargetIntCast {
				if asInt > targetAsInt {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asInt, nil
			}

			if okTargetFloatCast {
				if float64(asInt) > targetAsFloat {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asInt, nil
			}
		}

		asFloat, okFloatCast := input.(float64)
		if okFloatCast {
			if okTargetIntCast {
				if asFloat > float64(targetAsInt) {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asFloat, nil
			}

			if okTargetFloatCast {
				if asFloat > targetAsFloat {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asFloat, nil
			}
		}

		asString, okStringCast := input.(string)
		if okStringCast {
			if okTargetIntCast {
				if len(asString) > targetAsInt {
					return nil, fmt.Errorf("length must be less than %d characters", target)
				}
				return asString, nil
			}

			if okTargetFloatCast {
				return nil, errors.New("string length cannot be a floating point number")
			}
		}

		if !okIntCast && !okFloatCast && !okStringCast {
			return nil, errors.New("invalid data type provided")
		}

		return nil, nil
	}
}

// GreaterThan if the provided number is an int / float(64), check input is more
// than (but not equal) to the target. If the provided input is a string, check its
// length is more than (but not equal) to the target.
func GreaterThan(target any) Rule {
	return func(input any) (any, error) {
		targetAsInt, okTargetIntCast := target.(int)
		targetAsFloat, okTargetFloatCast := target.(float64)

		if !okTargetIntCast && !okTargetFloatCast {
			return nil, errors.New("invalid target type provided")
		}

		asInt, okIntCast := input.(int)
		if okIntCast {
			if okTargetIntCast {
				if asInt <= targetAsInt {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asInt, nil
			}

			if okTargetFloatCast {
				if float64(asInt) <= targetAsFloat {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asInt, nil
			}
		}

		asFloat, okFloatCast := input.(float64)
		if okFloatCast {
			if okTargetIntCast {
				if asFloat <= float64(targetAsInt) {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asFloat, nil
			}

			if okTargetFloatCast {
				if asFloat <= targetAsFloat {
					return nil, fmt.Errorf("number must be greater than %d", target)
				}
				return asFloat, nil
			}
		}

		asString, okStringCast := input.(string)
		if okStringCast {
			if okTargetIntCast {
				if len(asString) <= targetAsInt {
					return nil, fmt.Errorf("length must be more than %d characters", target)
				}
				return asString, nil
			}

			if okTargetFloatCast {
				return nil, errors.New("string length cannot be a floating point number")
			}
		}

		if !okIntCast && !okFloatCast && !okStringCast {
			return nil, errors.New("invalid data type provided")
		}

		return nil, nil
	}
}

// LessThan if the provided number is an int / float(64), check input is less
// than (but not equal) to the target. If the provided input is a string, check its
// length is less than (but not equal) to the target.
func LessThan(target any) Rule {
	return func(input any) (any, error) {
		targetAsInt, okTargetIntCast := target.(int)
		targetAsFloat, okTargetFloatCast := target.(float64)

		if !okTargetIntCast && !okTargetFloatCast {
			return nil, errors.New("invalid target type provided")
		}

		asInt, okIntCast := input.(int)
		if okIntCast {
			if okTargetIntCast {
				if asInt >= targetAsInt {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asInt, nil
			}

			if okTargetFloatCast {
				if float64(asInt) >= targetAsFloat {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asInt, nil
			}
		}

		asFloat, okFloatCast := input.(float64)
		if okFloatCast {
			if okTargetIntCast {
				if asFloat >= float64(targetAsInt) {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asFloat, nil
			}

			if okTargetFloatCast {
				if asFloat >= targetAsFloat {
					return nil, fmt.Errorf("number must be less than %d", target)
				}
				return asFloat, nil
			}
		}

		asString, okStringCast := input.(string)
		if okStringCast {
			if okTargetIntCast {
				if len(asString) >= targetAsInt {
					return nil, fmt.Errorf("length must be less than %d characters", target)
				}
				return asString, nil
			}

			if okTargetFloatCast {
				return nil, errors.New("string length cannot be a floating point number")
			}
		}

		if !okIntCast && !okFloatCast && !okStringCast {
			return nil, errors.New("invalid data type provided")
		}

		return nil, nil
	}
}

// Email check if the provide input is a valid email address.
func Email(input any) (any, error) {
	err := errors.New("please provide a valid email address")
	asString, ok := input.(string)
	if !ok {
		return nil, err
	}

	match, errMatch := regexp.MatchString(PATTERN_EMAIL, asString)
	if errMatch != nil || !match {
		return nil, err
	}

	return asString, nil
}

// StartsWith check if the provided input is a valid string and starts with the
// provided substring
func StartsWith(prefix string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("the input must start with '%s'", prefix)
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
		err := fmt.Errorf("the input must end with '%s'", suffix)
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
		err := fmt.Errorf("the input must not start with '%s'", prefix)
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
		err := fmt.Errorf("the input must end with '%s'", suffix)
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
		err := fmt.Errorf("the input must be the same as '%s'", targetName)
		if targetValue != input {
			return nil, err
		}
		return targetValue, nil
	}
}

// Enum check if the provided input matches any of the listed enumerations values
func Enum(enumValues ...string) Rule {
	return func(input any) (any, error) {
		err := fmt.Errorf("the input must match values %s", strings.Join(enumValues, ", "))
		asString, ok := input.(string)
		if !ok || !slices.Contains(enumValues, asString) {
			return nil, err
		}
		return asString, nil
	}
}

// URL check if the provided input is a valid string and a valid URL.
func URL(input any) (any, error) {
	err := errors.New("the input must be a valid URL")
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
		err := errors.New("the input doesn't match the required pattern")
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
	err := errors.New("the input must be a valid UUID string")
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
	err := errors.New("please provide a stronger password")
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
	err := errors.New("the input must be a valid JSON string")
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

// DateTime check if the provided input is a valid string and a valid ISO
// timestamp according to RFC3339: [Link](https://pkg.go.dev/time#pkg-constants).
func DateTime(input any) (any, error) {
	err := errors.New("please provide a valid date")
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

// Date check if the provided input is a valid date-only string. Date string
// must be in format e.g. 2023-10-05
// [Link](https://pkg.go.dev/time#pkg-constants).
func Date(input any) (any, error) {
	err := errors.New("please provide a valid date")
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
	err := errors.New("please provide a valid time")
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

// DateEqual check if the provided date is a date equal to the target date.
func DateEqual(target time.Time) Rule {
	return func(input any) (any, error) {
		inputAsTime, ok := input.(time.Time)
		if !ok {
			return nil, errors.New("please provide a valid date")
		}

		delta := inputAsTime.Sub(target)
		if delta != 0 {
			return nil, errors.New("the provided date must be " + target.String())
		}

		return inputAsTime, nil
	}
}

// DateBefore check if the provided input is a date before (but not equal) to
// the target date.
func DateBefore(target time.Time, inclusive bool) Rule {
	return func(input any) (any, error) {
		inputAsTime, ok := input.(time.Time)
		if !ok {
			return nil, errors.New("please provide a valid date")
		}

		delta := inputAsTime.Sub(target)
		if !inclusive && delta >= 0 {
			return nil, errors.New("the provided date must be before " + target.String())
		}

		if inclusive && delta > 0 {
			return nil, errors.New("the provided date must be before or equal to " + target.String())
		}

		return inputAsTime, nil
	}
}

// DateAfter check if the provided input is a date after the target date. If
// inclusive is set to true, target date will be included.
func DateAfter(target time.Time, inclusive bool) Rule {
	return func(input any) (any, error) {
		inputAsTime, ok := input.(time.Time)
		if !ok {
			return nil, errors.New("please provide a valid date")
		}

		delta := inputAsTime.Sub(target)
		if !inclusive && delta <= 0 {
			return nil, errors.New("the provided date must be after " + target.String())
		}

		if inclusive && delta < 0 {
			return nil, errors.New("the provided date must be after or equal to " + target.String())
		}

		return inputAsTime, nil
	}
}

// Latitude check if the provided input a valid map latitude value.
func Latitude(input any) (any, error) {
	err := errors.New("please provide a valid latitude value")

	inputAsFloat, ok := input.(float32)
	if !ok {
		return nil, err
	}

	if inputAsFloat < -90.0 || inputAsFloat > 90.0 {
		return nil, err
	}

	return inputAsFloat, nil
}

// Longitude check if the provided input a valid map longitude value.
func Longitude(input any) (any, error) {
	err := errors.New("please provide a valid longitude value")

	inputAsFloat, ok := input.(float32)
	if !ok {
		return nil, err
	}

	if inputAsFloat < -180.0 || inputAsFloat > 180.0 {
		return nil, err
	}

	return inputAsFloat, nil
}
