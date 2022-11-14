package socialmedia

import "gorm.io/gorm"

type Repository interface {
	CreateSocialMedia(socialMedia SocialMedia) (SocialMedia, error)
	UpdateSocialMedia(socialMedia SocialMedia) (SocialMedia, error)
	GetSocialMediaById(id int) (SocialMedia, error)
	GetSocialMedias() ([]SocialMedia, error)
	DeleteSocialMedia(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateSocialMedia(socialMedia SocialMedia) (SocialMedia, error) {
	if err := r.db.Debug().Create(&socialMedia).Error; err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *repository) GetSocialMedias() ([]SocialMedia, error) {
	photos := make([]SocialMedia, 0)

	if err := r.db.Preload("User").Find(&photos).Error; err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *repository) GetSocialMediaById(id int) (SocialMedia, error) {
	var socialMedia SocialMedia

	if err := r.db.Debug().Where("id = ?", id).First(&socialMedia).Error; err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *repository) UpdateSocialMedia(socialMedia SocialMedia) (SocialMedia, error) {
	if err := r.db.Debug().Save(&socialMedia).Error; err != nil {
		return socialMedia, err
	}

	return socialMedia, nil
}

func (r *repository) DeleteSocialMedia(id int) error {
	if err := r.db.Debug().Delete(&SocialMedia{}, id).Error; err != nil {
		return err
	}

	return nil
}
