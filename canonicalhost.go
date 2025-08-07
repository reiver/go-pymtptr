package pymtptr

func canonicalHost(host string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, host...)

	var length int = len(p)

	for i:=0; i<length; i++ {
		var b byte = p[i]

		switch {
		case 'A' <= b && b <= 'Z':
			p[i] += ('a' - 'A')
		}
	}

	return string(p)
}
