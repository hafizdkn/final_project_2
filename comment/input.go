package comment

type CommentCreateInput struct {
	Message string `json:"message" binding:"required"`
	PhotoId int    `json:"photo_id" binding:"required"`
}

type CommentUpdateInput struct {
	Message string `json:"message" binding:"required"`
}
