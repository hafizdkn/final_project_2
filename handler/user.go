package handler

import (
	"github.com/gin-gonic/gin"

	"final_project_2/auth"
	"final_project_2/helper"
	"final_project_2/user"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var input user.UserRegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	u, err := h.userService.CreateUser(input)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessCreateResponse(user.CreateResponse(u), "Success create user")
	helper.WriteJsonRespnse(c, response)
}

func (h *userHandler) UserLogin(c *gin.Context) {
	var input user.UserLogin
	userNotValid := "Invalidusername or password"

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	user, err := h.userService.UserLogin(input)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.UnauthorizedResponse(err, userNotValid))
		return
	}

	token, err := h.authService.GenerateToken(user.ID, user.Email)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.BadRequestResponse(err))
		return
	}

	loginResponse := helper.SuccessLoginResponse(user, token, "Login success")
	helper.WriteJsonRespnse(c, loginResponse)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var input user.UserUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validatorError := helper.FormatValidationError(err)
		helper.WriteJsonRespnse(c, helper.UnprocessAbleEntityResponse(validatorError))
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.EmailCurrentUser = currentUser.Email

	newUser, err := h.userService.UpdateUser(input)
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	response := helper.SuccessCreateResponse(user.UpdateResponse(newUser), "Success update user")
	helper.WriteJsonRespnse(c, response)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	if err := h.userService.DeleteUser(userId); err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	helper.WriteJsonRespnse(c, helper.SuccessResponse("", "Your account has been successfully deleted"))
}

func (h *userHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		helper.WriteJsonRespnse(c, helper.InternalServerError(err))
		return
	}

	helper.WriteJsonRespnse(c, helper.SuccessResponse(users, "Succes get users"))
}
