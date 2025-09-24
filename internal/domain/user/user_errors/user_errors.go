package user_errors

import "errors"

var ErrorUserNotFound error = errors.New("user not found")
var ErrorUserEmptyInsert error = errors.New("empty insert")
var ErrorUserIsAlreadyExist error = errors.New("user is already exist")
var ErrorInvalidPassword error = errors.New("invalid password")
var ErrorUserNotExist error = errors.New("user not exists")
