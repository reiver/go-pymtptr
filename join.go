package pymtptr

// Join returns the payment-pointer.
//
// An example payment-pointer looks like these:
//
//      $example.com
//
//      $joeblow.example.com
//
//      $example.com/joeblow
//
//      $xn--938h.com/%DB%B1/%DB%B2/%DB%B3/%DB%B4/%DB%B5
//
// Note that what Join returns is in ASCII.
//
// A host with non-ASCII characters is encoded into ASCII using
// punycode.
//
// A path with non-ASCII character is encoded into ASCII using
// urlencode.
func Join(host string, path string) (paymentPointer string) {
	if "" == host {
		return ""
	}

	var pp = PaymentPointer{
		Host:host,
		Path:path,
	}

	return pp.String()
}
