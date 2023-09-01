package dto

type LoginDto struct {
	Jwt   string `json:"jwt"`
	Email string `json:"email"`
}
