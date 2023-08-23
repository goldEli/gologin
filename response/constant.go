package response

const (
	ResponseCodeOk                  = 10000
	ResponseCodeUnauthorized        = 20000
	ResponseCodeForbidden           = 20001
	ResponseCodeUserPasswordError   = 20002
	ResponseCodeBadRequest          = 20003
	ResponseResourceError           = 20004
	ResponseCodeInternalServerError = 30000
)

var ResponseCodeMap = map[int]string{
	10000: "SUCCESS",
	20000: "未登录",
	20001: "没有权限，请联系管理员",
	20002: "用户名/密码错误",
	20003: "参数不合法",
	20004: "资源冲突",
	30000: "系统异常，请稍后再试！",
}

func ResponseMessage(code int) string {
	return ResponseCodeMap[code]
}
