package id

// noValidator is a Validator implementation that performs no validation.
type noValidator struct{}

// Sign returns the input number unchanged.
func (v *noValidator) Sign(number int64) int64 {
	return number
}

// Verify always returns true, indicating any number is valid.
func (v *noValidator) Verify(number int64) bool {
	return true
}

// NoValidator creates a new Validator that performs no validation.
func NoValidator() Validator {
	return &noValidator{}
}
