package request

type UserUpdateRequest struct {
	Username string `json:"username" validate:"alphanum"`
}

type UserUpdateEmailRequest struct {
	Email string `json:"email" validate:"required"`
}

type UserUpdateRoleRequest struct {
	Role string `json:"role" validate:"required,oneof=admin consumer"`
}
