package vo

type LoginVo struct {
	Email    string `json:"email" validate:"required" label:"邮件" field_error_info:"邮件不能为空"`
	Password string `json:"password" validate:"required" label:"密码" `
}
