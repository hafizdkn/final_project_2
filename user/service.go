package user

import (
	"final_project_2/helper"
)

type Service interface {
	CreateUser(input UserRegisterInput) (User, error)
	UserLogin(input UserLogin) (UserResponse, error)
	UpdateUser(input UserUpdateInput) (User, error)
	GetUserById(input int) (User, error)
	GetUsers() ([]User, error)
	DeleteUser(id int) error
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(input UserRegisterInput) (User, error) {
	var user User
	var err error

	user.Username = input.Username
	user.Email = input.Email
	user.Age = input.Age

	passwordHashed, err := helper.GeneratePasswordHash(input.Password)
	if err != nil {
		return user, err
	}
	user.Password = passwordHashed

	user, err = s.repository.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UserLogin(input UserLogin) (UserResponse, error) {
	var userResponse UserResponse

	email := input.Email
	InputPassword := input.Password

	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return userResponse, err
	}

	if err := helper.ComparePasswordHash(user.Password, InputPassword); err != nil {
		return userResponse, err
	}

	userResponse.ID = user.ID
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Age = user.Age

	return userResponse, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateUser(input UserUpdateInput) (User, error) {
	currentUser, err := s.repository.GetUserByEmail(input.EmailCurrentUser)
	if err != nil {
		return currentUser, err
	}

	passwordHashed, err := helper.GeneratePasswordHash(input.Password)
	if err != nil {
		return currentUser, err
	}

	currentUser.Username = input.Username
	currentUser.Password = passwordHashed
	currentUser.Email = input.Email
	currentUser.Age = input.Age

	newUser, err := s.repository.UpdateUser(currentUser)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) DeleteUser(id int) error {
	if err := s.repository.Deleteuser(id); err != nil {
		return err
	}

	return nil
}

func (s *service) GetUsers() ([]User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}
