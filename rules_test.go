package vld

import (
	"testing"
)

const (
	errValidFailed       = "valid input returned as invalid: %s"
	errInvalidPassed     = "invalid input returned as valid"
	errInvalidTypePassed = "invalid input type returned as valid"
)

/**
 * Rule: NonEmptyString
 *
 */
func TestNonEmptyStringValid(t *testing.T) {
	input := "Some non-empty string"
	err := NonEmptyString(input)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestNonEmptyStringInvalidEmpty(t *testing.T) {
	err := NonEmptyString("")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestNonEmptyStringInvalidType(t *testing.T) {
	err := NonEmptyString(10)
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
	err := Length(5)("q1w2e")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestLengthInvalid(t *testing.T) {
	err := Length(5)("q1w2e3r4")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = Length(4)("q1w")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestLengthInvalidType(t *testing.T) {
	err := Length(4)(300)
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
	err := MinLength(8)("q1w2e3r4t5")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = MinLength(4)("q1w2")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestMinLengthInvalidInput(t *testing.T) {
	err := MinLength(8)("q1w2e3")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMinLengthInvalidTypeInput(t *testing.T) {
	err := MinLength(8)(true)
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
	err := MaxLength(8)("q1w2e3r4")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = MaxLength(4)("q1w")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestMaxLengthInvalidInput(t *testing.T) {
	err := MaxLength(8)("q1w2e3r4t5")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMaxLengthInvalidType(t *testing.T) {
	err := MaxLength(8)(false)
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
	err := MinFloat(0)(0.0)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = MinFloat(-100.0)(50.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestMinFloatInvalidInput(t *testing.T) {
	err := MinFloat(10.4)(2.6)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = MinFloat(-50.333)(-80.3)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMinFloatInvalidTypeInput(t *testing.T) {
	err := MinFloat(20.0)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = MinFloat(200.6)(170)
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
	err := MaxFloat(10.0)(10.0)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = MaxFloat(-100.0)(-200.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestMaxFloatInvalidInput(t *testing.T) {
	err := MaxFloat(2.6)(10.666)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = MaxFloat(-50.333)(-30.3)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMaxFloatInvalidTypeInput(t *testing.T) {
	err := MaxFloat(20.0)("random")
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = MaxFloat(200.6)(170)
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
	err := MinInt(10)(20)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = MinInt(-200)(50)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestMinIntInvalidInput(t *testing.T) {
	err := MinInt(10)(2)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = MinInt(-50)(-90)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMinIntInvalidTypeInput(t *testing.T) {
	err := MinInt(20)(false)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = MinInt(-200)(-270.75)
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
	err := MaxInt(10)(10)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = MaxInt(-100)(-200)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestMaxIntInvalidInput(t *testing.T) {
	err := MaxInt(2)(10)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = MaxInt(-50)(-30)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestMaxIntInvalidTypeInput(t *testing.T) {
	err := MaxInt(20)(10.66666)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = MaxInt(200)("random")
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
	err := LessThanInt(11)(10)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = LessThanInt(-90)(-100)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestLessThanIntInvalidInput(t *testing.T) {
	err := LessThanInt(-3000)(2)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = LessThanInt(-80)(-50)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestLessThanIntInvalidTypeInput(t *testing.T) {
	err := LessThanInt(30)(20.66666)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = LessThanInt(200)("random")
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
	err := LessThanFloat(200.5)(10.66)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = LessThanFloat(-90.55)(-100.58585)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestLessThanFloatInvalidInput(t *testing.T) {
	err := LessThanFloat(-3000.21)(24.44)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = LessThanFloat(-80.31231)(-50.6554)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestLessThanFloatInvalidTypeInput(t *testing.T) {
	err := LessThanFloat(30.55)(20)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = LessThanFloat(200.6575)("random")
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
	err := GreaterThanInt(10)(11)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = GreaterThanInt(-100)(-90)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestGreaterThanIntInvalidInput(t *testing.T) {
	err := GreaterThanInt(2)(-3000)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = GreaterThanInt(-50)(-80)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestGreaterThanIntInvalidTypeInput(t *testing.T) {
	err := GreaterThanInt(20)(30.66666)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = GreaterThanInt(200)("random")
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
	err := GreaterThanFloat(10.66)(200.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = GreaterThanFloat(-100.58585)(-90.54)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestGreaterThanFloatInvalidInput(t *testing.T) {
	err := GreaterThanFloat(2.44)(-3000.21)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = GreaterThanFloat(-50.678)(-80.31231)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestGreaterThanFloatInvalidTypeInput(t *testing.T) {
	err := GreaterThanFloat(20.55)(30)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}

	err = GreaterThanFloat(200.6575)("random")
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
	err := Email("admin@site.com")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestEmailInvalidInput(t *testing.T) {
	err := Email("random.ascalscn")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestEmailInvalidType(t *testing.T) {
	err := Email(400)
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
	err := StartsWith("sample")("sample input")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestStartsWithInvalidInput(t *testing.T) {
	err := StartsWith("id_")("random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestStartsWithInvalidInputType(t *testing.T) {
	err := StartsWith("id_")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: EndsWith
 *
 */
func TestEndsWithValid(t *testing.T) {
	err := EndsWith("sample")("input sample")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestEndsWithInvalidInput(t *testing.T) {
	err := EndsWith("_end")("random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestEndsWithInvalidInputType(t *testing.T) {
	err := EndsWith("id")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: DoesntStartWith
 *
 */
func TestDoesntStartWithValid(t *testing.T) {
	err := DoesntStartWith("sample")("1sample input")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestDoesntStartWithInvalidInput(t *testing.T) {
	err := DoesntStartWith("id_")("id_random_input")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestDoesntStartWithInvalidInputType(t *testing.T) {
	err := DoesntStartWith("id_")(4000)
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
	err := DoesntEndWith("sample")("input samplex")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestDoesntEndWithInvalidInput(t *testing.T) {
	err := DoesntEndWith("_end")("random_input_end")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestDoesntEndWithInvalidInputType(t *testing.T) {
	err := DoesntEndWith("id")(4000)
	if err == nil {
		t.Error(errInvalidTypePassed)
	}
}

/**
 * Rule: Same
 *
 */
func TestSameValidInput(t *testing.T) {
	err := Same("Password", "confirmed_password")("confirmed_password")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}

	err = Same("Payment", 300.5)(300.5)
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestSameInvalidInput(t *testing.T) {
	err := Same("Repo name", "github.com/sample/example")("github.com/example/sample")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = Same("Random", false)(true)
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestSameMismatchBetweenTypes(t *testing.T) {
	err := Same("Example", 40.55)(true)
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
	err := Enum("A", "B", "C")("B")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestEnumInvalidInput(t *testing.T) {
	err := Enum("Left", "Right", "Up", "Down")("Middle")
	if err == nil {
		t.Errorf(errInvalidPassed)
		return
	}
}

func TestEnumInvalidInputType(t *testing.T) {
	err := Enum("A", "B", "C")(400.666)
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
	err := URL("https://site.com/abc?some=random")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestURLInvalidInput(t *testing.T) {
	err := URL("not-a-valid-url")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}

	err = URL("google.com")
	if err == nil {
		t.Errorf(errInvalidPassed)
		return
	}
}

func TestURLInvalidInputType(t *testing.T) {
	err := URL(500_000.67)
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
	err := Regexp("^[a-z0-9_-]{3,16}$")("random_alpha_321")
	if err != nil {
		t.Errorf(errValidFailed, err.Error())
		return
	}
}

func TestRegexpInvalidInput(t *testing.T) {
	err := Regexp("^hello$")("another")
	if err == nil {
		t.Error(errInvalidPassed)
		return
	}
}

func TestRegexpInvalidInputType(t *testing.T) {
	err := Regexp("^[a-z0-9]+(?:-[a-z0-9]+)*$")(500_30.22)
	if err == nil {
		t.Error(errInvalidTypePassed)
		return
	}
}
