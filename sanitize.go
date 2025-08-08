package pymtptr

// Sanitize returns the sanitized form of a Payment-Pointer.
func Sanitize(value string) string {
	var pp PaymentPointer
	err := pp.Parse(value)
	if nil != err {
		return ""
	}

	return pp.String()
}
