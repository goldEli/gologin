package vo

type LoginVo struct {
	Email    string `json:"email" validate:"required" label:"邮件"`
	Password string `json:"password" validate:"required" label:"密码"`
}
