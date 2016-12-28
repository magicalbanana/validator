package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{``, false},
		{`foo@bar.com`, true},
		{`x@x.x`, true},
		{`foo@bar.com.au`, true},
		{`foo+bar@bar.com`, true},
		{`foo@bar.coffee`, true},
		{`foo@bar.中文网`, true},
		{`invalidemail@`, false},
		{`invalid.com`, false},
		{`@invalid.com`, false},
		{`test|123@m端ller.com`, true},
		{`hans@m端ller.com`, true},
		{`hans.m端ller@test.com`, true},
		{`foo.Babbie@DomaIn.cOM`, true},
		{`manbearpig@DOMAIN.CO.UK`, true},
		{`very.(),:;<>[]".VERY."very@\ "very".unusual@strange.example.com`, true},
	}
	for _, test := range tests {
		actual := IsEmail(test.param)
		assert.Equal(t, test.expected, actual, test.param)
	}
}
