package pymtptr_test

import (
	"testing"

	"github.com/reiver/go-pymtptr"
)

func TestResolve(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
		},



		{
			Value:           "$example.com",
			Expected: "https://example.com/.well-known/pay",
		},
		{
			Value:           "$EXAMPLE.COM",
			Expected: "https://example.com/.well-known/pay",
		},



		{
			Value:           "$example.com/Path",
			Expected: "https://example.com/Path",
		},
		{
			Value:           "$EXAMPLE.COM/Path",
			Expected: "https://example.com/Path",
		},



		{
			Value:           "$example.com/Once/Twice/Thrice/Fource/",
			Expected: "https://example.com/Once/Twice/Thrice/Fource/",
		},
		{
			Value:           "$EXAMPLE.COM/Once/Twice/Thrice/Fource/",
			Expected: "https://example.com/Once/Twice/Thrice/Fource/",
		},



		{
			Value:           "$example.com/۱/۲/۳/۴/۵",
			Expected: "https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},
		{
			Value:           "$EXAMPLE.COM/۱/۲/۳/۴/۵",
			Expected: "https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},



		{
			Value:           "$example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},
		{
			Value:           "$EXAMPLE.COM/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},



		{
			Value:           "$example.com/~JoeBlow",
			Expected: "https://example.com/~JoeBlow",
		},
		{
			Value:           "$EXAMPLE.COM/~JoeBlow",
			Expected: "https://example.com/~JoeBlow",
		},



		{
			Value:           "$apple.BANANA.Cherry",
			Expected: "https://apple.banana.cherry/.well-known/pay",
		},
		{
			Value:           "$apple.BANANA.Cherry/Path",
			Expected: "https://apple.banana.cherry/Path",
		},



		{
			Value:           "$😈.com",
			Expected: "https://xn--m28h.com/.well-known/pay",
		},
		{
			Value:           "$😈.com/Path",
			Expected: "https://xn--m28h.com/Path",
		},



		{
			Value:           "$xn--m28h.com",
			Expected: "https://xn--m28h.com/.well-known/pay",
		},
		{
			Value:           "$xn--m28h.com/Path",
			Expected: "https://xn--m28h.com/Path",
		},



		{
			Value:           "$XN--M28H.com",
			Expected: "https://xn--m28h.com/.well-known/pay",
		},
		{
			Value:           "$XN--M28H.com/Path",
			Expected: "https://xn--m28h.com/Path",
		},



		{
			Value:           "$joeblow.example.com",
			Expected: "https://joeblow.example.com/.well-known/pay",
		},
		{
			Value:           "$JoeBlow.EXAMPLE.COM",
			Expected: "https://joeblow.example.com/.well-known/pay",
		},
		{
			Value:           "$JoeBlow.EXAMPLE.COM/pay-me",
			Expected: "https://joeblow.example.com/pay-me",
		},
	}

	for testNumber, test := range tests {

		actual := pymtptr.Resolve(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the resolved payment-pointer HTTP-URL is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
