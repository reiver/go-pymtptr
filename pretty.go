package pymtptr

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
func Pretty(paymentPointer string) string {
	var pp PaymentPointer
	err := pp.Parse(paymentPointer)
	if nil != err {
		return ""
	}

	return pp.Pretty()
}
