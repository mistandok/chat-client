package functional_error

import "errors"

type functionalError struct {
	err error
}

// NewFunctionalError ..
func NewFunctionalError(err error) *functionalError {
	return &functionalError{
		err: err,
	}
}

// Error ..
func (r *functionalError) Error() string {
	return r.err.Error()
}

// Unwrap ..
func (r *functionalError) Unwrap() error {
	return r.err
}

// IsFunctionalError ..
func IsFunctionalError(err error) bool {
	var ce *functionalError
	return errors.As(err, &ce)
}

// GetFunctionalError ..
func GetFunctionalError(err error) *functionalError {
	var ce *functionalError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
