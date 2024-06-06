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
 * Rule: MinLength
 *
 */
func TestMinLengthValid(t *testing.T) {
	_, err := MinLength(8)("q1w2e3r4t5")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	v, err := MinLength(4)("q1w2")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestMinLengthInvalidInput(t *testing.T) {
	_, err := MinLength(8)("q1w2e3")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMinLengthInvalidTypeInput(t *testing.T) {
	_, err := MinLength(8)(true)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: MaxLength
 *
 */
func TestMaxLengthValid(t *testing.T) {
	_, err := MaxLength(8)("q1w2e3r4")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	v, err := MaxLength(4)("q1w")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(string); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestMaxLengthInvalidInput(t *testing.T) {
	_, err := MaxLength(8)("q1w2e3r4t5")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMaxLengthInvalidType(t *testing.T) {
	_, err := MaxLength(8)(false)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: MinFloat
 *
 */
func TestMinFloatValid(t *testing.T) {
	v, err := MinFloat(0)(0.0)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = MinFloat(-100.0)(50.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(float64); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestMinFloatInvalidInput(t *testing.T) {
	_, err := MinFloat(10.4)(2.6)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = MinFloat(-50.333)(-80.3)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMinFloatInvalidTypeInput(t *testing.T) {
	_, err := MinFloat(20.0)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = MinFloat(200.6)(170)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: MaxFloat
 *
 */
func TestMaxFloatValid(t *testing.T) {
	v, err := MaxFloat(10.0)(10.0)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = MaxFloat(-100.0)(-200.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(float64); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestMaxFloatInvalidInput(t *testing.T) {
	_, err := MaxFloat(2.6)(10.666)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = MaxFloat(-50.333)(-30.3)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMaxFloatInvalidTypeInput(t *testing.T) {
	_, err := MaxFloat(20.0)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = MaxFloat(200.6)(170)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: MinInt
 *
 */
func TestMinIntValid(t *testing.T) {
	v, err := MinInt(10)(20)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = MinInt(-200)(50)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(int); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestMinIntInvalidInput(t *testing.T) {
	_, err := MinInt(10)(2)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = MinInt(-50)(-90)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMinIntInvalidTypeInput(t *testing.T) {
	_, err := MinInt(20)(false)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = MinInt(-200)(-270.75)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: MaxInt
 *
 */
func TestMaxIntValid(t *testing.T) {
	v, err := MaxInt(10)(10)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = MaxInt(-100)(-200)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(int); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestMaxIntInvalidInput(t *testing.T) {
	_, err := MaxInt(2)(10)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = MaxInt(-50)(-30)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMaxIntInvalidTypeInput(t *testing.T) {
	_, err := MaxInt(20)(10.66666)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = MaxInt(200)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: LessThanInt
 *
 */
func TestLessThanIntValid(t *testing.T) {
	v, err := LessThanInt(11)(10)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = LessThanInt(-90)(-100)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(int); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestLessThanIntInvalidInput(t *testing.T) {
	_, err := LessThanInt(-3000)(2)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = LessThanInt(-80)(-50)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestLessThanIntInvalidTypeInput(t *testing.T) {
	_, err := LessThanInt(30)(20.66666)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = LessThanInt(200)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: LessThanFloat
 *
 */
func TestLessThanFloatValid(t *testing.T) {
	v, err := LessThanFloat(200.5)(10.66)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = LessThanFloat(-90.55)(-100.58585)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(float64); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestLessThanFloatInvalidInput(t *testing.T) {
	_, err := LessThanFloat(-3000.21)(24.44)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = LessThanFloat(-80.31231)(-50.6554)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestLessThanFloatInvalidTypeInput(t *testing.T) {
	_, err := LessThanFloat(30.55)(20)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = LessThanFloat(200.6575)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: GreaterThanInt
 *
 */
func TestGreaterThanIntValid(t *testing.T) {
	v, err := GreaterThanInt(10)(11)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = GreaterThanInt(-100)(-90)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(int); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestGreaterThanIntInvalidInput(t *testing.T) {
	_, err := GreaterThanInt(2)(-3000)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = GreaterThanInt(-50)(-80)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestGreaterThanIntInvalidTypeInput(t *testing.T) {
	_, err := GreaterThanInt(20)(30.66666)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = GreaterThanInt(200)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}

/**
 * Rule: GreaterThanFloat
 *
 */
func TestGreaterThanFloatValid(t *testing.T) {
	v, err := GreaterThanFloat(10.66)(200.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	_, err = GreaterThanFloat(-100.58585)(-90.54)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	if _, ok := v.(float64); !ok {
		t.Error(errInvalidReturnType)
		return
	}
}

func TestGreaterThanFloatInvalidInput(t *testing.T) {
	_, err := GreaterThanFloat(2.44)(-3000.21)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	_, err = GreaterThanFloat(-50.678)(-80.31231)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestGreaterThanFloatInvalidTypeInput(t *testing.T) {
	_, err := GreaterThanFloat(20.55)(30)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	_, err = GreaterThanFloat(200.6575)("random")
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
 * Rule: ISODate
 *
 */
func TestISODateValidInput(t *testing.T) {
	inputs := []string{
		"2024-03-22T12:35:05.115Z",
		"2024-03-22T05:11:09.762Z",
		"2024-03-20T10:17:40.521Z",
	}

	for _, input := range inputs {
		v, err := ISODate(input)
		if err != nil {
			t.Errorf(errValidFailed, err.Error())
		}

		if _, ok := v.(time.Time); !ok {
			t.Error(errInvalidReturnType)
			return
		}
	}
}

func TestISODateInvalidInput(t *testing.T) {
	inputs := []string{
		"2020-10-15",
		"2024-15-22T05:11:09.762Z",
		"2024-03-40T10:17:40.521Z",
	}

	for _, input := range inputs {
		_, err := ISODate(input)
		if err == nil {
			t.Error(errInvalidPassed)
		}
	}
}

func TestISODateInvalidInputType(t *testing.T) {
	_, err := ISODate(30.44)
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
