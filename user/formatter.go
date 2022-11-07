package user

import "time"

type FormatUser struct {
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Age       int        `json:"age"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func CreateResponse(user User) FormatUser {
	var fUser FormatUser

	fUser.Username = user.Username
	fUser.Password = user.Password
	fUser.Email = user.Email
	fUser.Age = user.Age

	return fUser
}

func UpdateResponse(user User) FormatUser {
	var fUser FormatUser

	fUser.UpdatedAt = &user.UpdatedAt
	fUser.Username = user.Username
	fUser.Password = user.Password
	fUser.Email = user.Email
	fUser.Age = user.Age

	return fUser
}
