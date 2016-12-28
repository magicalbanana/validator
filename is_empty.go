package validator

// IsEmpty check string's length (including multi byte strings)
func IsEmpty(str string) bool {
	if len(str) > 0 {
		return false
	}

	return true
}
