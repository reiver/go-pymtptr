package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExampleJoin_onlyHost() {

	var host string = "example.com"
	var path string = ""

	str := pymtptr.Join(host, path)

	fmt.Println(str)

	// Output:
	// $example.com
}

func ExampleJoin_withPath() {

	var host string = "example.com"
	var path string = "/joeblow"

	str := pymtptr.Join(host, path)

	fmt.Println(str)

	// Output:
	// $example.com/joeblow
}

func ExampleJoin_withSubDomain() {

	var host string = "joeblow.example.com"
	var path string = ""

	str := pymtptr.Join(host, path)

	fmt.Println(str)

	// Output:
	// $joeblow.example.com
}

func ExampleJoin_nonAsciiHost() {

	var host string = "😈.com"
	var path string = ""

	str := pymtptr.Join(host, path)

	fmt.Println(str)

	// Output:
	// $xn--m28h.com
}

func ExampleJoin_idnaHost() {

	var host string = "xn--m28h.com"
	var path string = ""

	str := pymtptr.Join(host, path)

	fmt.Println(str)

	// Output:
	// $xn--m28h.com
}

func ExampleJoin_nonAsciiPath() {

	var host string = "example.com"
	var path string = "/۱/۲/۳/۴/۵"

	str := pymtptr.Join(host, path)

	fmt.Println(str)

	// Output:
	// $example.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
}
