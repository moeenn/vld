package vld

const (
	PATTERN_EMAIL = `^[^@]+@[^@]+\.[^@]+$`
	PATTERN_UUID  = `^[a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}$`

	// Password strength rules:
	// - Minimum eight characters
	// - At least one uppercase letter
	// - One lowercase letter
	// - One number
	// - One special character
	// **Note**: the pattern will match weak passwords instead of strong passwords
	PATTERN_PASSWORD_STRENGTH = `^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`
)
