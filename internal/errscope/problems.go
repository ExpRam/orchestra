package errscope

import "errors"

type Problems []error

func (p *Problems) Add(err error) {
	if err == nil {
		return
	}

	*p = append(*p, err)
}

func (p Problems) Err() error {
	if len(p) == 0 {
		return nil
	}

	return p
}

func (p Problems) Error() string {
	return errors.Join(p...).Error()
}

func (p Problems) Unwrap() []error {
	return p
}
