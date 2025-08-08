package pymtptr

// Resolve resolves a Payment-Pointer to an HTTP URL.
//
// If the Payment-Pointer is NOT a valid the Resolve returns an empty string.
func Resolve(paymentPointer string) (httpsURL string) {
	var pp PaymentPointer
	err := pp.Parse(paymentPointer)
	if nil != err {
		return ""
	}

	return pp.Resolve()
}
