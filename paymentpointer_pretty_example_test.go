package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExamplePaymentPointer_Pretty_onlyHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
	}

	str := paymentPointer.Pretty()

	fmt.Println(str)

	// Output:
	// $example.com
}

func ExamplePaymentPointer_Pretty_withPath() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
		Path: "/joeblow",
	}

	pretty := paymentPointer.Pretty()

	fmt.Println(pretty)

	// Output:
	// $example.com/joeblow
}

func ExamplePaymentPointer_Pretty_withSubDomain() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "joeblow.example.com",
	}

	pretty := paymentPointer.Pretty()

	fmt.Println(pretty)

	// Output:
	// $joeblow.example.com
}

func ExamplePaymentPointer_Pretty_nonAsciiHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "😈.com",
	}

	pretty := paymentPointer.Pretty()

	fmt.Println(pretty)

	// Output:
	// $😈.com
}

func ExamplePaymentPointer_Pretty_idnaHost() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "xn--m28h.com",
	}

	pretty := paymentPointer.Pretty()

	fmt.Println(pretty)

	// Output:
	// $😈.com
}

func ExamplePaymentPointer_Pretty_nonAsciiPath() {

	var paymentPointer = pymtptr.PaymentPointer{
		Host: "example.com",
		Path: "/۱/۲/۳/۴/۵",
	}

	pretty := paymentPointer.Pretty()

	fmt.Println(pretty)

	// Output:
	// $example.com/۱/۲/۳/۴/۵
}
