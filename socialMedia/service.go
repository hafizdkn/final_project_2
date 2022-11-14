package socialmedia

import (
	"errors"

	"final_project_2/helper"
)

type Service interface {
	UpdateSocialMedia(input MediaInput, photoId, currentUserId int) (SocialMedia, error)
	CreateSocialMedia(input MediaInput, userId int) (SocialMedia, error)
	GetSocialMediaById(socialMediaId, currentUserId int) (SocialMedia, error)
	DeleteSocialMedia(socialMediaId, currentUserId int) error
	GetSocialMedias() ([]SocialMedia, error)
}

type service struct {
	repository Repository
}

func NewServiceRepository(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateSocialMedia(input MediaInput, userId int) (SocialMedia, error) {
	var socialMedia SocialMedia

	socialMedia.SocialMeidaUrl = input.SocialMediaUrl
	socialMedia.Name = input.Name
	socialMedia.UserId = userId

	socialMedia, err := s.repository.CreateSocialMedia(socialMedia)
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (s *service) GetSocialMedias() ([]SocialMedia, error) {
	socialMedia, err := s.repository.GetSocialMedias()
	if err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (s *service) GetSocialMediaById(photoId, currentUserId int) (SocialMedia, error) {
	socialMedia, err := s.repository.GetSocialMediaById(photoId)
	if err != nil {
		return socialMedia, err
	}

	if socialMedia.UserId != currentUserId {
		return socialMedia, errors.New(helper.ErrUnauthorized.Error())
	}

	return socialMedia, nil
}

func (s *service) UpdateSocialMedia(input MediaInput, photoId, currentUserId int) (SocialMedia, error) {
	socialMedia, err := s.GetSocialMediaById(photoId, currentUserId)
	if err != nil {
		return socialMedia, err
	}

	socialMedia.Name = input.Name
	socialMedia.SocialMeidaUrl = input.SocialMediaUrl

	updatedSocialMedia, err := s.repository.UpdateSocialMedia(socialMedia)
	if err != nil {
		return socialMedia, err
	}

	return updatedSocialMedia, nil
}

func (s *service) DeleteSocialMedia(photoId, currentUserId int) error {
	_, err := s.GetSocialMediaById(photoId, currentUserId)
	if err != nil {
		return err
	}

	if err := s.repository.DeleteSocialMedia(photoId); err != nil {
		return err
	}

	return nil
}
