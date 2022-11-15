package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Token   string      `json:"token,omitempty"`
}

func WriteJsonRespnse(ctx *gin.Context, resp *Response) {
	ctx.JSON(resp.Status, resp)
}

func AbortJsonRespnse(ctx *gin.Context, resp *Response) {
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func SuccessLoginResponse(payload interface{}, token, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
		Token:   token,
	}
}

func SuccessResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Error:   err.Error(),
	}
}

func BadRequestResponse(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Error:   err.Error(),
	}
}

func UnprocessAbleEntityResponse(err interface{}) *Response {
	return &Response{
		Status:  http.StatusUnprocessableEntity,
		Message: "Unprocess able entity",
		Error:   err,
	}
}

func UnauthorizedResponse(err error, message string) *Response {
	return &Response{
		Status:  http.StatusUnauthorized,
		Message: message,
		Error:   err.Error(),
	}
}
