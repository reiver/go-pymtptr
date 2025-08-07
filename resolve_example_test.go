package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExampleResolve_onlyHost() {

	paymentPointer := "$example.com"

	url := pymtptr.Resolve(paymentPointer)

	fmt.Println(url)

	// Output:
	// https://example.com/.well-known/pay
}

func ExampleResolve_withPath() {

	paymentPointer := "$example.com/joeblow"

	url := pymtptr.Resolve(paymentPointer)

	fmt.Println(url)

	// Output:
	// https://example.com/joeblow
}

func ExampleResolve_withSubDomain() {

	paymentPointer := "$joeblow.example.com"

	url := pymtptr.Resolve(paymentPointer)

	fmt.Println(url)

	// Output:
	// https://joeblow.example.com/.well-known/pay
}

func ExampleResolve_nonAsciiHost() {

	paymentPointer := "$😈.com"

	url := pymtptr.Resolve(paymentPointer)

	fmt.Println(url)

	// Output:
	// https://xn--m28h.com/.well-known/pay
}

func ExampleResolve_idnaHost() {

	paymentPointer := "$xn--m28h.com"

	url := pymtptr.Resolve(paymentPointer)

	fmt.Println(url)

	// Output:
	// https://xn--m28h.com/.well-known/pay
}

func ExampleResolve_nonAsciiPath() {

	paymentPointer := "$example.com/۱/۲/۳/۴/۵"

	url := pymtptr.Resolve(paymentPointer)

	fmt.Println(url)

	// Output:
	// https://example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
}
