package validator

import "strings"

// IsEmail is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
// https://en.wikipedia.org/wiki/Email_address#Valid_email_addresses
func IsEmail(s string) bool {
	if !strings.Contains(s, "@") || string(s[0]) == "@" || string(s[len(s)-1]) == "@" {
		return false
	}
	return true
}
