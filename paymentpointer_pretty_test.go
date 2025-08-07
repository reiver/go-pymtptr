package pymtptr_test

import (
	"testing"

	"github.com/reiver/go-pymtptr"
)

func TestPaymentPointer_Pretty(t *testing.T) {

	tests := []struct{
		Host string
		Path string
		Expected string
	}{
		{
		},



		{
			Host:      "example.com",
			Expected: "$example.com",
		},
		{
			Host:      "EXAMPLE.COM",
			Expected: "$example.com",
		},



		{
			Host:      "example.com",
			Path:                  "Path",
			Expected: "$example.com/Path",
		},
		{
			Host:      "EXAMPLE.COM",
			Path:                  "Path",
			Expected: "$example.com/Path",
		},



		{
			Host:      "example.com",
			Path:                 "/Path",
			Expected: "$example.com/Path",
		},
		{
			Host:      "EXAMPLE.COM",
			Path:                 "/Path",
			Expected: "$example.com/Path",
		},



		{
			Host:      "example.com",
			Path:                  "Once/Twice/Thrice/Fource/",
			Expected: "$example.com/Once/Twice/Thrice/Fource/",
		},
		{
			Host:      "EXAMPLE.COM",
			Path:                  "Once/Twice/Thrice/Fource/",
			Expected: "$example.com/Once/Twice/Thrice/Fource/",
		},



		{
			Host:      "example.com",
			Path:                  "۱/۲/۳/۴/۵",
			Expected: "$example.com/۱/۲/۳/۴/۵",
		},
		{
			Host:      "EXAMPLE.COM",
			Path:                  "۱/۲/۳/۴/۵",
			Expected: "$example.com/۱/۲/۳/۴/۵",
		},



		{
			Host:      "example.com",
			Path:                  "%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "$example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},
		{
			Host:      "EXAMPLE.COM",
			Path:                  "%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "$example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},



		{
			Host:      "example.com",
			Path:                  "~JoeBlow",
			Expected: "$example.com/~JoeBlow",
		},
		{
			Host:      "EXAMPLE.COM",
			Path:                  "~JoeBlow",
			Expected: "$example.com/~JoeBlow",
		},



		{
			Host:      "apple.BANANA.Cherry",
			Expected: "$apple.banana.cherry",
		},
		{
			Host:      "apple.BANANA.Cherry",
			Path:                          "Path",
			Expected: "$apple.banana.cherry/Path",
		},
		{
			Host:      "apple.BANANA.Cherry",
			Path:                         "/Path",
			Expected: "$apple.banana.cherry/Path",
		},



		{
			Host:      "😈.com",
			Expected: "$😈.com",
		},
		{
			Host:      "😈.com",
			Path:                   "Path",
			Expected: "$😈.com/Path",
		},
		{
			Host:      "😈.com",
			Path:                  "/Path",
			Expected: "$😈.com/Path",
		},



		{
			Host:      "xn--m28h.com",
			Expected: "$😈.com",
		},
		{
			Host:      "xn--m28h.com",
			Path:                   "Path",
			Expected: "$😈.com/Path",
		},
		{
			Host:      "xn--m28h.com",
			Path:                  "/Path",
			Expected: "$😈.com/Path",
		},



		{
			Host:      "XN--M28H.com",
			Expected: "$😈.com",
		},
		{
			Host:      "XN--M28H.com",
			Path:                   "Path",
			Expected: "$😈.com/Path",
		},
		{
			Host:      "XN--M28H.com",
			Path:                  "/Path",
			Expected: "$😈.com/Path",
		},
	}

	for testNumber, test := range tests {

		actual := pymtptr.PaymentPointer{
			Host:test.Host,
			Path:test.Path,
		}.Pretty()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual pretty payment-pointer is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("HOST: %q", test.Host)
			t.Logf("PATH: %q", test.Path)
			continue
		}
	}
}
