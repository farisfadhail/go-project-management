package request

type RegisterRequest struct {
	FullName        string `json:"full_name" bson:"full_name" validate:"required"`
	Username        string `json:"username" bson:"username" validate:"required,alphanum"`
	Email           string `json:"email" bson:"email" validate:"required,email"`
	Password        string `json:"password" bson:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" bson:"confirm_password" validate:"required,min=6"`
	//Role            string `json:"role" bson:"role" validate:"required,oneof=admin user"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=6"`
}

type UpdatePassword struct {
	OldPassword string `json:"old_password" bson:"oldPassword" validate:"required,min=6"`
	NewPassword string `json:"new_password" bson:"newPassword" validate:"required,min=6"`
}
