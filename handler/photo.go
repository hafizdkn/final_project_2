package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"final_project_2/helper"
	"final_project_2/photo"
	"final_project_2/user"
)

type photoHandler struct {
	photoService photo.Service
}

func NewPhotoHandler(userService photo.Service) *photoHandler {
	return &photoHandler{userService}
}

func (h *photoHandler) CreatePhoto(c *gin.Context) {
	var input photo.PhotoInput

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	p, err := h.photoService.CreatePhoto(input, userId)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessCreateResponse(p, "Success create photo")
	helper.WriteJsonRespnse(c, response)
}

func (h *photoHandler) GetPhotos(c *gin.Context) {
	p, err := h.photoService.GetPhotos()
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse(p, "Success get photo")
	helper.WriteJsonRespnse(c, response)
}

func (h *photoHandler) UpdatePhoto(c *gin.Context) {
	var input photo.PhotoUpdateInput

	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	p, err := h.photoService.UpdatePhoto(input, photoId, currentUserId)
	if err != nil {
		if err.Error() == helper.ErrUnauthorized.Error() {
			helper.WriteJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse(p, "Success update photo")
	helper.WriteJsonRespnse(c, response)
}

func (h *photoHandler) DeletePhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	if err := h.photoService.DeletePhoto(photoId, currentUserId); err != nil {
		if err.Error() == helper.ErrUnauthorized.Error() {
			helper.WriteJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse(nil, "Your photo has been successfully deleted")
	helper.WriteJsonRespnse(c, response)
}
