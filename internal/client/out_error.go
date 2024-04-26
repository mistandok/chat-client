package client

import "errors"

const (
	errMsgUserNotFound      = "пользователь с такими данными не найден"
	errMsgUserAlreadyExists = "пользователь уже существует"
	errMsgIncorrectAuthData = "некорректные авторизационные данные"
	errMsgTooLongP          = "слишком длинный пароль"
)

var (
	ErrUserNotFound      = errors.New(errMsgUserNotFound)      // ErrUserNotFound ..
	ErrUserAlreadyExists = errors.New(errMsgUserAlreadyExists) // ErrUserAlreadyExists ..
	ErrIncorrectAuthData = errors.New(errMsgIncorrectAuthData) // ErrIncorrectAuthData ..
	ErrTooLongPass       = errors.New(errMsgTooLongP)          // ErrTooLongPass ..
)
