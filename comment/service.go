package comment

import (
	"errors"

	"final_project_2/helper"
)

type Service interface {
	UpdateComment(input CommentUpdateInput, commentId, currentUserId int) (Comment, error)
	CreateComment(input CommentCreateInput, userId int) (Comment, error)
	GetCommentById(commentId, currentUserId int) (Comment, error)
	DeleteComment(commentId, currentUserId int) error
	GetComments() ([]Comment, error)
}

type service struct {
	repository Repository
}

func NewServiceRepository(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateComment(input CommentCreateInput, userId int) (Comment, error) {
	var comment Comment

	comment.Message = input.Message
	comment.PhotoId = input.PhotoId
	comment.UserId = userId

	photo, err := s.repository.CreateComment(comment)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (s *service) GetComments() ([]Comment, error) {
	photos, err := s.repository.GetComments()
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (s *service) GetCommentById(commentId, currentUserId int) (Comment, error) {
	comment, err := s.repository.GetCommentById(commentId)
	if err != nil {
		return comment, err
	}

	if comment.UserId != currentUserId {
		return comment, errors.New(helper.ErrUnauthorized.Error())
	}

	return comment, nil
}

func (s *service) UpdateComment(input CommentUpdateInput, commentId, currentUserId int) (Comment, error) {
	comment, err := s.GetCommentById(commentId, currentUserId)
	if err != nil {
		return comment, err
	}

	comment.Message = input.Message

	updatedComment, err := s.repository.UpdateComment(comment)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func (s *service) DeleteComment(commentId, currentUserId int) error {
	_, err := s.GetCommentById(commentId, currentUserId)
	if err != nil {
		return err
	}

	if err := s.repository.DeleteComment(commentId); err != nil {
		return err
	}

	return nil
}
