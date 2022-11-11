package photo

import "gorm.io/gorm"

type Repository interface {
	CreatePhoto(photo Photo) (Photo, error)
	UpdatePhoto(photo Photo) (Photo, error)
	GetPhotoById(id int) (Photo, error)
	GetPhotos() ([]Photo, error)
	DeletePhoto(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePhoto(photo Photo) (Photo, error) {
	if err := r.db.Debug().Create(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) GetPhotos() ([]Photo, error) {
	photos := make([]Photo, 0)

	if err := r.db.Preload("User").Find(&photos).Error; err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *repository) GetPhotoById(id int) (Photo, error) {
	var photo Photo

	if err := r.db.Debug().Where("id = ?", id).First(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) UpdatePhoto(photo Photo) (Photo, error) {
	if err := r.db.Debug().Save(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) DeletePhoto(id int) error {
	if err := r.db.Debug().Delete(&Photo{}, id).Error; err != nil {
		return err
	}

	return nil
}
