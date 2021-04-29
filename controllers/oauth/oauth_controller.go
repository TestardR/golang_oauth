package oauth

import (
	"net/http"

	"example.com/domain/oauth"
	"example.com/services"
	"example.com/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
