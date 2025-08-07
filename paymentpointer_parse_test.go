package pymtptr_test

import (
	"testing"

	"github.com/reiver/go-pymtptr"
)

func TestPaymentPointer_Parse(t *testing.T) {

	tests := []struct{
		Value string
		ExpectedHost string
		ExpectedPath string
	}{
		{
			Value: "$",
		},



		{
			Value:       "$example.com",
			ExpectedHost: "example.com",
		},
		{
			Value:       "$EXAMPLE.COM",
			ExpectedHost: "EXAMPLE.COM",
		},



		{
			Value:       "$example.com/JaneDoe",
			ExpectedHost: "example.com",
			ExpectedPath:            "/JaneDoe",
		},
		{
			Value:       "$EXAMPLE.COM/JaneDoe",
			ExpectedHost: "EXAMPLE.COM",
			ExpectedPath:            "/JaneDoe",
		},



		{
			Value:       "$joeblow.example.com",
			ExpectedHost: "joeblow.example.com",
		},
		{
			Value:       "$JOEBLOW.EXAMPLE.COM",
			ExpectedHost: "JOEBLOW.EXAMPLE.COM",
		},



		{
			Value:       "$xn--938h.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
			ExpectedHost: "xn--938h.com",
			ExpectedPath:             "/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5",
		},
	}

	for testNumber, test := range tests  {

		var paymentPointer pymtptr.PaymentPointer
		err := paymentPointer.Parse(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %q", err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			actual := paymentPointer.Host
			expected := test.ExpectedHost

			if expected != actual {
				t.Errorf("For test #%d, the actual host is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			actual := paymentPointer.Path
			expected := test.ExpectedPath

			if expected != actual {
				t.Errorf("For test #%d, the actual path is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
