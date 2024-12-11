package id

// luhnValidator is a Validator implementation that uses the Luhn algorithm for signing and verifying numbers.
type luhnValidator struct{}

// Sign generates a Luhn checksum and appends it to the given number.
func (v *luhnValidator) Sign(number int64) int64 {
	sign := checksum(number)
	if sign != 0 {
		sign = 10 - sign
	}
	return (number * 10) + sign
}

// Verify checks if the given number is valid according to the Luhn algorithm.
func (v *luhnValidator) Verify(number int64) bool {
	return (number%10+checksum(number/10))%10 == 0
}

// LuhnValidator creates a new Validator that uses the Luhn algorithm.
func LuhnValidator() Validator {
	return &luhnValidator{}
}

// checksum calculates the Luhn checksum for the given number.
func checksum(number int64) int64 {
	var luhn int64
	double := false

	for number > 0 {
		cur := number % 10
		number /= 10

		if double {
			cur *= 2
			if cur > 9 {
				cur -= 9
			}
		}

		luhn += cur
		double = !double
	}

	return luhn % 10
}
