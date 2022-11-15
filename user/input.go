package user

type UserRegisterInput struct {
	Username string `json:"username" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,numeric,min=8"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateInput struct {
	Username         string `json:"username" binding:"required,min=5"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=6"`
	Age              int    `json:"age" binding:"required,numeric,min=8"`
	EmailCurrentUser string
}
