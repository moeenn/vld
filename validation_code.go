package vld

const (
	CODE_NON_EMPTY_STRING = "non-empty-string"
	CODE_LENGTH           = "length"
	CODE_MIN              = "min"
	CODE_MAX              = "max"
	CODE_LESS_THAN        = "less-than"
	CODE_EMAIL            = "email"
	CODE_HAS_PREFIX       = "has-prefix"
	CODE_HAS_SUFFIX       = "has-suffix"
	CODE_NOT_HAS_PREFIX   = "not-has-prefix"
	CODE_NOT_HAS_SUFFIX   = "not-has-suffix"
	CODE_EQUALS           = "equals"
	CODE_ENUM             = "enum"
	CODE_URL              = "url"
	CODE_REGEXP           = "regexp"
	CODE_UUID             = "uuid"
	CODE_PASSWORD         = "password"
	CODE_JSON             = "json"
	CODE_DATE_TIME        = "date-time"
	CODE_DATE             = "date"
	CODE_TIME             = "time"
	CODE_DATE_EQUAL       = "date-equal"  // TODO: merge with `Equals`
	CODE_DATE_BEFORE      = "date-before" // TODO: merge with `LessThan`
	CODE_DATE_AFTER       = "date-after"  // TODO: merge with `GreaterThan`
	CODE_LATITUDE         = "latitude"
	CODE_LONGITUDE        = "longitude"
)
