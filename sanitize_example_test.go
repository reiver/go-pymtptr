package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExampleSanitize_nonAscii() {

	var paymentPointer string = "$hello-world🙂.com/۱/۲/۳/۴/۵"

	sanitized := pymtptr.Sanitize(paymentPointer)

	fmt.Println(sanitized)

	// Output:
	// $xn--hello-world-lk27j.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
}

func ExampleSanitize_alreadyEncoded() {

	var paymentPointer string = "$xn--hello-world-lk27j.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5"

	sanitized := pymtptr.Sanitize(paymentPointer)

	fmt.Println(sanitized)

	// Output:
	// $xn--hello-world-lk27j.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
}

func ExampleSanitize_justHost() {

	var paymentPointer string = "$hello-world🙂.com"

	sanitized := pymtptr.Sanitize(paymentPointer)

	fmt.Println(sanitized)

	// Output:
	// $xn--hello-world-lk27j.com
}

func ExampleSanitize_upperCaseHost() {

	var paymentPointer string = "$HELLO-WORLD🙂.com"

	sanitized := pymtptr.Sanitize(paymentPointer)

	fmt.Println(sanitized)

	// Output:
	// $xn--hello-world-lk27j.com
}
