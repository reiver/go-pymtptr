package pymtptr

import (
	"net/url"

	"golang.org/x/net/idna"
)

// PaymentPointer represents a Payment-Pointer, as defined by:
// https://paymentpointers.org/ and https://paymentpointers.org/syntax/
type PaymentPointer struct {
	Host string
	Path string
}

// Pretty returns the "pretty" version of a payment-pointer.
//
// NOTE that the the "pretty" version of a payment-pointer
// might NOT be a valid payment-pointer!
//
// What pretty returns is meant to be a more human-legible form
// meant for displaying to people (and is NOT mean to be consumed
// by machines).
//
// Some example "pretty" payment-pointer are:
//
//	$example.com
//
//	$joeblow.example.com
//
//	$example.com/joeblow
//
//	$🙂.com/۱/۲/۳/۴/۵
//
// Note that the last example of a "pretty" version of a payment-pointer
// (i.e., "$🙂.com/۱/۲/۳/۴/۵") is NOT a valid payment-pointer!
//
// To get a valid payment-pointer, instead use [PaymentPointer.String].
func (receiver PaymentPointer) Pretty() string {
	if "" == receiver.Host {
		return ""
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, '$')

	{
		host := canonicalHost(receiver.Host)

		var err error
		host, err = idna.ToUnicode(host)
		if nil != err {
			return ""
		}

		p = append(p, host...)
	}

	if "" != receiver.Path {
		if '/' != receiver.Path[0] {
			p = append(p, '/')
		}

		p = append(p, receiver.Path...)
	}

	return string(p)
}

// String returns the payment-pointer.
//
// An example payment-pointer looks like these:
//
//	$example.com
//
//	$joeblow.example.com
//
//	$example.com/joeblow
//
//	$xn--938h.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
//
// Note that what String returns is in ASCII.
//
// A host with non-ASCII characters is encoded into ASCII using
// punycode.
//
// A path with non-ASCII character is encoded into ASCII using
// urlencode.
func (receiver PaymentPointer) String() string {
	if "" == receiver.Host {
		return ""
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, '$')

	{
		host, err := idna.ToASCII(receiver.Host)
		if nil != err {
			return ""
		}

		var length int = len(host)
		for i:=0; i<length; i++ {
			var b byte = host[i]
			switch {
			case 'A' <= b && b <= 'Z':
				b += ('a' - 'A')
				p = append(p, b)
			default:
				p = append(p, b)
			}
		}
	}

	if "" != receiver.Path {
		if '/' != receiver.Path[0] {
			p = append(p, '/')
		}

		var u = url.URL{
			Path: receiver.Path,
		}
		p = append(p, u.EscapedPath()...)
	}

	return string(p)
}
