package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExamplePaymentPointer_String_onlyHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
	}

	str := paymentPointer.String()

	fmt.Println(str)

	// Output:
	// $example.com
}

func ExamplePaymentPointer_String_withPath() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
		Path: "/joeblow",
	}

	str := paymentPointer.String()

	fmt.Println(str)

	// Output:
	// $example.com/joeblow
}

func ExamplePaymentPointer_String_withSubDomain() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "joeblow.example.com",
	}

	str := paymentPointer.String()

	fmt.Println(str)

	// Output:
	// $joeblow.example.com
}

func ExamplePaymentPointer_String_nonAsciiHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "😈.com",
	}

	str := paymentPointer.String()

	fmt.Println(str)

	// Output:
	// $xn--m28h.com
}

func ExamplePaymentPointer_String_idnaHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "xn--m28h.com",
	}

	str := paymentPointer.String()

	fmt.Println(str)

	// Output:
	// $xn--m28h.com
}

func ExamplePaymentPointer_String_nonAsciiPath() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
		Path: "/۱/۲/۳/۴/۵",
	}

	str := paymentPointer.String()

	fmt.Println(str)

	// Output:
	// $example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
}
