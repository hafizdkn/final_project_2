package photo

type PhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}

type PhotoUpdateInput struct {
	ID       string `json:"id"`
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}
