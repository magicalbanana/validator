package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStringNotInteger(t *testing.T) {
	t.Parallel()

	tests := []struct {
		param    string
		expected bool
	}{
		{"0", false},
		{"1", false},
		{"100", false},
		{"3256", false},
		{"03256", true},
		{"127.0.0.1", true},
		{"1,200", true},
		{"300.00", true},
		{"", true},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", true},
	}

	for _, test := range tests {
		actual := IsStringNotInteger(test.param)
		assert.Equal(t, test.expected, actual, test.param)
	}
}
