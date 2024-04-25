package client

import "errors"

const (
	errMsgUserNotFound      = "пользователь с такими данными не найден"
	errMsgUserAlreadyExists = "пользователь уже существует"
	errMsgIncorrectAuthData = "некорректные авторизационные данные"
	errMsgTooLongPass       = "слишком длинный пароль"
)

var (
	ErrUserNotFound      = errors.New(errMsgUserNotFound)
	ErrUserAlreadyExists = errors.New(errMsgUserAlreadyExists)
	ErrIncorrectAuthData = errors.New(errMsgIncorrectAuthData)
	ErrTooLongPass       = errors.New(errMsgTooLongPass)
)
