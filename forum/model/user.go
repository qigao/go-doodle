package model

type Response struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ProfileType struct {
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Bio       *string `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}
