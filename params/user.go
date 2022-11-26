package params

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

type CreateUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UpdateUser struct {
	UserId int    `json:"user_id" validate:"required"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
