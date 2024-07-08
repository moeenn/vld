package vld

import (
	"testing"
	"time"
)

const (
	errValidFailed       = "valid input returned as invalid: %s"
	errInvalidPassed     = "invalid input returned as valid"
	errInvalidTypePassed = "invalid input type returned as valid"
	errInvalidReturnType = "invalid type of data returned by validation rule"
)

/**
 * Rule: NonEmptyString
 *
 */
func TestNonEmptyStringValid(t *testing.T) {
	input := "Some non-empty string"
	v, err := NonEmptyString(input)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestNonEmptyStringInvalidEmpty(t *testing.T) {
	_, err := NonEmptyString("")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestNonEmptyStringInvalidType(t *testing.T) {
	_, err := NonEmptyString(10)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Length
 *
 */
func TestLengthValid(t *testing.T) {
	v, err := Length(5)("q1w2e")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestLengthInvalid(t *testing.T) {
	_, err := Length(5)("q1w2e3r4")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = Length(4)("q1w")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestLengthInvalidType(t *testing.T) {
	_, err := Length(4)(300)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Min
 *
 */
func TestMinValidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 50, target: 10},
		{input: 50, target: 8.5},
		{input: 20.55, target: 10},
		{input: 20.55, target: 8.5},
		{input: "Hello world", target: 10},
	}

	for _, testCase := range testCases {
		_, err := Min(testCase.target)(testCase.input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}
	}
}

func TestMinInvalidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 50, target: 100},
		{input: 50, target: 80.5},
		{input: 20.55, target: 100},
		{input: 20.55, target: 80.5},
		{input: "Hello world", target: 20},
	}

	for _, testCase := range testCases {
		_, err := Min(testCase.target)(testCase.input)
		if err == nil {
			t.Errorf(errInvalidPassed)
			return
		}
	}
}

func TestMinInvalidTypeInput(t *testing.T) {
	_, err := Min(8)(true)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Max
 *
 */
func TestMaxValidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 50, target: 100},
		{input: 50.23, target: 80.5},
		{input: 20.55, target: 100},
		{input: 20.55, target: 800.5},
		{input: "Hello world", target: 20},
	}

	for _, testCase := range testCases {
		_, err := Max(testCase.target)(testCase.input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}
	}
}

func TestMaxInvalidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 20, target: 10},
		{input: 50.56, target: 30.5},
		{input: 25, target: 10.5},
		{input: 20.55, target: 10},
		{input: "Hello world", target: 10},
	}

	for _, testCase := range testCases {
		_, err := Max(testCase.target)(testCase.input)
		if err == nil {
			t.Errorf(errInvalidPassed)
			return
		}
	}
}

func TestMaxInvalidType(t *testing.T) {
	_, err := Max(8)(false)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: LessThan
 *
 */
func TestLessThanValidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 50, target: 400},
		{input: 50.23, target: 80.5},
		{input: 20.55, target: 100},
		{input: 20.55, target: 50.5},
		{input: "Hello world", target: 20},
	}

	for _, testCase := range testCases {
		_, err := LessThan(testCase.target)(testCase.input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}
	}
}

func TestLessThanInvalidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 20, target: 10},
		{input: 50.56, target: 30.5},
		{input: 25, target: 10.5},
		{input: 20.55, target: 10},
		{input: "Hello world", target: 5},
	}

	for _, testCase := range testCases {
		_, err := LessThan(testCase.target)(testCase.input)
		if err == nil {
			t.Errorf(errInvalidPassed)
			return
		}
	}
}

func TestLessThanInvalidType(t *testing.T) {
	_, err := LessThan(8)(false)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: GreaterThan
 *
 */
func TestGreaterThanValidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 20, target: 10},
		{input: 50.56, target: 30.5},
		{input: 25, target: 10.5},
		{input: 20.55, target: 10},
		{input: "Hello world", target: 5},
	}

	for _, testCase := range testCases {
		_, err := GreaterThan(testCase.target)(testCase.input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}
	}
}

