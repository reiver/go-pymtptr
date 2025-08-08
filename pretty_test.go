package pymtptr_test

import (
	"testing"

	"github.com/reiver/go-pymtptr"
)

func TestPretty(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
		},



		{
			Value:    "$example.com",
			Expected: "$example.com",
		},
		{
			Value:    "$EXAMPLE.COM",
			Expected: "$example.com",
		},



		{
			Value:    "$example.com/Path",
			Expected: "$example.com/Path",
		},
		{
			Value:    "$EXAMPLE.COM/Path",
			Expected: "$example.com/Path",
		},



		{
			Value:    "$example.com/Once/Twice/Thrice/Fource/",
			Expected: "$example.com/Once/Twice/Thrice/Fource/",
		},
		{
			Value:    "$EXAMPLE.COM/Once/Twice/Thrice/Fource/",
			Expected: "$example.com/Once/Twice/Thrice/Fource/",
		},



		{
			Value:    "$example.com/۱/۲/۳/۴/۵",
			Expected: "$example.com/۱/۲/۳/۴/۵",
		},
		{
			Value:    "$EXAMPLE.COM/۱/۲/۳/۴/۵",
			Expected: "$example.com/۱/۲/۳/۴/۵",
		},



		{
			Value:    "$example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "$example.com/۱/۲/۳/۴/۵",
		},
		{
			Value:    "$EXAMPLE.COM/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "$example.com/۱/۲/۳/۴/۵",
		},



		{
			Value:    "$example.com/~JoeBlow",
			Expected: "$example.com/~JoeBlow",
		},
		{
			Value:    "$EXAMPLE.COM/~JoeBlow",
			Expected: "$example.com/~JoeBlow",
		},



		{
			Value:    "$apple.BANANA.Cherry",
			Expected: "$apple.banana.cherry",
		},
		{
			Value:    "$apple.BANANA.Cherry/Path",
			Expected: "$apple.banana.cherry/Path",
		},



		{
			Value:    "$😈.com",
			Expected: "$😈.com",
		},
		{
			Value:    "$😈.com/Path",
			Expected: "$😈.com/Path",
		},



		{
			Value:    "$xn--m28h.com",
			Expected: "$😈.com",
		},
		{
			Value:    "$xn--m28h.com/Path",
			Expected: "$😈.com/Path",
		},



		{
			Value:    "$XN--M28H.com",
			Expected: "$😈.com",
		},
		{
			Value:    "$XN--M28H.com/Path",
			Expected: "$😈.com/Path",
		},
	}

	for testNumber, test := range tests {

		actual := pymtptr.Pretty(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual pretty payment-pointer is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
