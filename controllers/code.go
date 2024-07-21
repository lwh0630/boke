package controllers

type CodeRes int

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken

	CodeErrorPassword
)

var codeMsg = map[CodeRes]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",

	CodeErrorPassword: "用户密码错误",
}

func (ctrl CodeRes) Msg() string {
	msg, ok := codeMsg[ctrl]
	if !ok {
		return codeMsg[CodeServerBusy]
	}
	return msg
}
