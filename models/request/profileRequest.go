package request

type ProfileRequest struct {
	FirstName string `json:"first_name" bson:"first_name" validate:"required"`
	LastName  string `json:"last_name" bson:"last_name" validate:"required"`
	Phone     string `json:"phone" bson:"phone" validate:"required"`
	About     string `json:"about" bson:"about" validate:"required"`
}

type ProfileUpdateRequest struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Phone     string `json:"phone" bson:"phone"`
	About     string `json:"about" bson:"about"`
}
