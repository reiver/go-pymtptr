package pymtptr_test

import (
	"fmt"

	"github.com/reiver/go-pymtptr"
)

func ExampleSplit_onlyHost() {

	var value string = "$example.com"

	host, path, err := pymtptr.Split(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", host)
	fmt.Printf("PATH: %q\n", path)

	// Output:
	// HOST: "example.com"
	// PATH: ""
}

func ExampleSplit_withPath() {

	var value string = "$example.com/joeblow"

	host, path, err := pymtptr.Split(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", host)
	fmt.Printf("PATH: %q\n", path)

	// Output:
	// HOST: "example.com"
	// PATH: "/joeblow"
}

func ExampleSplit_withSubDomain() {

	var value string = "$joeblow.example.com"

	host, path, err := pymtptr.Split(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", host)
	fmt.Printf("PATH: %q\n", path)

	// Output:
	// HOST: "joeblow.example.com"
	// PATH: ""
}

func ExampleSplit_nonAsciiHost() {

	var value string = "$😈.com"

	host, path, err := pymtptr.Split(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", host)
	fmt.Printf("PATH: %q\n", path)

	// Output:
	// HOST: "😈.com"
	// PATH: ""
}

func ExampleSplit_idnaHost() {

	var value string = "$xn--m28h.com"

	host, path, err := pymtptr.Split(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", host)
	fmt.Printf("PATH: %q\n", path)

	// Output:
	// HOST: "😈.com"
	// PATH: ""
}

func ExampleSplit_nonAsciiPath() {

	var value string = "$example.com/۱/۲/۳/۴/۵"

	host, path, err := pymtptr.Split(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("HOST: %q\n", host)
	fmt.Printf("PATH: %q\n", path)

	// Output:
	// HOST: "example.com"
	// PATH: "/۱/۲/۳/۴/۵"
}
