package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"final_project_2/helper"
	socialmedia "final_project_2/socialMedia"
	"final_project_2/user"
)

type socialMediaHandler struct {
	socialMediaService socialmedia.Service
}

func NewSocialMediahandler(userService socialmedia.Service) *socialMediaHandler {
	return &socialMediaHandler{userService}
}

func (h *socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var input socialmedia.MediaInput

	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	p, err := h.socialMediaService.CreateSocialMedia(input, userId)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessCreateResponse(p, "Success create social media")
	helper.WriteJsonRespnse(c, response)
}

func (h *socialMediaHandler) GetSocialMedias(c *gin.Context) {
	p, err := h.socialMediaService.GetSocialMedias()
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse(p, "Success get social media")
	helper.WriteJsonRespnse(c, response)
}

func (h *socialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	var input socialmedia.MediaInput

	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	p, err := h.socialMediaService.UpdateSocialMedia(input, socialMediaId, currentUserId)
	if err != nil {
		if err.Error() == helper.ErrUnauthorized.Error() {
			helper.WriteJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse(p, "Success update social media")
	helper.WriteJsonRespnse(c, response)
}

func (h *socialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	currentUserId := currentUser.ID

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	if err := h.socialMediaService.DeleteSocialMedia(socialMediaId, currentUserId); err != nil {
		if err.Error() == helper.ErrUnauthorized.Error() {
			helper.WriteJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessResponse(nil, "Your social media has been successfully deleted")
	helper.WriteJsonRespnse(c, response)
}
