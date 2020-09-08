package model

import "errors"

//各种错误
var (
	ErrUsrNotExist = errors.New("用户不存在")
	ErrUsrExisted  = errors.New("用户已存在")
	ErrUsrPwdErr   = errors.New("密码错误")
)
