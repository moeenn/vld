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
	issue := &Issue{
		Code:    CODE_NON_EMPTY_STRING,
		Message: "Please provide a non-empty string",
	}

	asString, ok := input.(string)
	if !ok || asString == "" {
		return nil, issue
	}

	return asString, nil
}

// Length check if the provided input is a string and its length is equal to
// the provided length.
func Length(length int) Rule {
	return func(input any) (any, error) {
		issue := &Issue{
			Code:    CODE_LENGTH,
			Message: fmt.Sprintf("The value must be %d characters in length", length),
			Value:   length,
		}

		asString, ok := input.(string)
		if !ok || len(asString) != length {
			return nil, issue
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

		switch t := input.(type) {
		case int:
			{
				if okTargetIntCast {
					if t < targetAsInt {
						return nil, Issue{
							Code:    CODE_MIN,
							Message: fmt.Sprintf("The number must be greater than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					if float64(t) < targetAsFloat {
						return nil, Issue{
							Code:    CODE_MIN,
							Message: fmt.Sprintf("The number must be greater than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}
			}

		case float64:
			{
				if okTargetIntCast {
					if t < float64(targetAsInt) {
						return nil, Issue{
							Code:    CODE_MIN,
							Message: fmt.Sprintf("The number must be greater than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					if t < targetAsFloat {
						return nil, Issue{
							Code:    CODE_MIN,
							Message: fmt.Sprintf("The number must be greater than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}
			}

		case string:
			{
				if okTargetIntCast {
					if len(t) < targetAsInt {
						return nil, Issue{
							Code:    CODE_MIN,
							Message: fmt.Sprintf("The length must be more than %d characters", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					return nil, errors.New("string length cannot be a floating point number")
				}
			}

		default:
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

		switch t := input.(type) {
		case int:
			{
				if okTargetIntCast {
					if t > targetAsInt {
						return nil, Issue{
							Code:    CODE_MAX,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					if float64(t) > targetAsFloat {
						return nil, Issue{
							Code:    CODE_MAX,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}
			}

		case float64:
			{
				if okTargetIntCast {
					if t > float64(targetAsInt) {
						return nil, Issue{
							Code:    CODE_MAX,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					if t > targetAsFloat {
						return nil, Issue{
							Code:    CODE_MAX,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}
			}

		case string:
			{
				if okTargetIntCast {
					if len(t) > targetAsInt {
						return nil, Issue{
							Code:    CODE_MAX,
							Message: fmt.Sprintf("The length must be less than %d characters", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					return nil, errors.New("string length cannot be a floating point number")
				}
			}

		default:
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

		switch t := input.(type) {
		case int:
			{
				if okTargetIntCast {
					if t <= targetAsInt {
						return nil, fmt.Errorf("number must be greater than %d", target)
					}
					return t, nil
				}

				if okTargetFloatCast {
					if float64(t) <= targetAsFloat {
						return nil, fmt.Errorf("number must be greater than %d", target)
					}
					return t, nil
				}
			}

		case float64:
			{
				if okTargetIntCast {
					if t <= float64(targetAsInt) {
						return nil, fmt.Errorf("number must be greater than %d", target)
					}
					return t, nil
				}

				if okTargetFloatCast {
					if t <= targetAsFloat {
						return nil, fmt.Errorf("number must be greater than %d", target)
					}
					return t, nil
				}
			}

		case string:
			{
				if okTargetIntCast {
					if len(t) <= targetAsInt {
						return nil, fmt.Errorf("length must be more than %d characters", target)
					}
					return t, nil
				}

				if okTargetFloatCast {
					return nil, errors.New("string length cannot be a floating point number")
				}
			}

		default:
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

		switch t := input.(type) {
		case int:
			{
				if okTargetIntCast {
					if t >= targetAsInt {
						return nil, Issue{
							Code:    CODE_LESS_THAN,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					if float64(t) >= targetAsFloat {
						return nil, Issue{
							Code:    CODE_LESS_THAN,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}
			}

		case float64:
			{
				if okTargetIntCast {
					if t >= float64(targetAsInt) {
						return nil, Issue{
							Code:    CODE_LESS_THAN,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					if t >= targetAsFloat {
						return nil, Issue{
							Code:    CODE_LESS_THAN,
							Message: fmt.Sprintf("The number must be less than %d", target),
							Value:   target,
						}
					}
					return t, nil
				}
			}

		case string:
			{
				if okTargetIntCast {
					if len(t) >= targetAsInt {
						return nil, Issue{
							Code:    CODE_LESS_THAN,
							Message: fmt.Sprintf("The length must be less than %d characters", target),
							Value:   target,
						}
					}
					return t, nil
				}

				if okTargetFloatCast {
					return nil, errors.New("string length cannot be a floating point number")
				}
			}

		default:
			return nil, errors.New("invalid data type provided")
		}

		return nil, nil
	}
}

// Email check if the provide input is a valid email address.
func Email(input any) (any, error) {
	issue := Issue{
		Code:    CODE_EMAIL,
		Message: "Please provide a valid email address",
	}

	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	match, errMatch := regexp.MatchString(PATTERN_EMAIL, asString)
	if errMatch != nil || !match {
		return nil, issue
	}

	return asString, nil
}

// HasPrefix check if the provided input is a valid string and starts with the
// provided substring.
func HasPrefix(prefix string) Rule {
	return func(input any) (any, error) {
		issue := Issue{
			Code:    CODE_HAS_PREFIX,
			Message: fmt.Sprintf("The input must start with '%s'", prefix),
			Value:   prefix,
		}

		asString, ok := input.(string)
		if !ok || !strings.HasPrefix(asString, prefix) {
			return nil, issue
		}
		return asString, nil
	}
}

// HasSuffix check if the provided input is a valid string and ends with the
// provided substring.
func HasSuffix(suffix string) Rule {
	return func(input any) (any, error) {
		issue := Issue{
			Code:    CODE_HAS_SUFFIX,
			Message: fmt.Sprintf("The input must end with '%s'", suffix),
			Value:   suffix,
		}

		asString, ok := input.(string)
		if !ok || !strings.HasSuffix(asString, suffix) {
			return nil, issue
		}
		return asString, nil
	}
}

// NotHasPrefix check if the provided input is a valid string and doesn't
// starts with the provided substring.
func NotHasPrefix(prefix string) Rule {
	return func(input any) (any, error) {
		issue := Issue{
			Code:    CODE_NOT_HAS_PREFIX,
			Message: fmt.Sprintf("The input must not start with '%s'", prefix),
			Value:   prefix,
		}

		asString, ok := input.(string)
		if !ok || strings.HasPrefix(asString, prefix) {
			return nil, issue
		}
		return asString, nil
	}
}

// NotHasSuffix check if the provided input is a valid string and ends with the
// provided substring.
func NotHasSuffix(suffix string) Rule {
	return func(input any) (any, error) {
		issue := Issue{
			Code:    CODE_NOT_HAS_SUFFIX,
			Message: fmt.Sprintf("The input must end with '%s'", suffix),
			Value:   suffix,
		}
		asString, ok := input.(string)
		if !ok || strings.HasSuffix(asString, suffix) {
			return nil, issue
		}
		return asString, nil
	}
}

// Equals check if the provided input is the same as the required input.
// TODO: merge DateEquals into this
// TODO: extend to allow comparison of numbers
func Equals(targetName string, targetValue any) Rule {
	return func(input any) (any, error) {
		issue := Issue{
			Code:    CODE_EQUALS,
			Message: fmt.Sprintf("The input must be the same as '%s'", targetName),
			Value:   targetName, // TODO: confirm targetName or targetValue
		}
		if targetValue != input {
			return nil, issue
		}
		return targetValue, nil
	}
}

// Enum check if the provided input matches any of the listed enumerations values
// TODO: allow numeric enums has well
func Enum(enumValues ...string) Rule {
	return func(input any) (any, error) {
		issue := Issue{
			Code:    CODE_ENUM,
			Message: fmt.Sprintf("The input must match values %s", strings.Join(enumValues, ", ")),
			Value:   enumValues,
		}
		asString, ok := input.(string)
		if !ok || !slices.Contains(enumValues, asString) {
			return nil, issue
		}
		return asString, nil
	}
}

// URL check if the provided input is a valid string and a valid URL.
func URL(input any) (any, error) {
	issue := Issue{
		Code:    CODE_URL,
		Message: "Please provide a valid URL",
	}
	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	_, err := url.ParseRequestURI(asString)
	if err != nil {
		return nil, issue
	}
	return asString, nil
}

// Regexp check if the provided input is a valid string and matches the required
// regular expression.
func Regexp(pattern string) Rule {
	return func(input any) (any, error) {
		issue := Issue{
			Code:    CODE_REGEXP,
			Message: "The input doesn't match the required pattern",
		}
		asString, ok := input.(string)
		if !ok {
			return nil, issue
		}

		match, errMatch := regexp.MatchString(pattern, asString)
		if errMatch != nil || !match {
			return nil, issue
		}
		return asString, nil
	}
}

// UUID check if the provided input is a valid string and a valid UUID.
func UUID(input any) (any, error) {
	issue := Issue{
		Code:    CODE_UUID,
		Message: "Please provide a valid UUID string",
	}
	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	match, errMatch := regexp.MatchString(PATTERN_UUID, asString)
	if errMatch != nil || !match {
		return nil, issue
	}
	return asString, nil
}

// Password check if the provided input is a valid string and a reasonably strong
// password.
func Password(input any) (any, error) {
	issue := Issue{
		Code:    CODE_PASSWORD,
		Message: "Please provide a stronger password",
	}
	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	match, errMatch := regexp.MatchString(PATTERN_PASSWORD_STRENGTH, asString)

	// NOTE: regexp pattern being used matches invalid passwords instead of
	// strong passwords. If match is true, it means password was weak.
	if errMatch != nil || match {
		return nil, issue
	}
	return asString, nil
}

// JSON check if the provided code is a valid string and a valid json.
func JSON(input any) (any, error) {
	issue := Issue{
		Code:    CODE_JSON,
		Message: "Please provide a valid JSON string",
	}
	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	var target any
	errUnmarshal := json.Unmarshal([]byte(asString), &target)
	if errUnmarshal != nil {
		return nil, issue
	}
	return asString, nil
}

// DateTime check if the provided input is a valid string and a valid ISO
// timestamp according to RFC3339: [Link](https://pkg.go.dev/time#pkg-constants).
func DateTime(input any) (any, error) {
	issue := Issue{
		Code:    CODE_DATE_TIME,
		Message: "Please provide a valid date",
	}
	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	asDate, parseErr := time.Parse(time.RFC3339, asString)
	if parseErr != nil {
		return nil, issue
	}
	return asDate, nil
}

// Date check if the provided input is a valid date-only string. Date string
// must be in format e.g. 2023-10-05
// [Link](https://pkg.go.dev/time#pkg-constants).
func Date(input any) (any, error) {
	issue := Issue{
		Code:    CODE_DATE,
		Message: "Please provide a valid date",
	}
	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	asDate, parseErr := time.Parse(time.DateOnly, asString)
	if parseErr != nil {
		return nil, issue
	}
	return asDate, nil
}

// Time check if the provided input is a valid string and a valid time-only
// string. Time string must be in 24-hours format: e.g. 10:20:00
// [Link](https://pkg.go.dev/time#pkg-constants).
func Time(input any) (any, error) {
	issue := Issue{
		Code:    CODE_TIME,
		Message: "Please provide a valid time",
	}
	asString, ok := input.(string)
	if !ok {
		return nil, issue
	}

	asTime, parseErr := time.Parse(time.TimeOnly, asString)
	if parseErr != nil {
		return nil, issue
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
			return nil, Issue{
				Code:    CODE_DATE_EQUAL,
				Message: "The provided date must be " + target.String(),
				Value:   target.String(),
			}
		}

		return inputAsTime, nil
	}
}

// DateBefore check if the provided input is a date before (but not equal) to
// the target date.
// TODO: merge into LessThan
func DateBefore(target time.Time, inclusive bool) Rule {
	return func(input any) (any, error) {
		inputAsTime, ok := input.(time.Time)
		if !ok {
			return nil, errors.New("please provide a valid date")
		}

		delta := inputAsTime.Sub(target)
		if !inclusive && delta >= 0 {
			return nil, Issue{
				Code:    CODE_DATE_BEFORE,
				Message: "The provided date must be before " + target.String(),
				Value:   target.String(),
			}
		}

		if inclusive && delta > 0 {
			return nil, Issue{
				Code:    CODE_DATE_BEFORE,
				Message: "The provided date must be before or equal to " + target.String(),
				Value:   target.String(),
			}
		}

		return inputAsTime, nil
	}
}

// DateAfter check if the provided input is a date after the target date. If
// inclusive is set to true, target date will be included.
// TODO: merge into GreaterThan
func DateAfter(target time.Time, inclusive bool) Rule {
	return func(input any) (any, error) {
		inputAsTime, ok := input.(time.Time)
		if !ok {
			return nil, errors.New("please provide a valid date")
		}

		delta := inputAsTime.Sub(target)
		if !inclusive && delta <= 0 {
			return nil, Issue{
				Code:    CODE_DATE_AFTER,
				Message: "The provided date must be after " + target.String(),
				Value:   target.String(),
			}
		}

		if inclusive && delta < 0 {
			return nil, Issue{
				Code:    CODE_DATE_AFTER,
				Message: "The provided date must be after or equal to " + target.String(),
				Value:   target.String(),
			}
		}

		return inputAsTime, nil
	}
}

// Latitude check if the provided input a valid map latitude value.
func Latitude(input any) (any, error) {
	issue := Issue{
		Code:    CODE_LATITUDE,
		Message: "Please provide a valid latitude value",
	}

	inputAsFloat, ok := input.(float32)
	if !ok {
		return nil, issue
	}

	if inputAsFloat < -90.0 || inputAsFloat > 90.0 {
		return nil, issue
	}

	return inputAsFloat, nil
}

// Longitude check if the provided input a valid map longitude value.
func Longitude(input any) (any, error) {
	issue := Issue{
		Code:    CODE_LONGITUDE,
		Message: "Please provide a valid longitude value",
	}

	inputAsFloat, ok := input.(float32)
	if !ok {
		return nil, issue
	}

	if inputAsFloat < -180.0 || inputAsFloat > 180.0 {
		return nil, issue
	}

	return inputAsFloat, nil
}
