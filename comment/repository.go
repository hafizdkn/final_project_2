package comment

import (
	"gorm.io/gorm"
)

type Repository interface {
	UpdateComment(comment Comment) (Comment, error)
	CreateComment(comment Comment) (Comment, error)
	GetCommentById(id int) (Comment, error)
	GetComments() ([]Comment, error)
	DeleteComment(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateComment(comment Comment) (Comment, error) {
	if err := r.db.Debug().Create(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) GetComments() ([]Comment, error) {
	users := make([]Comment, 0)

	if err := r.db.Preload("User").Preload("Photo").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) GetCommentById(id int) (Comment, error) {
	var comment Comment

	if err := r.db.Debug().Where("id = ?", id).Take(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) UpdateComment(comment Comment) (Comment, error) {
	if err := r.db.Debug().Save(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *repository) DeleteComment(id int) error {
	if err := r.db.Debug().Delete(&Comment{}, id).Error; err != nil {
		return err
	}

	return nil
}
