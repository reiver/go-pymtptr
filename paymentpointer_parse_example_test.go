package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExamplePaymentPointer_Parse_onlyHost() {

	var value string = "$example.com"

	var paymentPointer pymtptr.PaymentPointer

	err := paymentPointer.Parse(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", paymentPointer.Host)
	fmt.Printf("PATH: %q\n", paymentPointer.Path)

	// Output:
	// HOST: "example.com"
	// PATH: ""
}

func ExamplePaymentPointer_Parse_withPath() {

	var value string = "$example.com/joeblow"

	var paymentPointer pymtptr.PaymentPointer

	err := paymentPointer.Parse(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", paymentPointer.Host)
	fmt.Printf("PATH: %q\n", paymentPointer.Path)

	// Output:
	// HOST: "example.com"
	// PATH: "/joeblow"
}

func ExamplePaymentPointer_Parse_withSubDomain() {

	var value string = "$joeblow.example.com"

	var paymentPointer pymtptr.PaymentPointer

	err := paymentPointer.Parse(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", paymentPointer.Host)
	fmt.Printf("PATH: %q\n", paymentPointer.Path)

	// Output:
	// HOST: "joeblow.example.com"
	// PATH: ""
}

func ExamplePaymentPointer_Parse_nonAsciiHost() {

	var value string = "$😈.com"

	var paymentPointer pymtptr.PaymentPointer

	err := paymentPointer.Parse(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", paymentPointer.Host)
	fmt.Printf("PATH: %q\n", paymentPointer.Path)

	// Output:
	// HOST: "😈.com"
	// PATH: ""
}

func ExamplePaymentPointer_Parse_idnaHost() {

	var value string = "$xn--m28h.com"

	var paymentPointer pymtptr.PaymentPointer

	err := paymentPointer.Parse(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", paymentPointer.Host)
	fmt.Printf("PATH: %q\n", paymentPointer.Path)

	// Output:
	// HOST: "xn--m28h.com"
	// PATH: ""
}

func ExamplePaymentPointer_Parse_nonAsciiPath() {

	var value string = "$example.com/۱/۲/۳/۴/۵"

	var paymentPointer pymtptr.PaymentPointer

	err := paymentPointer.Parse(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", paymentPointer.Host)
	fmt.Printf("PATH: %q\n", paymentPointer.Path)

	// Output:
	// HOST: "example.com"
	// PATH: "/۱/۲/۳/۴/۵"
}