func TestGreaterThanInvalidInputs(t *testing.T) {
	type testCase struct {
		input  any
		target any
	}

	testCases := []testCase{
		{input: 50, target: 400},
		{input: 50.23, target: 80.5},
		{input: 20.55, target: 100},
		{input: 20.55, target: 50.5},
		{input: "Hello world", target: 20},
	}

	for _, testCase := range testCases {
		_, err := GreaterThan(testCase.target)(testCase.input)
		if err == nil {
			t.Errorf(errInvalidPassed)
			return
		}
	}
}

func TestGreaterThanInvalidType(t *testing.T) {
	_, err := GreaterThan(8)(false)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Email
 *
 */
func TestEmailValid(t *testing.T) {
	v, err := Email("admin@site.com")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestEmailInvalidInput(t *testing.T) {
	_, err := Email("random.ascalscn")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestEmailInvalidType(t *testing.T) {
	_, err := Email(400)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: HasPrefix
 *
 */
func TestHasPrefixValid(t *testing.T) {
	v, err := HasPrefix("sample")("sample input")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestHasPrefixInvalidInput(t *testing.T) {
	_, err := HasPrefix("id_")("random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestHasPrefixInvalidInputType(t *testing.T) {
	_, err := HasPrefix("id_")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: HasSuffix
 *
 */
func TestHasSuffixValid(t *testing.T) {
	v, err := HasSuffix("sample")("input sample")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestHasSuffixInvalidInput(t *testing.T) {
	_, err := HasSuffix("_end")("random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestHasSuffixInvalidInputType(t *testing.T) {
	_, err := HasSuffix("id")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: NotHasPrefix
 *
 */
func TestNotHasPrefixValid(t *testing.T) {
	v, err := NotHasPrefix("sample")("1sample input")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestNotHasPrefixInvalidInput(t *testing.T) {
	_, err := NotHasPrefix("id_")("id_random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestNotHasPrefixInvalidInputType(t *testing.T) {
	_, err := NotHasPrefix("id_")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: NotHasSuffix
 *
 */
func TestNotHasSuffixValid(t *testing.T) {
	v, err := NotHasSuffix("sample")("input samplex")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestNotHasSuffixInvalidInput(t *testing.T) {
	_, err := NotHasSuffix("_end")("random_input_end")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestNotHasSuffixInvalidInputType(t *testing.T) {
	_, err := NotHasSuffix("id")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: Equals
 *
 */
func TestEqualsValidInput(t *testing.T) {
	v, err := Equals("Password", "confirmed_password")("confirmed_password")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = Equals("Payment", 300.5)(300.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestEqualsInvalidInput(t *testing.T) {
	_, err := Equals("Repo name", "github.com/sample/example")("github.com/example/sample")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = Equals("Random", false)(true)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestEqualsMismatchBetweenTypes(t *testing.T) {
	_, err := Equals("Example", 40.55)(true)
	if err == nil {
		t.Errorf(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Enum
 *
 */
func TestEnumValidInput(t *testing.T) {
	v, err := Enum("A", "B", "C")("B")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestEnumInvalidInput(t *testing.T) {
	_, err := Enum("Left", "Right", "Up", "Down")("Middle")
	if err == nil {
		t.Errorf(errInvalidPassed)
		return
	}
}

func TestEnumInvalidInputType(t *testing.T) {
	_, err := Enum("A", "B", "C")(400.666)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

/**
 * Rule: URL
 *
 */
func TestURLValidInput(t *testing.T) {
	v, err := URL("https://site.com/abc?some=random")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestURLInvalidInput(t *testing.T) {
	_, err := URL("not-a-valid-url")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = URL("google.com")
	if err == nil {
		t.Errorf(errInvalidPassed)
		return
	}
}

func TestURLInvalidInputType(t *testing.T) {
	_, err := URL(500_000.67)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Regexp
 *
 */
func TestRegexpValidInput(t *testing.T) {
	v, err := Regexp("^[a-z0-9_-]{3,16}$")("random_alpha_321")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestRegexpInvalidInput(t *testing.T) {
	_, err := Regexp("^hello$")("another")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestRegexpInvalidInputType(t *testing.T) {
	_, err := Regexp("^[a-z0-9]+(?:-[a-z0-9]+)*$")(500_30.22)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: UUID
 *
 */
func TestUUIDValidInput(t *testing.T) {
	v, err := UUID("1b810d1a-0f3b-4bff-86fe-039258c5b20e")
	if err != nil {
		t.Errorf(errInvalidPassed)
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestUUIDInvalidInput(t *testing.T) {
	_, err := UUID("some.random-string.not-uuid")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestUUIDInvalidInputType(t *testing.T) {
	_, err := UUID(true)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Password
 *
 */
func TestPasswordValidInput(t *testing.T) {
	v, err := Password("Q1w2e3r4#!@")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestPasswordInvalidInput(t *testing.T) {
	_, err := Password("abc123")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestPasswordInvalidInputType(t *testing.T) {
	_, err := Password(40_000_234.4)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: JSON
 *
 */
func TestJSONValidInput(t *testing.T) {
	input := `{"a":1,"b":"random","c":true,"d":[1,2,3]}`
	v, err := JSON(input)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestJSONInvalidInput(t *testing.T) {
	input := `{"a":1,"b':"random"`
	_, err := JSON(input)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestJSONInvalidInputType(t *testing.T) {
	_, err := JSON(true)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: DateTime
 *
 */
func TestDateTimeValidInput(t *testing.T) {
	inputs := []string{
		"2024-03-22T12:35:05.115Z",
		"2024-03-22T05:11:09.762Z",
		"2024-03-20T10:17:40.521Z",
	}

	for _, input := range inputs {
		v, err := DateTime(input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
		}

		if _, ok := v.(time.Time); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestDateTimeInvalidInput(t *testing.T) {
	inputs := []string{
		"2020-10-15",
		"2024-15-22T05:11:09.762Z",
		"2024-03-40T10:17:40.521Z",
	}

	for _, input := range inputs {
		_, err := DateTime(input)
		if err == nil {
			t.Error(errInvalidPassed)
		}
	}
}

func TestDateTimeInvalidInputType(t *testing.T) {
	_, err := DateTime(30.44)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: Date
 *
 */
func TestDateValidInput(t *testing.T) {
	inputs := []string{
		"2024-03-22",
		"2024-02-29",
		"1970-01-01",
	}

	for _, input := range inputs {
		v, err := Date(input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
		}

		if _, ok := v.(time.Time); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestDateInvalidInput(t *testing.T) {
	inputs := []string{
		"2020/10/15",
		"2024-43-22",
		"2020-02-30",
		"2024-03-10T10:17:40.521Z",
	}

	for _, input := range inputs {
		_, err := Date(input)
		if err == nil {
			t.Error(errInvalidPassed)
		}
	}
}

func TestDateInvalidInputType(t *testing.T) {
	_, err := Date(false)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: Time
 *
 */
func TestTimeValidInput(t *testing.T) {
	inputs := []string{
		"11:30:00",
		"15:00:00",
	}

	for _, input := range inputs {
		v, err := Time(input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
		}

		if _, ok := v.(time.Time); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestTimeInvalidInput(t *testing.T) {
	inputs := []string{
		"25:00:00",
		"10:70:00",
		"02:15:00 AM",
		"2024-03-10T10:17:40.521Z",
	}

	for _, input := range inputs {
		_, err := Time(input)
		if err == nil {
			t.Error(errInvalidPassed)
		}
	}
}

func TestTimeInvalidInputType(t *testing.T) {
	_, err := Time(true)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: DateEqual
 *
 */
func TestDateEqualValidInput(t *testing.T) {
	target := time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC)
	input := time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC) // same as target

	result, err := DateEqual(target)(input)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := result.(time.Time); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestDateEqualInvalidInput(t *testing.T) {
	inputs := []time.Time{
		time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 5, 12, 0, 0, 0, 0, time.UTC),
	}

	target := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	v := DateEqual(target)
	for _, input := range inputs {
		_, err := v(input)
		if err == nil {
			t.Error(errInvalidPassed)
			return
		}
	}
}

func TestDateEqualInvalidInputType(t *testing.T) {
	v := DateEqual(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC))
	if _, err := v(true); err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: DateBefore
 *
 */
func TestDateBeforeValidInput(t *testing.T) {
	target := time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		date      time.Time
		inclusive bool
	}{
		{date: time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(2005, 1, 2, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(1970, 2, 5, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC), inclusive: true},
		{date: target, inclusive: true},
	}

	for _, testCase := range testCases {
		result, err := DateBefore(target, testCase.inclusive)(testCase.date)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}

		if _, ok := result.(time.Time); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestDateBeforeInvalidInput(t *testing.T) {
	target := time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		date      time.Time
		inclusive bool
	}{
		{date: time.Date(2030, 10, 23, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(2050, 1, 2, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(3050, 2, 5, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(2035, 2, 5, 0, 0, 0, 0, time.UTC), inclusive: true},
		{date: target, inclusive: false},
	}

	for _, testCase := range testCases {
		_, err := DateBefore(target, testCase.inclusive)(testCase.date)
		if err == nil {
			t.Error(errInvalidPassed)
			return
		}
	}
}

func TestDateBeforeInvalidInputType(t *testing.T) {
	v := DateBefore(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC), false)
	if _, err := v(false); err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: DateAfter
 *
 */
func TestDateAfterValidInput(t *testing.T) {
	target := time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		date      time.Time
		inclusive bool
	}{
		{date: time.Date(2030, 10, 23, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(2050, 1, 2, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(3050, 2, 5, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), inclusive: true},
		{date: target, inclusive: true},
	}

	for _, testCase := range testCases {
		result, err := DateAfter(target, testCase.inclusive)(testCase.date)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}

		if _, ok := result.(time.Time); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestDateAfterInvalidInput(t *testing.T) {
	target := time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		date      time.Time
		inclusive bool
	}{
		{date: time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(2005, 1, 2, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(1970, 2, 5, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC), inclusive: false},
		{date: target, inclusive: false},
	}

	for _, testCase := range testCases {
		_, err := DateAfter(target, testCase.inclusive)(testCase.date)
		if err == nil {
			t.Error(errInvalidPassed)
			return
		}
	}
}

func TestDateAfterInvalidInputType(t *testing.T) {
	v := DateAfter(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC), true)
	if _, err := v(nil); err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Latitude
 *
 */
func TestLatitudeValidInput(t *testing.T) {
	inputs := []float32{31.475240, 67.724989, 82.326122, -6.538025}
	for _, lat := range inputs {
		result, err := Latitude(lat)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}

		if _, ok := result.(float32); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestLatitudeInvalidInput(t *testing.T) {
	inputs := []float32{-91.475240, 100.724989}
	for _, lat := range inputs {
		_, err := Latitude(lat)
		if err == nil {
			t.Error(errInvalidPassed)
			return
		}
	}
}

func TestLatitudeInvalidInputType(t *testing.T) {
	if _, err := Latitude("abc"); err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: Longitude
 *
 */
func TestLongitudeValidInput(t *testing.T) {
	inputs := []float32{74.365614, 176.140846, -77.687278, 23.554154}
	for _, lng := range inputs {
		result, err := Longitude(lng)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
			return
		}

		if _, ok := result.(float32); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestLongitudeInvalidInput(t *testing.T) {
	inputs := []float32{-191.475240, 190.724989}
	for _, lng := range inputs {
		_, err := Longitude(lng)
		if err == nil {
			t.Error(errInvalidPassed)
			return
		}
	}
}

func TestLongitudeInvalidInputType(t *testing.T) {
	if _, err := Longitude(true); err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}
