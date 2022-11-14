package socialmedia

type MediaInput struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}
