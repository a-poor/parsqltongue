package expr

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	if len(values) < 1 {
		return ErrTooFewValues
	}
	if len(values) > 1 {
		return ErrTooManyValues
	}
	*b = values[0] == "true"
	return nil
}
