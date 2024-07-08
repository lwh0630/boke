package mysql

import "errors"

var (
	ErrUserNotExist    = errors.New("用户不存在")
	ErrUserExist       = errors.New("用户已存在")
	ErrInvalidPassWord = errors.New("用户名或密码错误")

	ErrSelectCommunityList = errors.New("community select 错误")
	ErrInvalidCommunityId  = errors.New("community id 无效")
)
