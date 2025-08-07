package pymtptr

// Split parses a Payment-Pointer, and if valid, extracts the host and path
// from the Payment-Pointer and returns them.
func Split(paymentPointer string) (host string, path string, err error) {
	var pp PaymentPointer
	err = pp.Parse(paymentPointer)
	if nil != err {
		return
	}

	return pp.Host, pp.Path, nil
}
