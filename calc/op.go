package calc

import "errors"

var (
	ErrTooFewValues     = errors.New("too few values specified")
	ErrTooManyValues    = errors.New("too many values specified")
	ErrUnknownOperation = errors.New("unknown operation supplied")
)

type Operation int

const (
	_ Operation = iota
	OpAdd
	OpSub
	OpMul
	OpDiv
)

func (o *Operation) Capture(s []string) error {
	if len(s) < 1 {
		return ErrTooFewValues
	}
	if len(s) > 1 {
		return ErrTooManyValues
	}

	switch s[0] {
	case "+":
		*o = OpAdd
	case "-":
		*o = OpSub
	case "*":
		*o = OpMul
	case "/":
		*o = OpDiv
	default:
		return ErrUnknownOperation
	}
	return nil
}
