package socialmedia

import (
	"time"

	"final_project_2/database"
)

type SocialMedia database.SocialMedia

type SocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}
