package repository

import "errors"

const errMsgTokensNotFound = "Токены не найдены" // #nosec G101

var ErrTokensNotFound = errors.New(errMsgTokensNotFound) // ErrTokensNotFound ..
