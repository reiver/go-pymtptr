package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExamplePretty_onlyHost() {

	var paymentPointer string = "$example.com"

	pretty := pymtptr.Pretty(paymentPointer)

	fmt.Println(pretty)

	// Output:
	// $example.com
}

func ExamplePretty_withPath() {

	var paymentPointer string = "$example.com/joeblow"

	pretty := pymtptr.Pretty(paymentPointer)

	fmt.Println(pretty)

	// Output:
	// $example.com/joeblow
}

func ExamplePretty_withSubDomain() {

	var paymentPointer string = "$joeblow.example.com"

	pretty := pymtptr.Pretty(paymentPointer)

	fmt.Println(pretty)

	// Output:
	// $joeblow.example.com
}

func ExamplePretty_nonAsciiHost() {

	var paymentPointer string = "$😈.com"

	pretty := pymtptr.Pretty(paymentPointer)

	fmt.Println(pretty)

	// Output:
	// $😈.com
}

func ExamplePretty_idnaHost() {

	var paymentPointer string = "$xn--m28h.com"

	pretty := pymtptr.Pretty(paymentPointer)

	fmt.Println(pretty)

	// Output:
	// $😈.com
}

func ExamplePretty_nonAsciiPath() {

	var paymentPointer string = "$example.com/۱/۲/۳/۴/۵"

	pretty := pymtptr.Pretty(paymentPointer)

	fmt.Println(pretty)

	// Output:
	// $example.com/۱/۲/۳/۴/۵
}
