package client

import "errors"

const (
	errMsgUserNotFound      = "пользователь с такими данными не найден"
	errMsgUserAlreadyExists = "пользователь уже существует"
	errMsgIncorrectAuthData = "некорректные авторизационные данные"
	errMsgTooLongPass       = "слишком длинный пароль" // #nosec G101
)

var (
	ErrUserNotFound      = errors.New(errMsgUserNotFound)      // ErrUserNotFound ..
	ErrUserAlreadyExists = errors.New(errMsgUserAlreadyExists) // ErrUserAlreadyExists ..
	ErrIncorrectAuthData = errors.New(errMsgIncorrectAuthData) // ErrIncorrectAuthData ..
	ErrTooLongPass       = errors.New(errMsgTooLongPass)       // ErrTooLongPass ..
)
