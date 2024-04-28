package repository

import "errors"

const errMsgTokensNotFound = "Токены не найдены"

var ErrTokensNotFound = errors.New(errMsgTokensNotFound) // ErrTokensNotFound ..
