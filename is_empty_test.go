package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"1", false},
		{"", true},
	}
	for _, test := range tests {
		actual := IsEmpty(test.param)
		assert.Equal(t, test.expected, actual, test.param)
	}
}
