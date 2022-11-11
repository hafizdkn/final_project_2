package middleware

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"final_project_2/auth"
	"final_project_2/helper"
	"final_project_2/user"
)

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		unauthorizedResponse := "Unauthorized"
		cantParseTokenReponse := "Can't parse token"
		tokenIvalidResponse := "Invalid Token"

		headerToken := c.Request.Header.Get("Authorization")
		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			err := errors.New(unauthorizedResponse)
			helper.AbortJsonRespnse(c, helper.UnauthorizedResponse(err, unauthorizedResponse))
			return
		}

		stringToken := strings.Split(headerToken, " ")
		if len(stringToken) != 2 {
			err := errors.New(cantParseTokenReponse)
			helper.AbortJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
		}

		token, err := authService.ValidateToken(stringToken[1])
		if err != nil {
			helper.AbortJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			err := errors.New(tokenIvalidResponse)
			helper.AbortJsonRespnse(c, helper.UnauthorizedResponse(err, ""))
			return
		}

		userId := int(claim["id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			helper.WriteJsonRespnse(c, helper.InternalServerError(err))
			return
		}

		c.Set("currentUser", user)
	}
}
