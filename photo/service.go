package photo

import (
	"errors"

	"final_project_2/helper"
)

type Service interface {
	UpdatePhoto(input PhotoUpdateInput, photoId, currentUserId int) (Photo, error)
	CreatePhoto(input PhotoInput, userId int) (Photo, error)
	GetPhotoById(photoId, currentUserId int) (Photo, error)
	DeletePhoto(photoId, currentUserId int) error
	GetPhotos() ([]Photo, error)
}

type service struct {
	repository Repository
}

func NewServiceRepository(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePhoto(input PhotoInput, userId int) (Photo, error) {
	var photo Photo

	photo.Title = input.Title
	photo.PhotoUrl = input.PhotoUrl
	photo.Caption = input.Caption
	photo.UserId = userId

	photo, err := s.repository.CreatePhoto(photo)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (s *service) GetPhotos() ([]Photo, error) {
	photos, err := s.repository.GetPhotos()
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (s *service) GetPhotoById(photoId, currentUserId int) (Photo, error) {
	photo, err := s.repository.GetPhotoById(photoId)
	if err != nil {
		return photo, err
	}

	if photo.UserId != currentUserId {
		return photo, errors.New(helper.ErrUnauthorized.Error())
	}

	return photo, nil
}

func (s *service) UpdatePhoto(input PhotoUpdateInput, photoId, currentUserId int) (Photo, error) {
	/*
		method ini akan memanggil method GetPhotoById dan mengembalikan nilai photo berdasarkan request photoId,
		kemudian nilai kembalian akan di cek apakah field photo.UserId sama dengan currentUserid, jika tidak sama,
		kembalikan error, jika sama update nilai photo.
	*/
	photo, err := s.GetPhotoById(photoId, currentUserId)
	if err != nil {
		return photo, err
	}

	photo.Title = input.Title
	photo.PhotoUrl = input.PhotoUrl
	if input.Caption != "" {
		photo.Caption = input.Caption
	}

	updatedPhoto, err := s.repository.UpdatePhoto(photo)
	if err != nil {
		return photo, err
	}

	return updatedPhoto, nil
}

func (s *service) DeletePhoto(photoId, currentUserId int) error {
	/*
		method ini akan memanggil method GetPhotoById dan mengembalikan nilai photo berdasarkan request photoId,
		kemudian nilai kembalian akan di cek apakah field photo.UserId sama dengan currentUserid, jika tidak sama,
		kembalikan error, jika sama delete data photo
	*/
	_, err := s.GetPhotoById(photoId, currentUserId)
	if err != nil {
		return err
	}

	if err := s.repository.DeletePhoto(photoId); err != nil {
		return err
	}

	return nil
}
