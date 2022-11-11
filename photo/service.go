package photo

type Service interface {
	UpdatePhoto(input PhotoUpdateInput, id int) (Photo, error)
	CreatePhoto(input PhotoInput, userId int) (Photo, error)
	GetPhotoById(id int) (Photo, error)
	GetPhotos() ([]Photo, error)
	DeletePhoto(id int) error
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

func (s *service) GetPhotoById(id int) (Photo, error) {
	photo, err := s.repository.GetPhotoById(id)
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (s *service) UpdatePhoto(input PhotoUpdateInput, id int) (Photo, error) {
	photo, err := s.repository.GetPhotoById(id)
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

func (s *service) DeletePhoto(id int) error {
	if err := s.repository.DeletePhoto(id); err != nil {
		return err
	}

	return nil
}
