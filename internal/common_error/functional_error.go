package common_error

import "errors"

type commonError struct {
	msg string
	err error
}

// NewCommonError ..
func NewCommonError(msg string, err error) *commonError {
	return &commonError{
		msg: msg,
		err: err,
	}
}

// Error ..
func (r *commonError) Error() string {
	return r.msg
}

// Unwrap ..
func (r *commonError) Unwrap() error {
	return r.err
}

// IsCommonError ..
func IsCommonError(err error) bool {
	var ce *commonError
	return errors.As(err, &ce)
}

// GetCommonError ..
func GetCommonError(err error) *commonError {
	var ce *commonError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
