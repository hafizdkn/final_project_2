package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(email string) (User, error)
	UpdateUser(user User) (User, error)
	CreateUser(user User) (User, error)
	GetUserById(id int) (User, error)
	GetUsers() ([]User, error)
	Deleteuser(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user User) (User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUsers() ([]User, error) {
	users := make([]User, 0)

	if err := r.db.Debug().Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) GetUserByEmail(email string) (User, error) {
	var user User

	if err := r.db.Debug().Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetUserById(id int) (User, error) {
	var user User

	if err := r.db.Debug().Where("id = ?", id).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	if err := r.db.Debug().Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Deleteuser(id int) error {
	if err := r.db.Debug().Delete(&User{}, id).Error; err != nil {
		return err
	}

	return nil
}
