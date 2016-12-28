package validator

import "regexp"

// IsStringNotInteger returns true if the given value does not match the regex
// pattern used and returns false if it matches.
func IsStringNotInteger(value string) bool {
	regx := regexp.MustCompile(`^([1-9]\d*|0)$`)
	return !regx.MatchString(value)
}
