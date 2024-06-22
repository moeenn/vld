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
 * Rule: StartsWith
 *
 */
func TestStartsWithValid(t *testing.T) {
	v, err := StartsWith("sample")("sample input")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestStartsWithInvalidInput(t *testing.T) {
	_, err := StartsWith("id_")("random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestStartsWithInvalidInputType(t *testing.T) {
	_, err := StartsWith("id_")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: EndsWith
 *
 */
func TestEndsWithValid(t *testing.T) {
	v, err := EndsWith("sample")("input sample")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestEndsWithInvalidInput(t *testing.T) {
	_, err := EndsWith("_end")("random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestEndsWithInvalidInputType(t *testing.T) {
	_, err := EndsWith("id")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: DoesntStartWith
 *
 */
func TestDoesntStartWithValid(t *testing.T) {
	v, err := DoesntStartWith("sample")("1sample input")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestDoesntStartWithInvalidInput(t *testing.T) {
	_, err := DoesntStartWith("id_")("id_random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestDoesntStartWithInvalidInputType(t *testing.T) {
	_, err := DoesntStartWith("id_")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: DoesntEndWith
 *
 */
func TestDoesntEndWithValid(t *testing.T) {
	v, err := DoesntEndWith("sample")("input samplex")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestDoesntEndWithInvalidInput(t *testing.T) {
	_, err := DoesntEndWith("_end")("random_input_end")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestDoesntEndWithInvalidInputType(t *testing.T) {
	_, err := DoesntEndWith("id")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: Same
 *
 */
func TestSameValidInput(t *testing.T) {
	v, err := Same("Password", "confirmed_password")("confirmed_password")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = Same("Payment", 300.5)(300.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestSameInvalidInput(t *testing.T) {
	_, err := Same("Repo name", "github.com/sample/example")("github.com/example/sample")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = Same("Random", false)(true)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestSameMismatchBetweenTypes(t *testing.T) {
	_, err := Same("Example", 40.55)(true)
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
 * Rule: DateBefore
 *
 */
func TestDateBeforeValidInput(t *testing.T) {
	inputs := []time.Time{
		time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2005, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(1970, 2, 5, 0, 0, 0, 0, time.UTC),
		time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	v := DateBefore(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC))
	for _, input := range inputs {
		result, err := v(input)
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
	inputs := []time.Time{
		time.Date(2030, 10, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2050, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(3050, 2, 5, 0, 0, 0, 0, time.UTC),
		time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	v := DateBefore(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC))
	for _, input := range inputs {
		_, err := v(input)
		if err == nil {
			t.Error(errInvalidPassed)
			return
		}
	}
}

func TestDateBeforeInvalidInputType(t *testing.T) {
	v := DateBefore(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC))
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
	inputs := []time.Time{
		time.Date(2030, 10, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2050, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(3050, 2, 5, 0, 0, 0, 0, time.UTC),
		time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	v := DateAfter(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC))
	for _, input := range inputs {
		result, err := v(input)
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
	inputs := []time.Time{
		time.Date(2023, 10, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2005, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(1970, 2, 5, 0, 0, 0, 0, time.UTC),
		time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	v := DateAfter(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC))
	for _, input := range inputs {
		_, err := v(input)
		if err == nil {
			t.Error(errInvalidPassed)
			return
		}
	}
}

func TestDateAfterInvalidInputType(t *testing.T) {
	v := DateAfter(time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC))
	if _, err := v(nil); err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}
