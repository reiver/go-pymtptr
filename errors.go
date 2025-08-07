package pymtptr

import (
	"github.com/reiver/go-erorr"
)

const (
	errDollarSignPrefixNotFound = erorr.Error("pymtptr: dollar-sign ($) prefix not found")
	errEmptyString              = erorr.Error("pymtptr: empty string")
	errNilReceiver              = erorr.Error("pymtptr: nil receiver")
)
