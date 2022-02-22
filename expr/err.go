package expr

import "errors"

var (
	ErrTooFewValues     = errors.New("too few values specified")
	ErrTooManyValues    = errors.New("too many values specified")
	ErrUnknownOperation = errors.New("unknown operation supplied")
)
