package vo

type LoginVo struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
