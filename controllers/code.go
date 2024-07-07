package controllers

type CodeRes int

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeErrorPassword
	CodeServerBusy
	CodeErrorOther

	CodeNeedLogin
	COdeInvalidToken
)

var codeMsg = map[CodeRes]string{
	CodeSuccess:       "成功",
	CodeInvalidParams: "请求参数错误",
	CodeUserExist:     "用户已经存在",
	CodeUserNotExist:  "用户不存在",
	CodeErrorPassword: "用户名或密码错误",
	CodeServerBusy:    "服务器忙",
	CodeErrorOther:    "未知错误",

	CodeNeedLogin:    "需要登录",
	COdeInvalidToken: "无效Token",
}

func (ctrl CodeRes) Msg() string {
	msg, ok := codeMsg[ctrl]
	if !ok {
		return codeMsg[CodeErrorOther]
	}
	return msg
}
