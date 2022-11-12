package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"final_project_2/comment"
	"final_project_2/helper"
	"final_project_2/user"
)

type commentHandler struct {
	userService comment.Service
}

func NewCommentHandler(commetService comment.Service) *commentHandler {
	return &commentHandler{commetService}
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	var input comment.CommentCreateInput

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	comment, err := h.userService.CreateComment(input, userId)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessCreateResponse(comment, "Success create comment")
	helper.WriteJsonRespnse(c, response)
}

func (h *commentHandler) GetComments(c *gin.Context) {
	comments, err := h.userService.GetComments()
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessCreateResponse(comments, "Success get comments")
	helper.WriteJsonRespnse(c, response)
}

func (h *commentHandler) UpdateComment(c *gin.Context) {
	var input comment.CommentUpdateInput

	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	p, err := h.userService.UpdateComment(input, commentId, currentUserId)
	if err != nil {
		if err.Error() == helper.ErrUnauthorized.Error() {
			helper.WriteJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse(p, "Success update comment")
	helper.WriteJsonRespnse(c, response)
}

func (h *commentHandler) DeleteComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	err = h.userService.DeleteComment(commentId, currentUserId)
	if err != nil {
		if err.Error() == helper.ErrUnauthorized.Error() {
			helper.WriteJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse("", "Your comment has benn successfully deleted")
	helper.WriteJsonRespnse(c, response)
}
