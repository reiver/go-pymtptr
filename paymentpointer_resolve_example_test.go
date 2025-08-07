package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExamplePaymentPointer_Resolve_onlyHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
	}

	url := paymentPointer.Resolve()

	fmt.Println(url)

	// Output:
	// https://example.com/.well-known/pay
}

func ExamplePaymentPointer_Resolve_withPath() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
		Path: "/joeblow",
	}

	url := paymentPointer.Resolve()

	fmt.Println(url)

	// Output:
	// https://example.com/joeblow
}

func ExamplePaymentPointer_Resolve_withSubDomain() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "joeblow.example.com",
	}

	url := paymentPointer.Resolve()

	fmt.Println(url)

	// Output:
	// https://joeblow.example.com/.well-known/pay
}

func ExamplePaymentPointer_Resolve_nonAsciiHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "😈.com",
	}

	url := paymentPointer.Resolve()

	fmt.Println(url)

	// Output:
	// https://xn--m28h.com/.well-known/pay
}

func ExamplePaymentPointer_Resolve_idnaHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "xn--m28h.com",
	}

	url := paymentPointer.Resolve()

	fmt.Println(url)

	// Output:
	// https://xn--m28h.com/.well-known/pay
}

func ExamplePaymentPointer_Resolve_nonAsciiPath() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
		Path: "/۱/۲/۳/۴/۵",
	}

	url := paymentPointer.Resolve()

	fmt.Println(url)

	// Output:
	// https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
}
