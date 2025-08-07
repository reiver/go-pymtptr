package pymtptr_test

import (
	"testing"

	"github.com/reiver/go-pymtptr"
)

func TestPaymentPointer_Resolve(t *testing.T) {

	tests := []struct{
		Host string
		Path string
		Expected string
	}{
		{
		},



		{
			Host:             "example.com",
			Expected: "https://example.com/.well-known/pay",
		},
		{
			Host:             "EXAMPLE.COM",
			Expected: "https://example.com/.well-known/pay",
		},



		{
			Host:             "example.com",
			Path:                         "Path",
			Expected: "https://example.com/Path",
		},
		{
			Host:             "EXAMPLE.COM",
			Path:                         "Path",
			Expected: "https://example.com/Path",
		},



		{
			Host:             "example.com",
			Path:                        "/Path",
			Expected: "https://example.com/Path",
		},
		{
			Host:             "EXAMPLE.COM",
			Path:                        "/Path",
			Expected: "https://example.com/Path",
		},



		{
			Host:             "example.com",
			Path:                         "Once/Twice/Thrice/Fource/",
			Expected: "https://example.com/Once/Twice/Thrice/Fource/",
		},
		{
			Host:             "EXAMPLE.COM",
			Path:                         "Once/Twice/Thrice/Fource/",
			Expected: "https://example.com/Once/Twice/Thrice/Fource/",
		},



		{
			Host:             "example.com",
			Path:                         "۱/۲/۳/۴/۵",
			Expected: "https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},
		{
			Host:             "EXAMPLE.COM",
			Path:                         "۱/۲/۳/۴/۵",
			Expected: "https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},



		{
			Host:             "example.com",
			Path:                         "%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "https://example.com/%25DB%25B1/%25DB%25B2/%25DB%25B3/%25DB%25B4/%25DB%25B5",
		},
		{
			Host:             "EXAMPLE.COM",
			Path:                         "%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			Expected: "https://example.com/%25DB%25B1/%25DB%25B2/%25DB%25B3/%25DB%25B4/%25DB%25B5",
		},



		{
			Host:              "example.com",
			Path:                         "~JoeBlow",
			Expected: "https://example.com/~JoeBlow",
		},
		{
			Host:             "EXAMPLE.COM",
			Path:                         "~JoeBlow",
			Expected: "https://example.com/~JoeBlow",
		},



		{
			Host:             "apple.BANANA.Cherry",
			Expected: "https://apple.banana.cherry/.well-known/pay",
		},
		{
			Host:             "apple.BANANA.Cherry",
			Path:                                 "Path",
			Expected: "https://apple.banana.cherry/Path",
		},
		{
			Host:             "apple.BANANA.Cherry",
			Path:                                "/Path",
			Expected: "https://apple.banana.cherry/Path",
		},



		{
			Host:             "😈.com",
			Expected: "https://xn--m28h.com/.well-known/pay",
		},
		{
			Host:             "😈.com",
			Path:                          "Path",
			Expected: "https://xn--m28h.com/Path",
		},
		{
			Host:             "😈.com",
			Path:                         "/Path",
			Expected: "https://xn--m28h.com/Path",
		},



		{
			Host:             "xn--m28h.com",
			Expected: "https://xn--m28h.com/.well-known/pay",
		},
		{
			Host:             "xn--m28h.com",
			Path:                          "Path",
			Expected: "https://xn--m28h.com/Path",
		},
		{
			Host:             "xn--m28h.com",
			Path:                         "/Path",
			Expected: "https://xn--m28h.com/Path",
		},



		{
			Host:             "XN--M28H.com",
			Expected: "https://xn--m28h.com/.well-known/pay",
		},
		{
			Host:             "XN--M28H.com",
			Path:                          "Path",
			Expected: "https://xn--m28h.com/Path",
		},
		{
			Host:             "XN--M28H.com",
			Path:                         "/Path",
			Expected: "https://xn--m28h.com/Path",
		},



		{
			Host:             "joeblow.example.com",
			Expected: "https://joeblow.example.com/.well-known/pay",
		},
		{
			Host:             "JoeBlow.EXAMPLE.COM",
			Expected: "https://joeblow.example.com/.well-known/pay",
		},
	}

	for testNumber, test := range tests {

		actual := pymtptr.PaymentPointer{
			Host:test.Host,
			Path:test.Path,
		}.Resolve()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the resolved payment-pointer HTTP-URL is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("HOST: %q", test.Host)
			t.Logf("PATH: %q", test.Path)
			continue
		}
	}
}
