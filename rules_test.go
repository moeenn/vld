package vld

import (
	"testing"
)

/**
 * Rule: NonEmptyString
 *
 */
func TestNonEmptyStringValid(t *testing.T) {
	input := "Some non-empty string"
	err := NonEmptyString(input)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestNonEmptyStringInvalidEmpty(t *testing.T) {
	input := ""
	err := NonEmptyString(input)
	if err == nil {
		t.Error("empty string returned as valid")
		return
	}
}

func TestNonEmptyStringInvalidType(t *testing.T) {
	err := NonEmptyString(10)
	if err == nil {
		t.Error("NonEmptyString: invalid type input successfully validated")
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
		t.Errorf("valid data returned as invalid: %+v", err.Error())
		return
	}
}

func TestLengthInvalid(t *testing.T) {
	err := Length(5)("q1w2e3r4")
	if err == nil {
		t.Error("invalid data returned as valid")
		return
	}

	err = Length(4)("q1w")
	if err == nil {
		t.Error("invalid data returned as valid")
		return
	}
}

func TestLengthInvalidType(t *testing.T) {
	err := Length(4)(300)
	if err == nil {
		t.Error("invalid data type returned as valid")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = MinLength(4)("q1w2")
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestMinLengthInvalidInput(t *testing.T) {
	err := MinLength(8)("q1w2e3")
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestMinLengthInvalidTypeInput(t *testing.T) {
	err := MinLength(8)(true)
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = MaxLength(4)("q1w")
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestMaxLengthInvalidInput(t *testing.T) {
	err := MaxLength(8)("q1w2e3r4t5")
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestMaxLengthInvalidType(t *testing.T) {
	err := MaxLength(8)(false)
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = MinFloat(-100.0)(50.5)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestMinFloatInvalidInput(t *testing.T) {
	err := MinFloat(10.4)(2.6)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = MinFloat(-50.333)(-80.3)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestMinFloatInvalidTypeInput(t *testing.T) {
	err := MinFloat(20.0)("random")
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = MinFloat(200.6)(170)
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = MaxFloat(-100.0)(-200.5)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestMaxFloatInvalidInput(t *testing.T) {
	err := MaxFloat(2.6)(10.666)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = MaxFloat(-50.333)(-30.3)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestMaxFloatInvalidTypeInput(t *testing.T) {
	err := MaxFloat(20.0)("random")
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = MaxFloat(200.6)(170)
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = MinInt(-200)(50)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestMinIntInvalidInput(t *testing.T) {
	err := MinInt(10)(2)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = MinInt(-50)(-90)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestMinIntInvalidTypeInput(t *testing.T) {
	err := MinInt(20)(false)
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = MinInt(-200)(-270.75)
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = MaxInt(-100)(-200)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestMaxIntInvalidInput(t *testing.T) {
	err := MaxInt(2)(10)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = MaxInt(-50)(-30)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestMaxIntInvalidTypeInput(t *testing.T) {
	err := MaxInt(20)(10.66666)
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = MaxInt(200)("random")
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = LessThanInt(-90)(-100)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestLessThanIntInvalidInput(t *testing.T) {
	err := LessThanInt(-3000)(2)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = LessThanInt(-80)(-50)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestLessThanIntInvalidTypeInput(t *testing.T) {
	err := LessThanInt(30)(20.66666)
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = LessThanInt(200)("random")
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = LessThanFloat(-90.55)(-100.58585)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestLessThanFloatInvalidInput(t *testing.T) {
	err := LessThanFloat(-3000.21)(24.44)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = LessThanFloat(-80.31231)(-50.6554)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestLessThanFloatInvalidTypeInput(t *testing.T) {
	err := LessThanFloat(30.55)(20)
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = LessThanFloat(200.6575)("random")
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = GreaterThanInt(-100)(-90)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestGreaterThanIntInvalidInput(t *testing.T) {
	err := GreaterThanInt(2)(-3000)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = GreaterThanInt(-50)(-80)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestGreaterThanIntInvalidTypeInput(t *testing.T) {
	err := GreaterThanInt(20)(30.66666)
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = GreaterThanInt(200)("random")
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}

	err = GreaterThanFloat(-100.58585)(-90.54)
	if err != nil {
		t.Errorf("validation failed for valid input: %+v", err.Error())
		return
	}
}

func TestGreaterThanFloatInvalidInput(t *testing.T) {
	err := GreaterThanFloat(2.44)(-3000.21)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}

	err = GreaterThanFloat(-50.678)(-80.31231)
	if err == nil {
		t.Error("invalid input validated successfully")
		return
	}
}

func TestGreaterThanFloatInvalidTypeInput(t *testing.T) {
	err := GreaterThanFloat(20.55)(30)
	if err == nil {
		t.Error("invalid type input validated successfully")
		return
	}

	err = GreaterThanFloat(200.6575)("random")
	if err == nil {
		t.Error("invalid type input validated successfully")
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
		t.Errorf("valid email returned as invalid: %s", err.Error())
		return
	}
}

func TestEmailInvalidInput(t *testing.T) {
	err := Email("random.ascalscn")
	if err == nil {
		t.Error("invalid email returned as valid")
		return
	}
}

func TestEmailInvalidType(t *testing.T) {
	err := Email(400)
	if err == nil {
		t.Error("invalid input type returned as valid")
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
		t.Errorf("valid input returned as invalid: %s", err.Error())
		return
	}
}

func TestStartsWithInvalidInput(t *testing.T) {
	err := StartsWith("id_")("random_input")
	if err == nil {
		t.Error("invalid input returned as valid")
		return
	}
}

func TestStartsWithInvalidInputType(t *testing.T) {
	err := StartsWith("id_")(4000)
	if err == nil {
		t.Error("invalid input type returned as valid")
	}
}

/**
 * Rule: EndsWith
 *
 */
func TestEndsWithValid(t *testing.T) {
	err := EndsWith("sample")("input sample")
	if err != nil {
		t.Errorf("valid input returned as invalid: %s", err.Error())
		return
	}
}

func TestEndsWithInvalidInput(t *testing.T) {
	err := EndsWith("_end")("random_input")
	if err == nil {
		t.Error("invalid input returned as valid")
		return
	}
}

func TestEndsWithInvalidInputType(t *testing.T) {
	err := EndsWith("id")(4000)
	if err == nil {
		t.Error("invalid input type returned as valid")
	}
}

/**
 * Rule: DoesntStartWith
 *
 */
func TestDoesntStartWithValid(t *testing.T) {
	err := DoesntStartWith("sample")("1sample input")
	if err != nil {
		t.Errorf("valid input returned as invalid: %s", err.Error())
		return
	}
}

func TestDoesntStartWithInvalidInput(t *testing.T) {
	err := DoesntStartWith("id_")("id_random_input")
	if err == nil {
		t.Error("invalid input returned as valid")
		return
	}
}

func TestDoesntStartWithInvalidInputType(t *testing.T) {
	err := DoesntStartWith("id_")(4000)
	if err == nil {
		t.Error("invalid input type returned as valid")
	}
}

/**
 * Rule: DoesntEndWith
 *
 */
func TestDoesntEndWithValid(t *testing.T) {
	err := DoesntEndWith("sample")("input samplex")
	if err != nil {
		t.Errorf("valid input returned as invalid: %s", err.Error())
		return
	}
}

func TestDoesntEndWithInvalidInput(t *testing.T) {
	err := DoesntEndWith("_end")("random_input_end")
	if err == nil {
		t.Error("invalid input returned as valid")
		return
	}
}

func TestDoesntEndWithInvalidInputType(t *testing.T) {
	err := DoesntEndWith("id")(4000)
	if err == nil {
		t.Error("invalid input type returned as valid")
	}
}
